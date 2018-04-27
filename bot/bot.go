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

	slackClient.Command("oi", "healthcheck", bot.HealthCheck)
	slackClient.DefaultCommand(bot.Default)

	logger.Info()

	err = slackClient.Listen()
	if err != nil {
		logger.Fatal(err)
		return
	}
}
