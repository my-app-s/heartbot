package heartbot

import (
	"os"
    "log"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv" // For work getenv
)

// GlobalHeartBot get bot by token without local file env.
func GlobalHeartBot() *tgbotapi.BotAPI{
	// Realization
	token:=os.Getenv("TOKEN_API")  // Get TOKEN_API_BotTG_GetVacanciesTBot
	if token=="" { log.Fatal("[FATAL GlobalHeartBot]TOKEN is not found.") }

	bot, err:=tgbotapi.NewBotAPI(token) // Create BotTG
	if err != nil { log.Panicf("[PANIC GlobalHeartBot]TOKEN error: %s.",err) }

	token="*****" // Token clear

	return bot
}

// LocalHeartBot get bot by token with local file env.
func LocalHeartBot() *tgbotapi.BotAPI{
	// Realization
	err:=godotenv.Load() // Initialization env local file
	if err!=nil { log.Fatal("[FATAL LocalHeartBot]Not .env file.") }

	token:=os.Getenv("TOKEN_API")  // Get TOKEN_API_BotTG_GetVacanciesTBot
	if token=="" { log.Fatal("[FATAL LocalHeartBot]TOKEN is not found.") }

	bot, err:=tgbotapi.NewBotAPI(token) // Create BotTG
	if err != nil { log.Panicf("[PANIC LocalHeartBot]TOKEN error: %s.",err) }

	token="*****" // Token clear

	return bot
}
