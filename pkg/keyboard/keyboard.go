package keyboard

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CmdKeyboard interface {
	TradeBaseKeyboard() CmdKeyboard;
}

type cmdKeyboard struct {
	cmdKeyboard *tgbotapi.ReplyKeyboardMarkup;
}

func (cmdk *cmdKeyboard) TradeBaseKeyboard() CmdKeyboard {
	var tmp tgbotapi.ReplyKeyboardMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/trade_stock"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/evaluation"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/watchlist"),
		),
	);

	return &cmdKeyboard{
		cmdKeyboard: &tmp,
	};
}