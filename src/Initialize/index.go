/* Nome do pacote atual */
package initialize

/* Importa os módulos necessários */
import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	cases "ceres/src/Commands/Cases"
	color "ceres/src/Functions/Color"
	config "ceres/src/Functions/Config"
	updater "ceres/src/Functions/Updater"

	"github.com/gorilla/websocket"
)

/* WebSocketConfig contém a configuração para a conexão WebSocket. */
type WebSocketConfig struct {
	URL      string
	Username string
	Password string
	Interval time.Duration
}

/* CreateWebSocketDialer cria e retorna um dialer WebSocket configurado. */
func CreateWebSocketDialer() *websocket.Dialer {
	return &websocket.Dialer{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		HandshakeTimeout: 60 * time.Second, /* Aumenta o tempo limite para 60 segundos */
	}
}

/* AddAuthHeaders adiciona cabeçalhos de autenticação básica. */
func AddAuthHeaders(username, password string) http.Header {
	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))
	headers := http.Header{}
	headers.Add("Authorization", authHeader)
	return headers
}

/* ConnectToWebSocket estabelece uma conexão com o WebSocket. */
func ConnectToWebSocket(dialer *websocket.Dialer, url string, headers http.Header) (*websocket.Conn, error) {
	for i := 0; i < 3; i++ { /* Tentar 3 vezes */
		conn, _, err := dialer.Dial(url, headers)
		/* Se houver erro ao conectar */
		if err != nil {
			log.Printf("Falha ao conectar ao WebSocket %s (tentativa %d): %v", url, i+1, err)
			time.Sleep(2 * time.Second) /* Esperar antes de tentar novamente */
			continue
		}

		/* Retorna a conexão se bem-sucedida */
		return conn, nil
	}

	/* Retorna erro se falhar após várias tentativas */
	return nil, fmt.Errorf("falha ao conectar ao WebSocket %s após várias tentativas", url)
}

/* ProcessMessages lê e processa mensagens do WebSocket. */
func ProcessMessages(conn *websocket.Conn, wg *sync.WaitGroup) {
	/* Define terminado */
	defer wg.Done()

	/* Roda em looping */
	for {
		messageType, msg, err := conn.ReadMessage()
		/* Se houver erro ao ler a mensagem */
		if err != nil {
			/* E for de conexão encerrada */
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				log.Println("Conexão fechada normalmente")
				return
			}

			/* For de JSON inválido */
			log.Printf("Erro ao ler a mensagem: %v", err)
			return
		}

		/* Variável para armazenar o JSON recebido */
		var data map[string]interface{}

		/* Se houver erro ao decodificar JSON */
		if err := json.Unmarshal(msg, &data); err != nil {
			log.Printf("Erro ao decodificar JSON: %v", err)
			continue
		}

		/* Exibir a mensagem se contiver "printerMessage" */
		if printerMessage, ok := data["printerMessage"].(string); ok {
			fmt.Printf("%s\n", printerMessage)
		}

		/* Encaminhar os dados JSON completos para handleCommand */
		_, err = cases.HandleCommand(data)

		/* Se houver erro ao processar o comando */
		if err != nil {
			log.Printf("Erro ao processar comando: %v", err)
		}

		/* Se a mensagem for de fechamento */
		if messageType == websocket.CloseMessage {
			fmt.Println("Conexão fechada pelo servidor")
			return
		}
	}
}

/* CheckUpdates verifica periodicamente se há atualizações. */
func CheckUpdates(interval time.Duration) {
	/* Executa em looping */
	for {
		/* Executa a busca por atualizações */
		updater.CheckUpdates(interval)

		/* Aguarda o tempo para buscar de novo */
		time.Sleep(interval)
	}
}

/* Run inicializa a aplicação, conectando-se ao WebSocket e iniciando as rotinas necessárias. */
func Run() {
	/* Carrega a configuração */
	err := config.LoadConfig("src/Settings/config.json")

	/* Se houver falha ao carregar a configuração */
	if err != nil {
		log.Fatalf("Falha ao carregar a configuração: %v", err)
	}

	/* Recupera os valores de configuração */
	cfg := config.GetConfig()
	webSocketURL := cfg.WebSocket.Value
	username := cfg.Auth.Username
	password := cfg.Auth.Password
	updateInterval := time.Duration(cfg.UpdateInterval.Value) * time.Second

	/* Cria um dialer WebSocket */
	dialer := CreateWebSocketDialer()

	/* Adiciona cabeçalhos de autenticação */
	headers := AddAuthHeaders(username, password)

	/* Conecta-se ao WebSocket */
	conn, err := ConnectToWebSocket(dialer, webSocketURL, headers)

	/* Se houver falha ao conectar */
	if err != nil {
		log.Fatal(err)
	}

	/* Fecha a conexão */
	defer conn.Close()

	/* Avisa que conectou */
	fmt.Println(
		color.ColorPrint("[START] ", "green"),
		color.ColorPrint("→ Tudo pronto para começar!", "yellow"),
	)

	/* Cria um WaitGroup para processar mensagens */
	var wg sync.WaitGroup
	wg.Add(1)
	go ProcessMessages(conn, &wg)

	/* Captura sinais de interrupção do sistema */
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	/* Verifica atualizações periodicamente */
	go CheckUpdates(updateInterval)

	/* Se receber pedido para desligar */
	<-sigChan
	fmt.Println("Interrupção recebida, encerrando...")
}
