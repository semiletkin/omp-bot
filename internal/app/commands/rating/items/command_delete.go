package items

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Delete обработка сообщений удаления объекта с заданным идентификатором
func (c *RatingItemsCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	//получаем идентификатор
	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		c.Answer(inputMessage, fmt.Sprintf(WrongArgs, args))
		return
	}

	//пробуем удалить
	var ok bool
	ok, err = c.itemsService.Remove(idx)
	if err != nil || !ok {
		c.Answer(inputMessage, fmt.Sprintf(IdNotFound, args))
		return
	}

	//возвращаем сообщение об успешном удалении
	c.Answer(inputMessage, fmt.Sprintf(IdDeleteOk, args))

}
