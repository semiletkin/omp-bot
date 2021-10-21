package items

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RatingItemsCommander) Answer(inputMessage *tgbotapi.Message, msg string) {

	answer := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		msg,
	)

	c.bot.Send(answer)
}
