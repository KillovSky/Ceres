/* Nome do pacote atual */
package kill

/* Importa os módulos necessários */
import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	config "ceres/src/Functions/Config"
)

/* postMessage envia uma mensagem via POST request e retorna a resposta do servidor */
func postMessage(url string, data map[string]interface{}, timeout time.Duration) (*http.Response, error) {
	/* Converte os dados para JSON */
	jsonData, err := json.Marshal(data)

	/* Se ocorrer um erro ao converter para JSON */
	if err != nil {
		return nil, fmt.Errorf("erro ao converter dados para JSON: %w", err)
	}

	/* Loga os dados da solicitação para depuração */
	fmt.Printf("Enviando solicitação POST para %s com os dados: %s\n", url, string(jsonData))

	/* Cria uma solicitação POST */
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	/* Se ocorrer um erro ao criar a solicitação */
	if err != nil {
		return nil, fmt.Errorf("erro ao criar solicitação: %w", err)
	}

	/* Define o header como JSON */
	req.Header.Set("Content-Type", "application/json")

	/* Cria um cliente HTTP com o timeout especificado e transporte personalizado */
	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, /* Ignora erros de certificado SSL, para localhosts funcionarem */
			},
		},
	}

	/* Envia a solicitação */
	resp, err := client.Do(req)

	/* Se ocorrer um erro ao enviar a solicitação */
	if err != nil {
		return nil, fmt.Errorf("erro ao enviar solicitação: %w", err)
	}

	/* Verifica se houve erros HTTP */
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close() /* Fecha o body da resposta */
		fmt.Printf("Resposta de erro do servidor: %s\n", string(body))
		return nil, fmt.Errorf("erro de resposta do servidor: %s - %s", resp.Status, string(body))
	}

	/* Retorna a resposta */
	return resp, nil
}

/* SendRaw envia uma mensagem bruta via POST request e retorna a resposta do servidor */
func SendRaw(chatId string, msg map[string]interface{}, reply interface{}, code interface{}) (map[string]interface{}, error) {
	/* Se o código for nulo, define como falso */
	if code == nil {
		code = false
	}

	/* Prepara os dados do POST com os valores ajustados */
	postData := map[string]interface{}{
		"username": config.ConfigData.Auth.Username,
		"password": config.ConfigData.Auth.Password,
		"code":     code,
		"chatId":   chatId,
		"quoted":   reply,
		"message":  msg,
	}

	/* Obtém a URL do servidor */
	url := config.ConfigData.PostRequest.Value

	/* Envia a solicitação POST e obtém a resposta */
	resp, err := postMessage(url, postData, 10*time.Second)

	/* Se ocorrer um erro ao enviar a solicitação */
	if err != nil {
		return nil, err
	}

	/* Fecha o body */
	defer resp.Body.Close()

	/* Analisa o corpo da resposta */
	var responseData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return nil, fmt.Errorf("erro ao decodificar resposta: %w", err)
	}

	/* Retorna os dados da resposta */
	return responseData, nil
}

/* SendMessage envia uma mensagem via POST request e retorna a resposta do servidor */
func SendMessage(chatId string, msg map[string]interface{}, reply interface{}, code interface{}) (map[string]interface{}, error) {
	/* Envia a mensagem personalizada */
	return SendRaw(chatId, msg, reply, code)
}
