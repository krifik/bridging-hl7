package bot

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/krifik/bridging-hl7/helper"
	"github.com/krifik/bridging-hl7/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/k0kubun/pp"
)

func StartBot() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			utils.SendMessage("LINE 23 \nLog Type: Error\n" + "Error: \n" + r.(error).Error() + "\n")
			wd, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				return
			}
			cmd := exec.Command("go", "run", "bot.go")
			cmd.Dir = wd
			output, err := cmd.Output()
			if err != nil {
				fmt.Println(err)
				fmt.Println(cmd)
			}
			fmt.Println(string(output))
		}
	}()
	godotenv.Load()
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		fmt.Println(err)
	}

	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = 30

	// Start polling Telegram for updates.

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	resDir := os.Getenv("RESULTDIR")
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		params := update.Message.CommandArguments()
		switch update.Message.Command() {
		case "createfile":
			msg.Text = "I understand /sayhi and /status."
			latestFileName := helper.GetLatestFileName(resDir)
			splittedStr := strings.Split(latestFileName, "-")
			var seq string
			if len(splittedStr) == 0 || len(splittedStr) < 1 {
				seq = ""
			} else {
				seq = splittedStr[1]
			}
			if seq != "" {
				seq = strings.Split(seq, ".")[0]
			}
			filteredStr := helper.RemoveAlphabet(seq)
			i, err := strconv.Atoi(filteredStr)
			if err != nil {
				utils.SendMessage("LINE 94 \nLog Type: Error\n" + "Error: \n" + err.Error() + "\n")
			}
			i += 1
			s := strconv.Itoa(i)
			msg.Text += " " + filteredStr
			o := resDir + time.Now().Format("20060102-"+s+".txt")
			f, err := os.Create(o)
			if err != nil {
				msg.Text = err.Error()
			}
			err = f.Close()
			if err != nil {
				utils.SendMessage("LINE 107 \nLog Type: Error\n" + "Error: \n" + err.Error() + "\n")
			}
			// os.OpenFile(resDir, os.O_RDONLY, 0644)
		case "files":
			entries, err := os.ReadDir(resDir)
			if err != nil {
				utils.SendMessage("LINE 113 \nLog Type: Error\n" + "Error: \n" + err.Error() + "\n")
			}
			var results []string
			for _, entri := range entries {
				results = append(results, entri.Name())
			}
			files := strings.Join(results[:], "\n")
			msg.Text = "List File: " + "\n" + files
		case "file":
			msg.Text = params
			fmt.Println(msg.Text)
			if msg.Text == "" {
				msg.Text = "Silahkan ketikan /file namafile.txt"
			} else {

				fileContent := helper.GetContent(resDir + "/" + params)
				j := utils.TransformToRightJson(fileContent)
				b, err := json.MarshalIndent(j, "", "    ")
				if err != nil {
					msg.Text = err.Error()
				}
				c, err := json.MarshalIndent(fileContent, "", "    ")
				if err != nil {
					msg.Text = err.Error()
				}
				msg.Text = "JSON: \n" + string(b)
				msg.Text += "\n\nFILE: \n" + string(c)
			}
		case "list":
			commands, err := bot.GetMyCommands()
			if err != nil {
				fmt.Println(err.Error())
			}
			var slCommands []string
			for _, c := range commands {
				slCommands = append(slCommands, c.Command+" - "+c.Description)
			}
			fmt.Println(slCommands)
			msg.Text = strings.Join(slCommands[:], "\n")
		case "status":
			msg.Text = "I'm ok."
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			pp.Print(err)
		}
	}
}
