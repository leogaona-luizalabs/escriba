package bot

import (
	"time"

	"github.com/labstack/gommon/log"
	"github.com/shomali11/slacker"
)

// Start ...
func Start(slackToken string) {

	bot := slacker.NewClient(slackToken)

	bot.DefaultCommand(func(request *slacker.Request, response slacker.ResponseWriter) {
		response.Typing()
		time.Sleep(3 * time.Second)
		response.Reply("Desculpe, não entendi o que você quis dizer")
	})

	bot.Command("oi", "checa se o escriba está acordado", HealthCheck)

	err := bot.Listen()
	if err != nil {
		log.Error(err)
	}
}
