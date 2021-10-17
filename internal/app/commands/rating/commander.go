package rating

import (
	"log"

	service "github.com/ozonmp/omp-bot/internal/service/rating/items"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/rating/items"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

// Commander интерфейс обработчика команд
type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

// RatingCommander обработчик команд Rating
type RatingCommander struct {
	bot            *tgbotapi.BotAPI
	itemsCommander Commander //обработчик команд Items
}

// NewRatingCommander Конструктор обработчика команд Rating
func NewRatingCommander(bot *tgbotapi.BotAPI) *RatingCommander {
	//Создаем сервис Items
	itemsService := service.NewDummyItemsService()
	return &RatingCommander{
		bot:            bot,
		itemsCommander: items.NewItemsCommander(bot, itemsService),
	}
}

// HandleCallback метод обработки функций обратного вызова
func (c *RatingCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "items":
		c.itemsCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("RatingCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

// HandleCommand метод обработки команд
func (c *RatingCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "items":
		c.itemsCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("RatingCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
