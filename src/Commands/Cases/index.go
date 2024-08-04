/* Nome do pacote atual */
package cases

/* Importa os módulos necessários */
import (
	config "ceres/src/Functions/Config"
	kill "ceres/src/Functions/Kill"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/* Define uma variável global para armazenar a configuração */
var configData config.Config

/* checkCase verifica se o comando corresponde ao nome especificado */
func checkCase(command, name string, onlyCmd, isCmd bool) bool {
	/* Retorna true se a combinação de valores corresponder aos critérios */
	return (strings.Contains(command, name) && !onlyCmd) || (command == name && onlyCmd && isCmd)
}

/* executeGoCode executa um código Go passado como string e retorna a saída padrão, erro padrão e um possível erro */
func executeGoCode(code string) (string, string, error) {
	/* Cria um arquivo temporário para armazenar o código Go */
	file, err := os.CreateTemp("", "*.go")

	/* Se ocorrer um erro ao criar o arquivo temporário */
	if err != nil {
		/* Retorna valores vazios e o erro */
		return "", "", err
	}

	/* Remove o arquivo temporário após a execução */
	defer os.Remove(file.Name())

	/* Escreve o código Go no arquivo temporário */
	_, err = file.WriteString(code)

	/* Se falhar na escrita do código */
	if err != nil {
		/* Retorna valores vazios e o erro se a escrita falhar */
		return "", "", err
	}

	/* Fecha o arquivo para garantir que os dados sejam gravados corretamente */
	file.Close()

	/* Cria um comando para executar o código Go usando `go run` */
	cmd := exec.Command("go", "run", file.Name())

	/* Captura a saída padrão e o erro padrão do comando */
	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	/* Executa o comando e captura erros */
	err = cmd.Run()

	/* Se der erro na execução do código */
	if err != nil {
		/* Retorna a saída e o erro formatado se a execução falhar */
		return stdout.String(), stderr.String(), fmt.Errorf("failed to execute code: %w", err)
	}

	/* Retorna a saída padrão, o erro padrão e nil se a execução for bem-sucedida */
	return stdout.String(), stderr.String(), nil
}

/* HandleCommand processa o comando recebido e executa ações conforme necessário */
func HandleCommand(env map[string]interface{}) (map[string]interface{}, error) {
	/* Extrai informações do ambiente fornecido */
	chatId, _ := env["chatId"].(string)
	isCmd, _ := env["isCmd"].(bool)
	command, _ := env["command"].(string)
	isOwner, _ := env["isOwner"].(bool)
	arg, _ := env["arg"].(string)

	/* Se não for um comando */
	if !isCmd {
		/* Usa o body da mensagem como comando */
		command, _ = env["body"].(string)
	}

	/* Define a variável de resposta */
	var reply interface{}

	/* Se existir corretamente */
	if r, ok := env["reply"].(map[string]interface{}); ok {
		/* Define como o valor da Object */
		reply = r

		/* Caso contrário define como nil */
	} else {
		reply = nil
	}

	/* Processa o comando com base em diferentes casos usando switch */
	switch {

	/* Teste de comando sem prefix */
	case checkCase(command, "gotest123+@", false, isCmd),
		checkCase(command, "go test 123 +@", false, isCmd),
		checkCase(command, "GO TEST 123 +@", false, isCmd):
		/* Envia mensagem de sucesso para o chat */
		data, err := kill.SendMessage(chatId, map[string]interface{}{"text": "✔️ OK!"}, reply, nil)

		/* Se a mensagem falhar no envio */
		if err != nil {
			return nil, err
		}

		/* Retorna os dados enviados e nil para indicar sucesso */
		return data, nil

	/* Golang Eval */
	case checkCase(command, "goeval", true, isCmd) && isOwner:
		/* Executa o código Go passado como argumento */
		stdout, stderr, err := executeGoCode(arg)

		/* Se der erro na execução do eval do GO */
		if err != nil {
			/* Cria uma variável para receber os dados do erro */
			var message string

			/* Erro e Stdout */
			if len(stderr) > 0 {
				message = fmt.Sprintf("❌ Erro ao executar o código:\nStderr:\n%s\nErro:\n%s", stderr, err)

				/* Só o erro */
			} else {
				message = fmt.Sprintf("❌ Erro ao executar o código:\nErro:\n%s", err)
			}

			/* Envia mensagem de erro para o chat */
			data, err := kill.SendMessage(chatId, map[string]interface{}{"text": message}, reply, nil)

			/* Se falhar no envio geral da mensagem */
			if err != nil {
				return nil, err
			}

			/* Retorna os dados enviados e nil para indicar sucesso */
			return data, nil
		}

		/* Cria uma mensagem com base em stdout e stderr */
		var message string

		/* Se stderr existir, insere ele */
		if len(stderr) > 0 {
			message = fmt.Sprintf("Stdout:\n%s\nStderr:\n%s", stdout, stderr)

			/* Se não, só envia o Stdout mesmo */
		} else {
			message = fmt.Sprintf("Stdout:\n%s", stdout)
		}

		/* Envia mensagem com a saída do código para o chat */
		data, err := kill.SendMessage(chatId, map[string]interface{}{"text": message}, reply, nil)

		/* Se a mensagem de eval falhar em ser enviada */
		if err != nil {
			return nil, err
		}

		/* Retorna os dados enviados e nil para indicar sucesso */
		return data, nil

	/* Caso de comando não existente */
	case isCmd && configData.Cases.Value:
		/* Envia mensagem informando que o comando não existe */
		data, err := kill.SendRaw(chatId, map[string]interface{}{"text": "Esse comando não existe ainda!"}, reply, nil)

		/* Retorna erro se o envio da mensagem falhar */
		if err != nil {
			return nil, err
		}

		/* Retorna os dados enviados e nil para indicar sucesso */
		return data, nil

	/* Caso não caia em nenhum comando acima */
	default:
		/* Retorna nil se nenhum comando foi correspondido */
		return nil, nil
	}
}
