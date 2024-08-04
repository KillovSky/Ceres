/* Pacote para impressão de mensagens coloridas no terminal */
package color

/* Importa os pacotes necessários */
import (
	"fmt"
	"log"
	"os"
)

/* Mapa de cores ANSI */
var COLORS = map[string][2]string{
	"reset":           {"\033[0m", "\033[0m"},
	"bold":            {"\033[1m", "\033[22m"},
	"dim":             {"\033[2m", "\033[22m"},
	"italic":          {"\033[3m", "\033[23m"},
	"underline":       {"\033[4m", "\033[24m"},
	"inverse":         {"\033[7m", "\033[27m"},
	"hidden":          {"\033[8m", "\033[28m"},
	"strikethrough":   {"\033[9m", "\033[29m"},
	"black":           {"\033[30m", "\033[39m"},
	"red":             {"\033[31m", "\033[39m"},
	"green":           {"\033[32m", "\033[39m"},
	"yellow":          {"\033[33m", "\033[39m"},
	"blue":            {"\033[34m", "\033[39m"},
	"magenta":         {"\033[35m", "\033[39m"},
	"cyan":            {"\033[36m", "\033[39m"},
	"white":           {"\033[37m", "\033[39m"},
	"gray":            {"\033[90m", "\033[39m"},
	"grey":            {"\033[90m", "\033[39m"},
	"brightRed":       {"\033[91m", "\033[39m"},
	"brightGreen":     {"\033[92m", "\033[39m"},
	"brightYellow":    {"\033[93m", "\033[39m"},
	"brightBlue":      {"\033[94m", "\033[39m"},
	"brightMagenta":   {"\033[95m", "\033[39m"},
	"brightCyan":      {"\033[96m", "\033[39m"},
	"brightWhite":     {"\033[97m", "\033[39m"},
	"bgBlack":         {"\033[40m", "\033[49m"},
	"bgRed":           {"\033[41m", "\033[49m"},
	"bgGreen":         {"\033[42m", "\033[49m"},
	"bgYellow":        {"\033[43m", "\033[49m"},
	"bgBlue":          {"\033[44m", "\033[49m"},
	"bgMagenta":       {"\033[45m", "\033[49m"},
	"bgCyan":          {"\033[46m", "\033[49m"},
	"bgWhite":         {"\033[47m", "\033[49m"},
	"bgGray":          {"\033[100m", "\033[49m"},
	"bgGrey":          {"\033[100m", "\033[49m"},
	"bgBrightRed":     {"\033[101m", "\033[49m"},
	"bgBrightGreen":   {"\033[102m", "\033[49m"},
	"bgBrightYellow":  {"\033[103m", "\033[49m"},
	"bgBrightBlue":    {"\033[104m", "\033[49m"},
	"bgBrightMagenta": {"\033[105m", "\033[49m"},
	"bgBrightCyan":    {"\033[106m", "\033[49m"},
	"bgBrightWhite":   {"\033[107m", "\033[49m"},
}

/* Função para imprimir uma mensagem no console com a cor especificada */
func ColorPrint(text string, color string) string {
	/* Mensagem padrão */
	defaultMessage := fmt.Sprintf(
		"\033[31m[%s]\033[39m → \033[33mThe operation cannot be completed because no text has been sent.\033[39m",
		os.Getenv("PWD"),
	)

	/* Adquire a cor ou usa uma padrão */
	selectedColor, exists := COLORS[color]
	if !exists {
		selectedColor = COLORS["green"]
	}

	/* Verifica se o texto é uma string não vazia */
	if len(text) > 0 {
		/* Define como o texto usado */
		return fmt.Sprintf("%s%s%s", selectedColor[0], text, selectedColor[1])
	}

	/* Caso contrário, faz um aviso de erro */
	log.Printf("A mensagem deve ser uma string. Mensagem recebida: %v", text)

	/* E retorna a mensagem padrão */
	return defaultMessage
}
