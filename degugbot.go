package heartbot

import (
    "strconv"
	"os"
    "log"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv" // For work getenv
)

// SetDebugBot set status debug from env.
func SetDebugBot(bot *tgbotapi.BotAPI) {
	// Realization
	err:=godotenv.Load() // Initialization env file
	if err!=nil { log.Fatal("[FATAL LocalHeartBot]Not .env file.") }

	status:=false // Default status false
	if	statusFromEnv:=os.Getenv("STATUS_DEBUG");statusFromEnv!="" {
		parsed,err:=strconv.ParseBool(statusFromEnv)
		if err!=nil {
			log.Printf("[WARNING SetDebugBot]Invalid STATUS_DEBUG value: %q, error: %v.",statusFromEnv,err)
		} else {
			status=parsed
			bot.Debug=status
			log.Printf("[INFO SetDebugBot]Status debug set: %t.",status)
		}
	} else {
		log.Printf("[INFO SetDebugBot]Status debug set default: %t.",status)
	}
}