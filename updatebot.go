package heartbot

import "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// CreateUpdatesChannel create channel updates.
func CreateUpdatesChannel(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	// Realization
	u:=tgbotapi.NewUpdate(0)
	u.Timeout=60
	updates:=bot.GetUpdatesChan(u)
	return updates
}