package keyboard

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CmdKeyboard interface {
	hello()
	saybye()
}

type cmdKeyboard struct {
	cmdKeyboard *tgbotapi.ReplyKeyboardMarkup
}

func (k *cmdKeyboard) hello() {
	print("Hello !")
}

func (k *cmdKeyboard) saybye() {
	print("Bye Bye !")
}

var cmdKey *cmdKeyboard

func New() CmdKeyboard {
	var cmdKeyboards = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/set_todo"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/delete_todo"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/show_all_todos"),
		),
	)

	cmdKey = &cmdKeyboard{
		cmdKeyboard: &cmdKeyboards,
	}

	return cmdKey
}