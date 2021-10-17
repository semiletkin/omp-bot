package items

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Delete обработка сообщений удаления объекта с заданным идентификатором
func (c *RatingItemsCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	//получаем идентификатор
	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	//пробуем удалить
	var ok bool
	ok, err = c.itemsService.Remove(idx)
	if err != nil {
		log.Printf("fail to get product with idx %d: %v", idx, err)
		return
	}

	if !ok {
		log.Printf("fail to get product with idx %d: %v", idx, err)
		return
	}

	//возвращаем сообщение об успешном удалении
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Items ID "+args+" delete ok",
	)

	c.bot.Send(msg)
}
