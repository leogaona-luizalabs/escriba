package bot

import (
	"time"

	"github.com/luizalabs/escriba/services"
	"github.com/shomali11/slacker"
)

// Approve adiciona uma aprovação no draft
func (b *Bot) Approve(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	time.Sleep(1 * time.Second)

	url := request.StringParam("url", "")
	if url == "" {
		response.Reply("Me desculpe, não consegui capturar a URL do artigo que você deseja aprovar")
		return
	}

	err := b.service.Approve(url)
	if err != nil {
		switch err.(type) {
		case *services.NotFoundError:
			response.Reply("Me desculpe, não consegui encontrar o artigo com esta URL. Poderia checar e mandar o comando novamente?")
		default:
			response.Reply("Me desculpe, ocorreu um erro ao tentar processar a sua solicitação")
		}
		return
	}

	response.Reply("O artigo foi atualizado com sua aprovação")
}
