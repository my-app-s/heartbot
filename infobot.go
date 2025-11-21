package heartbot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// SendAboutChat send in chat bot info.
func SendAboutChat(bot *tgbotapi.BotAPI, chatID int64) {
	// Realization
    msg := tgbotapi.NewMessage(chatID,
        "📦 *About the Bot / О боте:*\n" +
        "This bot is a personal pet project, created for **learning, experimentation, and demonstration of development skills**.\n" +
        "(_Этот бот — личный pet‑проект, созданный для обучения, экспериментов и демонстрации опыта разработки._)\n\n" +

        "🔐 *Access Notice / Уведомление о доступе:*\n" +
        "Access is restricted. If you obtained the token, link, or code unintentionally or without permission, please refrain from using it.\n" +
        "(_Доступ ограничен. Если вы получили токен, ссылку или код случайно или без разрешения — пожалуйста, не используйте его._)\n\n" +

        "⚠️ *Disclaimer / Отказ от ответственности:*\n" +
        "`THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT ANY WARRANTY OF ANY KIND, WHETHER EXPRESS OR IMPLIED.`\n" +
        "`IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY CLAIMS, DAMAGES, OR OTHER LIABILITY, WHETHER IN CONTRACT, TORT, OR OTHERWISE, ARISING FROM, OUT OF, OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.`")
    msg.ParseMode = "Markdown"
    bot.Send(msg)
}
