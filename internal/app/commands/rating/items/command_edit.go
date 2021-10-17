package items

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/ozonmp/omp-bot/internal/model/rating"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Edit обработка сообщений редактирвания объекта (объект здается в формате JSON)
func (c *RatingItemsCommander) Edit(inputMessage *tgbotapi.Message) {
	//получаем параметры сообщения
	args := inputMessage.CommandArguments()
	args = strings.TrimSpace(args)

	//пробуем разобрать JSON-объект
	var entity rating.Items
	err := json.Unmarshal([]byte(args), &entity)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	//пытаемся обновить объект с указанным идентификатором
	err = c.itemsService.Update(entity.ID, entity)
	if err != nil {
		log.Printf("fail to update items with idx %d: %v", entity.ID, err)
		return
	}

	//возвращаем сообщение об успешном обновлении
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("%s edit ok", entity),
	)

	c.bot.Send(msg)
}
