package items

import (
	"encoding/json"
	"fmt"
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
	var entity rating.Item
	err := json.Unmarshal([]byte(args), &entity)
	if err != nil {
		c.Answer(inputMessage, fmt.Sprintf(JsonError, err))
		return
	}

	//пытаемся обновить объект с указанным идентификатором
	err = c.itemsService.Update(entity.ID, entity)
	if err != nil {
		c.Answer(inputMessage, fmt.Sprintf(IdNotFound, entity.ID))
		return
	}

	//возвращаем сообщение об успешном обновлении
	c.Answer(inputMessage, fmt.Sprintf(UpdateOk, entity))
}
