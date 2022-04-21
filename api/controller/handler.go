package controller

import (
	"github.com/asadbekGo/telegram-message-service/api/clienthttp"
	"github.com/asadbekGo/telegram-message-service/config"
	l "github.com/asadbekGo/telegram-message-service/pkg/logger"
	"github.com/asadbekGo/telegram-message-service/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Controller ...
type Controller struct {
	log     l.Logger
	conf    config.Config
	storage storage.IStorage
	bot     *tgbotapi.BotAPI
	client  *clienthttp.ClientAPI
}

// NewController ...
func NewController(
	log l.Logger,
	conf config.Config,
	bot *tgbotapi.BotAPI,
	storage storage.IStorage,
	clientHttp *clienthttp.ClientAPI,
) Controller {
	return Controller{
		bot:     bot,
		log:     log,
		conf:    conf,
		storage: storage,
		client:  clientHttp,
	}
}
