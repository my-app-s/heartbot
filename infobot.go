package heartbot

import (
	"fmt"
	"time"
)

const MSG = "📦 *About the Bot / О боте:*\n" +
	"This bot is a personal pet project, created for **learning, experimentation, and demonstration of development skills**.\n" +
	"(_Этот бот — личный pet‑проект, созданный для обучения, экспериментов и демонстрации опыта разработки._)\n\n" +

	"🔐 *Access Notice / Уведомление о доступе:*\n" +
	"Access is restricted. If you obtained the token, link, or code unintentionally or without permission, please refrain from using it.\n" +
	"(_Доступ ограничен. Если вы получили токен, ссылку или код случайно или без разрешения — пожалуйста, не используйте его._)\n\n" +

	"⚠️ *Disclaimer / Отказ от ответственности:*\n" +
	"`THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT ANY WARRANTY OF ANY KIND, WHETHER EXPRESS OR IMPLIED.`\n" +
	"`IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY CLAIMS, DAMAGES, OR OTHER LIABILITY, WHETHER IN CONTRACT, TORT, OR OTHERWISE, ARISING FROM, OUT OF, OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.`"

// WelcomeAboutByStart send in chat bot welcom and about with sleep time.
func WelcomeAboutMessage(bot *tgbotapi.BotAPI, chatID int64) {
	bot.Send(tgbotapi.NewMessage(chatID, fmt.Sprintf("Welcome.(Добро пожаловать.)")))
	msg := tgbotapi.NewMessage(chatID, MSG)
	msg.ParseMode = "Markdown"
	bot.Send(msg)
	bot.Send(tgbotapi.NewMessage(chatID, fmt.Sprintf("⏳ About info please wait.")))
	time.Sleep(10 * time.Second)
}

// SendAboutChat send in chat bot info.
func SendAboutMessage(bot *tgbotapi.BotAPI, chatID int64) {
	// Realization
	msg := tgbotapi.NewMessage(chatID, MSG)
	msg.ParseMode = "Markdown"
	bot.Send(msg)
}
