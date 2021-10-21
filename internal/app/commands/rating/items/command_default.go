package items

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Default обработка сообщений с нераспознанной командой
func (c *RatingItemsCommander) Default(inputMessage *tgbotapi.Message) {
	c.Answer(inputMessage, fmt.Sprintf(UnknownCmd, inputMessage.Text))
}
