package api

import (
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/asadbekGo/telegram-message-service/api/clienthttp"
	"github.com/asadbekGo/telegram-message-service/api/controller"
	"github.com/asadbekGo/telegram-message-service/config"
	l "github.com/asadbekGo/telegram-message-service/pkg/logger"
	"github.com/asadbekGo/telegram-message-service/storage"
)

// Option ...
type Option struct {
	Conf    config.Config
	Log     l.Logger
	Bot     *tgbotapi.BotAPI
	Updates tgbotapi.UpdatesChannel
	Storage storage.IStorage
}

// NewRouter
func NewRouter(option Option) {
	clientHttp, err := clienthttp.New(&clienthttp.ClientAPI{
		HttpClient: http.Client{},
	})
	if err != nil {
		option.Log.Error("failed to connection client http", l.Error(err))
		return
	}

	c := controller.NewController(
		option.Log,
		option.Conf,
		option.Bot,
		option.Storage,
		clientHttp,
	)

	for update := range option.Updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				cmdText := update.Message.Command()
				if cmdText == "menu" {
					c.ShowMenu(update.Message.Chat.ID)
				}
			}

			if update.Message.Text == "🎵 Audio" {
				c.GetAudio(update.Message.From.ID)
			} else if update.Message.Text == "🖼 Images" {
				c.GetPhoto(update.Message.From.ID)
			} else if update.Message.Text == "📹 Video" {
				c.GetVideo(update.Message.From.ID)
			} else if update.Message.Text == "🗒 Document" {
				c.GetFile(update.Message.From.ID)
			} else if update.Message.Text == "📁 Send File" {
				option.Bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "send document"))
			}

			if update.Message.Audio != nil {
				c.ReceiveAudio(update.Message.Chat.ID, update.Message)
			} else if update.Message.Photo != nil {
				c.ReceivePhoto(update.Message.Chat.ID, update.Message)
			} else if update.Message.Video != nil {
				c.ReceiveVideo(update.Message.Chat.ID, update.Message)
			} else if update.Message.Document != nil {
				c.ReceiveFile(update.Message.Chat.ID, update.Message)
			}
		}
	}

}
