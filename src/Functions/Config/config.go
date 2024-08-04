/* Nome do pacote atual */
package config

/* Importa os módulos necessários */
import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

/* AuthConfig contém detalhes de autenticação */
type AuthConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/* WebSocketConfig contém configuração relacionada ao WebSocket */
type WebSocketConfig struct {
	Description string `json:"description"`
	Value       string `json:"value"`
}

/* PostRequestConfig contém configuração para solicitações post */
type PostRequestConfig struct {
	Description string `json:"description"`
	Value       string `json:"value"`
}

/* CasesConfig contém configuração para casos com valor booleano */
type CasesConfig struct {
	Description string `json:"description"`
	Value       bool   `json:"value"`
}

/* UpdateIntervalConfig contém configuração para intervalos de atualização */
type UpdateIntervalConfig struct {
	Description string `json:"description"`
	Value       int    `json:"value"`
}

/* Config contém todas as configurações da aplicação */
type Config struct {
	WebSocket      WebSocketConfig      `json:"WebSocket"`
	Auth           AuthConfig           `json:"Auth"`
	Cases          CasesConfig          `json:"Cases"`
	PostRequest    PostRequestConfig    `json:"PostRequest"`
	UpdateInterval UpdateIntervalConfig `json:"UpdateInterval"`
}

/* ConfigData é a variável global para armazenar a configuração */
var ConfigData Config

/* once garante que a configuração seja carregada apenas uma vez */
var once sync.Once

/* LoadConfig lê e analisa o arquivo de configuração JSON */
func LoadConfig(filePath string) error {
	/* Variável para armazenar o erro */
	var err error

	/* Garante que a configuração seja carregada apenas uma vez */
	once.Do(func() {
		/* Abre o arquivo de configuração */
		file, openErr := os.Open(filePath)

		/* Se ocorrer um erro ao abrir o arquivo, retorna log */
		if openErr != nil {
			err = fmt.Errorf("Falha ao abrir o arquivo %s: %w", filePath, openErr)
			return
		}

		/* Fecha o arquivo após a execução */
		defer file.Close()

		/* Cria um decoder para ler o arquivo JSON */
		decoder := json.NewDecoder(file)

		/* Decodifica o arquivo JSON para a variável ConfigData */
		if decodeErr := decoder.Decode(&ConfigData); decodeErr != nil {
			err = fmt.Errorf("Falha ao decodificar o arquivo %s: %w", filePath, decodeErr)
		}
	})

	/* Retorna o erro, se houver */
	return err
}

/* GetConfig retorna a configuração carregada */
func GetConfig() *Config {
	return &ConfigData
}
