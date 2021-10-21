package items

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Help обработка сообщений на запрос справки
func (c *RatingItemsCommander) Help(inputMessage *tgbotapi.Message) {
	msg := "/help__rating__items - print list of commands\n" +
		"/get__rating__items - get a entity with ID\n" +
		"/list__rating__items - get a list of entities\n" +
		"/delete__rating__items - delete an existing entity with ID\n" +
		"/new__rating__items - create a new entity with Title\n" +
		"/edit__rating__items — edit a entity {\"ID:\":identificator, \"Title\":\"yourText\"}"

	c.Answer(inputMessage, msg)
}
