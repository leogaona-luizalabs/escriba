package bot

import (
	"time"

	"github.com/shomali11/slacker"
)

// Publish marca um artigo como publicado
func (b *Bot) Publish(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	time.Sleep(1 * time.Second)

	url := request.StringParam("url", "")
	if url == "" {
		response.Reply("Me desculpe, não consegui capturar a URL do artigo publicado")
		return
	}

	err := b.service.MarkAsPublished(url)
	if err != nil {
		response.Reply("Me desculpe, ocorreu um erro ao tentar processar o seu pedido")
		return
	}

	response.Reply("O artigo foi marcado como publicado e não aparecerá mais nas pesquisas")
}
