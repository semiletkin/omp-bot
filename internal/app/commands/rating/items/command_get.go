package items

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Get обработка сообщений с запросом на выдачу объекта с указанным идентификатором
func (c *RatingItemsCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	//получаем идентификатор
	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Println("RatingItemsCommander.Get wrong args", args)
		return
	}
	//пытаемся получить объект
	entity, err := c.itemsService.Describe(idx)
	if err != nil {
		log.Printf("RatingItemsCommander.Get fail to get product with idx %d: %v", idx, err)
		return
	}
	//возвращаем сообщение с строковым представлением объекта
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("%s", entity),
	)

	c.bot.Send(msg)
}
