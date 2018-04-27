package bot

import (
	"time"

	"github.com/shomali11/slacker"
)

// Default handler padrão, executado quando um comando não foi reconhecido
func (bot *Bot) Default(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	time.Sleep(1 * time.Second)
	response.Reply("Me desculpe jovem, não entendi o que você quer dizer")
}
