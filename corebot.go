package heartbot

import (
	"os"
    "log"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv" // For work getenv
)

func init() {
	if FileExists(".env") {
		err:=godotenv.Load() // Initialization env file
		if err!=nil { log.Fatal("[FATAL Initialization package heartbot/corebot]Not .env file.") }
	}
}

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
	token:=os.Getenv("TOKEN_API")
	if token=="" { log.Fatal("[FATAL LocalHeartBot]TOKEN is not found.") }

	bot, err:=tgbotapi.NewBotAPI(token)
	if err != nil { log.Panicf("[PANIC LocalHeartBot]TOKEN error: %v.",err) }

	token="*****" // Token clear

	return bot
}