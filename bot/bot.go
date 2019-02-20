package bot

import (
	"github.com/luizalabs/escriba/services/draft"
	"github.com/luizalabs/escriba/util"
	"github.com/nlopes/slack"
	"github.com/shomali11/slacker"
	"github.com/sirupsen/logrus"
)

// Bot struct
type Bot struct {
	slackClient *slacker.Slacker
	idGen       slack.IDGenerator
	service     draft.ServiceIface
}

// Start inicia o bot, registra os handlers dos comandos e começa a escutar mensagens no slack
func Start(slackToken string, mysqlDSN string, draftApprovals int) {
	logger := util.GetLogger().WithFields(logrus.Fields{
		"module":         "bot",
		"operation_name": "start",
	})

	// cria o client do slacker e a conexão com o mysql
	slackClient := slacker.NewClient(slackToken)
	db, err := util.OpenMySQLConnection(mysqlDSN)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"error":       err.Error(),
			"error_stage": "openMySQLConnection",
		}).Fatal()
		return
	}

	bot := Bot{
		slackClient: slackClient,
		idGen:       slack.NewSafeID(1),
		service:     draft.New(db, draftApprovals),
	}

	// registra os comandos do bot
	slackClient.DefaultCommand(bot.Default)
	slackClient.Command("hello", "healthcheck do bot", bot.HealthCheck)
	slackClient.Command("pending reviews", "lista artigos pendentes de revisão", bot.ListPendingReviews)
	slackClient.Command("pending publications", "lista artigos pendentes de publicação", bot.ListPendingPublications)
	slackClient.Command("add <url>", "adiciona artigo", bot.Add)
	slackClient.Command("approve <url>", "aprova um artigo", bot.Approve)
	slackClient.Command("remove <url>", "remove um artigo que já foi publicado", bot.Remove)

	logger.Info()

	err = slackClient.Listen()
	if err != nil {
		logger.Fatal(err)
		return
	}
}
