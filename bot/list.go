package bot

import (
	"fmt"
	"time"

	"github.com/ryanuber/columnize"
	"github.com/shomali11/slacker"
)

// ListPendingReviews ...
func (bot *Bot) ListPendingReviews(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	time.Sleep(1 * time.Second)

	items, err := bot.service.ListPendingReviews()
	if err != nil {
		response.Reply("Me desculpe, ocorreu um erro ao tentar processar o seu pedido")
		return
	}

	output := []string{"Aprovações | URL | Data de envio"}
	for i := range items {
		item := items[i]
		output = append(output, fmt.Sprintf("%d | %s | %s", item.Approvals, item.URL, item.CreatedAt))
	}

	response.Reply(fmt.Sprintf("Temos %d artigos pendentes de revisão: \n```%s```", len(items), columnize.SimpleFormat(output)))
	return
}

// ListPendingPublications ...
func (bot *Bot) ListPendingPublications(request *slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	time.Sleep(1 * time.Second)

	items, err := bot.service.ListPendingPublications()
	if err != nil {
		response.Reply("Me desculpe, ocorreu um erro ao tentar processar o seu pedido")
		return
	}

	output := []string{"Aprovações | URL | Data de envio"}
	for i := range items {
		item := items[i]
		output = append(output, fmt.Sprintf("%d | %s | %s", item.Approvals, item.URL, item.CreatedAt))
	}

	response.Reply(fmt.Sprintf("Temos %d artigos prontos para serem publicados no blog: \n```%s```", len(items), columnize.SimpleFormat(output)))
	return
}
