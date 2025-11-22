package heartbot

import (
	"os"
    "log"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv" // For work getenv
)

// GlobalHeartBot get bot by token from scope env.
func GlobalHeartBot() *tgbotapi.BotAPI{
	// Realization
	token:=os.Getenv("TOKEN_API")
	if token=="" { log.Fatal("[FATAL GlobalHeartBot]TOKEN is not found.") }

	bot, err:=tgbotapi.NewBotAPI(token)
	if err != nil { log.Panicf("[PANIC GlobalHeartBot]TOKEN error: %v.",err) }

	token="*****" // Token clear

	return bot
}

// LocalHeartBot get bot by token from local file env.
func LocalHeartBot() *tgbotapi.BotAPI{
	// Realization
	err:=godotenv.Load()
	if err!=nil { log.Fatal("[FATAL LocalHeartBot]Not .env file.") }

	token:=os.Getenv("TOKEN_API")
	if token=="" { log.Fatal("[FATAL LocalHeartBot]TOKEN is not found.") }

	bot, err:=tgbotapi.NewBotAPI(token)
	if err != nil { log.Panicf("[PANIC LocalHeartBot]TOKEN error: %v.",err) }

	token="*****" // Token clear

	return bot
}