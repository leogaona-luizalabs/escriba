package bot

import (
	"time"

	"github.com/luizalabs/escriba/services"
	"github.com/shomali11/slacker"
)

// Remove marca um artigo como publicado
func (b *Bot) Remove(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	time.Sleep(1 * time.Second)

	url := request.StringParam("url", "")
	if url == "" {
		response.Reply("Me desculpe, não consegui capturar a URL do artigo publicado")
		return
	}

	err := b.service.MarkAsPublished(url)
	if err != nil {
		switch err.(type) {
		case *services.NotFoundError:
			response.Reply("Me desculpe, não consegui encontrar o artigo com esta URL. Poderia checar e mandar o comando novamente?")
		default:
			response.Reply("Me desculpe, ocorreu um erro ao tentar processar a sua solicitação")
		}
		return
	}

	response.Reply("O artigo foi removido com sucesso")
}
