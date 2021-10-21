package items

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Get обработка сообщений с запросом на выдачу объекта с указанным идентификатором
func (c *RatingItemsCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	//получаем идентификатор
	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		c.Answer(inputMessage, fmt.Sprintf(WrongArgs, args))
		return
	}
	//пытаемся получить объект
	entity, err := c.itemsService.Describe(idx)
	if err != nil {
		c.Answer(inputMessage, fmt.Sprintf(IdNotFound, args))
		return
	}

	//возвращаем сообщение с строковым представлением объекта
	c.Answer(inputMessage, entity.String())

}
