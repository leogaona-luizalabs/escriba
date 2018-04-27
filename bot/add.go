package bot

import (
	"time"

	"github.com/shomali11/slacker"
)

// Add adiciona um draft no escriba
func (b *Bot) Add(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	time.Sleep(1 * time.Second)

	url := request.StringParam("url", "")
	if url == "" {
		response.Reply("Me desculpe, não consegui capturar a URL do seu artigo")
		return
	}

	err := b.service.Add(url)
	if err != nil {
		response.Reply("Me desculpe, ocorreu um erro ao tentar processar o seu pedido")
		return
	}

	response.Reply("Seu artigo foi adicionado em meus arquivos com sucesso. Muito obrigado por contribuir com o blog")
}
