/* Nome do pacote atual */
package updater

/* Importa os módulos necessários */
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	color "ceres/src/Functions/Color"
)

/* PackageInfo contém as informações do package.json */
type PackageInfo struct {
	Version   string `json:"version"`
	BuildDate string `json:"build_date"`
	BuildName string `json:"build_name"`
	Homepage  string `json:"homepage"`
}

/* CheckUpdates verifica se há atualizações comparando os arquivos package.json local e remoto */
func CheckUpdates(timeout time.Duration) bool {
	/* Váriavel para armazenar os dados do JSON */
	var localPack, remotePack PackageInfo

	/* Lê o arquivo package.json local */
	localFile, err := os.ReadFile("package.json")

	/* Se ocorrer um erro ao ler o arquivo local */
	if err != nil {
		fmt.Printf("Erro ao ler o arquivo local: %v\n", err)
		return false
	}

	/* Decodifica o JSON local */
	if err := json.Unmarshal(localFile, &localPack); err != nil {
		fmt.Printf("Erro ao decodificar o JSON local: %v\n", err)
		return false
	}

	/* Define a URL para o package.json remoto */
	remoteURL := "https://raw.githubusercontent.com/KillovSky/Ceres/main/package.json"

	/* Busca o package.json remoto */
	client := &http.Client{Timeout: timeout}
	resp, err := client.Get(remoteURL)

	/* Se ocorrer um erro ao buscar a versão remota */
	if err != nil {
		fmt.Printf("Erro ao buscar a versão remota: %v\n", err)
		return false
	}

	/* Fecha o body */
	defer resp.Body.Close()

	/* Se o status da resposta não for OK */
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Erro ao acessar o arquivo remoto: Status %d\n", resp.StatusCode)
		return false
	}

	/* Lê e decodifica o JSON remoto */
	remoteData, err := io.ReadAll(resp.Body)

	/* Se ocorrer um erro ao ler a resposta remota */
	if err != nil {
		fmt.Printf("Erro ao ler a resposta remota: %v\n", err)
		return false
	}

	/* Decodifica o JSON remoto */
	if err := json.Unmarshal(remoteData, &remotePack); err != nil {
		/* Se der erro retorna false e printa o erro */
		fmt.Printf("Erro ao decodificar o JSON remoto: %v\n", err)
		return false
	}

	/* Compara versões e outras informações */
	if localPack.Version == remotePack.Version &&
		localPack.BuildDate == remotePack.BuildDate &&
		localPack.BuildName == remotePack.BuildName {
		fmt.Println(
			color.ColorPrint("[VERSION]", "green"),
			color.ColorPrint("Valeu por me manter atualizada!", "green"),
		)

		/* Retorna false para sinalizar sem updates */
		return false
	}

	/* Exibe informações sobre a atualização disponível */
	fmt.Println(
		color.ColorPrint("[VERSION]", "green"),
		color.ColorPrint("ATUALIZAÇÃO DISPONÍVEL → ", "red"),
		"[",
		color.ColorPrint(remotePack.Version, "magenta"),
		"~",
		color.ColorPrint(strings.ToUpper(remotePack.BuildName), "blue"),
		"~",
		color.ColorPrint(strings.ToUpper(remotePack.BuildDate), "yellow"),
		"] |",
		color.ColorPrint(remotePack.Homepage, "green"),
	)

	/* Retorna true para sinalizar que tem updates */
	return true
}
