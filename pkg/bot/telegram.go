package bot

import (
	"github.com/asadbekGo/telegram-message-service/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewConnectionBot(cfg *config.Config) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramBot.Token)
	if err != nil {
		return nil, err
	}

	bot.Debug = false

	return bot, nil
}
