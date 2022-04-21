package controller

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var mainMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🎵 Audio"),
		tgbotapi.NewKeyboardButton("🖼 Images"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("📹 Video"),
		tgbotapi.NewKeyboardButton("🗒 Document"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("📁 Send File"),
	),
)

func (c *Controller) ShowMenu(chatId int64) {
	msg := tgbotapi.NewMessage(chatId, "Main Menu")
	msg.ReplyMarkup = mainMenu
	c.bot.Send(msg)
}
