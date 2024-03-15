package utils

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/k0kubun/pp"
)

// SendMessage sends a message to a Telegram group chat using Telegram Bot API.
//
// message: The message to be sent to the group chat.
func SendMessage(message string) {
	godotenv.Load()
	pp.Println(fmt.Sprintf("ERROR : %s", message))
	// bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	// if err != nil {
	// 	pp.Print(err.Error())
	// }

	// groupChatID, err := strconv.Atoi(os.Getenv("GROUPCHAT_ID"))
	// if err != nil {
	// 	pp.Print(err.Error())
	// }
	// msg := tgbotapi.NewMessage(int64(groupChatID), message)
	// if _, err := bot.Send(msg); err != nil {
	// 	pp.Print(err.Error())
	// }

}
