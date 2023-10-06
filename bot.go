package main

import (
  "net/http"
  tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
  "log"
  "database/sql"
  _ "github.com/lib/pq"
  "github.com/joho/godotenv"
  "fmt"
)

type Bot struct {
  bot *tgbotapi.BotAPI
}

func main() {
  err := godotenv.Load("token.env")
  if err != nil {
    log.Fatalf("No .env file")
  }

  bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
  if err != nil {
    log.Fatal(err)
  }

  bot.Debug = true
  var b Bot = Bot{bot: bot}

  b.Start()
}

func (b *Bot) Start() {
  u := tgbotapi.NewUpdate(0)
  u.Timeout = 60

  updates, err := bot.GetUpdateChan(u)
	if err != nil {
		log.Fatal(err)
	}
  for update := range updates {
    if update.Message == nil {
      continue
    }

    log.Printf("[%s]: %s", update.Message.From.UserName, update.Message.Text)

    if update.Message.IsCommand() {
      err := bot.handleCommand()
      if err != nil {
        err2 := fmt.Errof("Command handle error: %w", err)
        log.Print("Error: ", err2)
      }
      continue
    } else {
      msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Use correct command")
      _, err := bot.Send(msg)
      if err != nil {
        log.Print("Error: ", err)
      }
    }
  }
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
  switch message.Command() {
    default:
      ans := "Incorrect comand"
  }
  msg := tgbotapi.NewMessage(message.Chat.ID, ans)
  b.bot.Send(msg)
}
