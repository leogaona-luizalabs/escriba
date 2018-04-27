package bot

import (
	"time"

	"github.com/shomali11/slacker"
)

// HealthCheck comando para checar se o bot está ativo ou não
func (b *Bot) HealthCheck(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	time.Sleep(1 * time.Second)
	response.Reply("Olá, tudo bem?")
}
