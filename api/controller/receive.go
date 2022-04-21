package controller

import (
	"context"
	"fmt"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/asadbekGo/telegram-message-service/models"
	l "github.com/asadbekGo/telegram-message-service/pkg/logger"
	"github.com/asadbekGo/telegram-message-service/pkg/utils"
)

const (
	GetFile = "getFile?file_id="
)

func (c *Controller) ReceiveAudio(chatId int64, message *tgbotapi.Message) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.conf.Server.CtxDefaultTimeout))
	defer cancel()

	var response models.FileResponse
	err := c.client.Do(
		ctx,
		"GET",
		fmt.Sprintf("%s/bot%s/%s", c.conf.TelegramBot.Url, c.conf.TelegramBot.Token, GetFile+message.Audio.FileID),
		[]byte{},
		&response,
	)
	if err != nil {
		c.log.Error("failed client request", l.Error(err))
		return
	}

	err = utils.Write(message.Audio.FileName, "audio", response.Result.FilePath)
	if err != nil {
		c.log.Error("failed receive audio utils", l.Error(err))
	}

	err = c.storage.Message().CreateAudio(message, response.Result.FilePath)
	if err != nil {
		c.log.Error("failed crated audio", l.Error(err))
	}

	c.bot.Send(tgbotapi.NewMessage(chatId, "OK"))
}

func (c *Controller) ReceivePhoto(chatId int64, message *tgbotapi.Message) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.conf.Server.CtxDefaultTimeout))
	defer cancel()

	if len(message.Photo) == 0 {
		c.log.Error("failed photo not found")
		return
	}

	var response models.FileResponse
	err := c.client.Do(
		ctx,
		"GET",
		fmt.Sprintf("%s/bot%s/%s", c.conf.TelegramBot.Url, c.conf.TelegramBot.Token, GetFile+message.Photo[len(message.Photo)-1].FileID),
		[]byte{},
		&response,
	)
	if err != nil {
		c.log.Error("failed client request", l.Error(err))
		return
	}

	filename := strings.Split(response.Result.FilePath, "/")[1]
	err = utils.Write(filename, "photo", response.Result.FilePath)
	if err != nil {
		c.log.Error("failed receive file utils", l.Error(err))
	}

	err = c.storage.Message().CreatePhoto(message, response.Result.FilePath)
	if err != nil {
		c.log.Error("failed crated photo", l.Error(err))
	}

	c.bot.Send(tgbotapi.NewMessage(chatId, "OK"))
}

func (c *Controller) ReceiveVideo(chatId int64, message *tgbotapi.Message) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.conf.Server.CtxDefaultTimeout))
	defer cancel()

	var response models.FileResponse
	err := c.client.Do(
		ctx,
		"GET",
		fmt.Sprintf("%s/bot%s/%s", c.conf.TelegramBot.Url, c.conf.TelegramBot.Token, GetFile+message.Video.FileID),
		[]byte{},
		&response,
	)
	if err != nil {
		c.log.Error("failed client request", l.Error(err))
		return
	}

	err = utils.Write(message.Video.FileName, "video", response.Result.FilePath)
	if err != nil {
		c.log.Error("failed receive video utils", l.Error(err))
	}

	err = c.storage.Message().CreateVideo(message, response.Result.FilePath)
	if err != nil {
		c.log.Error("failed crated video", l.Error(err))
	}

	c.bot.Send(tgbotapi.NewMessage(chatId, "OK"))
}

func (c *Controller) ReceiveFile(chatId int64, message *tgbotapi.Message) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.conf.Server.CtxDefaultTimeout))
	defer cancel()

	var response models.FileResponse
	err := c.client.Do(
		ctx,
		"GET",
		fmt.Sprintf("%s/bot%s/%s", c.conf.TelegramBot.Url, c.conf.TelegramBot.Token, GetFile+message.Document.FileID),
		[]byte{},
		&response,
	)
	if err != nil {
		c.log.Error("failed client request", l.Error(err))
		return
	}

	err = utils.Write(message.Document.FileName, "document", response.Result.FilePath)
	if err != nil {
		c.log.Error("failed receive file utils", l.Error(err))
	}

	err = c.storage.Message().CreateDocument(message, response.Result.FilePath)
	if err != nil {
		c.log.Error("failed crated document", l.Error(err))
	}

	c.bot.Send(tgbotapi.NewMessage(chatId, "OK"))
}
