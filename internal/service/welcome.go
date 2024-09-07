package service

// import (
// 	"artificial-dialogue/pkg/keyboard"
// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// )

// func Start(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
// 	text := "Hi, here you can create todos for your todolist."
// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
// 	msg.ReplyMarkup = keyboard.New()
// 	if _, err := bot.Send(msg); err != nil {
// 		panic(err)
// 	}
// }