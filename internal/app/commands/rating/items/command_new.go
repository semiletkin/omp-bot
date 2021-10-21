package items

import (
	"fmt"
	"strings"

	"github.com/ozonmp/omp-bot/internal/model/rating"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// New обработка сообщений с созданием нового объекта
func (c *RatingItemsCommander) New(inputMessage *tgbotapi.Message) {

	//получаем аргумент команды
	args := inputMessage.CommandArguments()
	args = strings.TrimSpace(args)
	entity := rating.Item{Title: args}

	//пытаемся создать новый объект
	id, err := c.itemsService.Create(entity)
	if err != nil {
		c.Answer(inputMessage, fmt.Sprintf(CreateError, err))
		return
	}

	//присваиваем назначенный сервисом идентификатор
	entity.ID = id
	//возвращаем сообщение об успешном создании объекта

	c.Answer(inputMessage, fmt.Sprintf(CreateOk, entity))
}
