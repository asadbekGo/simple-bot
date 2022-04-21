package controller

import (
	"context"
	"fmt"

	"github.com/asadbekGo/telegram-message-service/models"
	l "github.com/asadbekGo/telegram-message-service/pkg/logger"
)

const (
	SendPhoto    = "sendPhoto?chat_id"
	SendVideo    = "sendVideo?chat_id"
	SendAudio    = "sendAudio?chat_id"
	SendDocument = "sendDocument?chat_id"
)

func (c *Controller) GetAudio(fromID int64) {
	fileIds, err := c.storage.Message().GetAudios(fromID)
	if err != nil {
		c.log.Error("failed get photos storage", l.Error(err))
		return
	}

	for _, id := range fileIds {
		var response models.SendFile
		err := c.client.Do(
			context.Background(),
			"GET",
			fmt.Sprintf("%s/bot%s/%s=%d&&audio=%s",
				c.conf.TelegramBot.Url, c.conf.TelegramBot.Token, SendAudio, fromID, id),
			[]byte{},
			&response,
		)
		if err != nil {
			c.log.Error("failed client get audio", l.Error(err))
			return
		}

		if !response.Ok {
			c.log.Info(fmt.Sprintf("faild client send audio: %v", response))
		}
	}
}

func (c *Controller) GetPhoto(fromID int64) {
	fileIds, err := c.storage.Message().GetPhotos(fromID)
	if err != nil {
		c.log.Error("failed get photos storage", l.Error(err))
		return
	}

	for _, id := range fileIds {
		var response models.SendFile
		err := c.client.Do(
			context.Background(),
			"GET",
			fmt.Sprintf("%s/bot%s/%s=%d&&photo=%s",
				c.conf.TelegramBot.Url, c.conf.TelegramBot.Token, SendPhoto, fromID, id),
			[]byte{},
			&response,
		)
		if err != nil {
			c.log.Error("failed client get photo", l.Error(err))
			return
		}

		if !response.Ok {
			c.log.Info(fmt.Sprintf("faild client send photo: %v", response))
		}
	}
}

func (c *Controller) GetVideo(fromID int64) {
	fileIds, err := c.storage.Message().GetVideos(fromID)
	if err != nil {
		c.log.Error("failed get photos storage", l.Error(err))
		return
	}

	for _, id := range fileIds {
		var response models.SendFile
		err := c.client.Do(
			context.Background(),
			"GET",
			fmt.Sprintf("%s/bot%s/%s=%d&&video=%s",
				c.conf.TelegramBot.Url, c.conf.TelegramBot.Token, SendVideo, fromID, id),
			[]byte{},
			&response,
		)
		if err != nil {
			c.log.Error("failed client get video", l.Error(err))
			return
		}

		if !response.Ok {
			c.log.Info(fmt.Sprintf("faild client send video: %v", response))
		}
	}
}

func (c *Controller) GetFile(fromID int64) {
	fileIds, err := c.storage.Message().GetDocuments(fromID)
	if err != nil {
		c.log.Error("failed get photos storage", l.Error(err))
		return
	}

	for _, id := range fileIds {
		var response models.SendFile
		err := c.client.Do(
			context.Background(),
			"GET",
			fmt.Sprintf("%s/bot%s/%s=%d&&document=%s",
				c.conf.TelegramBot.Url, c.conf.TelegramBot.Token, SendDocument, fromID, id),
			[]byte{},
			&response,
		)
		if err != nil {
			c.log.Error("failed client get document", l.Error(err))
			return
		}

		if !response.Ok {
			c.log.Info(fmt.Sprintf("faild client send document: %v", response))
		}
	}
}
