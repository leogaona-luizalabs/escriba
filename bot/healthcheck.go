package bot

import (
	"github.com/shomali11/slacker"
)

// HealthCheck comando para checar se o bot está ativo ou não
func (b *Bot) HealthCheck(request *slacker.Request, response slacker.ResponseWriter) {
	response.Reply("Olá, tudo bem?")
}
