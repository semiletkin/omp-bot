package items

import (
	"fmt"
	"log"
	"strings"

	"github.com/ozonmp/omp-bot/internal/model/rating"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// New обработка сообщений с созданием нового объекта
func (c *RatingItemsCommander) New(inputMessage *tgbotapi.Message) {

	//получаем аргумент команды
	args := inputMessage.CommandArguments()
	args = strings.TrimSpace(args)
	entity := rating.Items{Title: args}

	//пытаемся создать новый объект
	id, err := c.itemsService.Create(entity)
	if err != nil {
		log.Printf("fail to create product with idx %d: %v", id, err)
		return
	}

	//присваиваем назначенный сервисом идентификатор
	entity.ID = id
	//возвращаем сообщение об успешном создании объекта
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf(" %s create ok", entity),
	)

	c.bot.Send(msg)
}
