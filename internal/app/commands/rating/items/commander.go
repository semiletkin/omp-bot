package items

import (
	"log"

	service "github.com/ozonmp/omp-bot/internal/service/rating/items"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const (
	HelpCmd   = "help"
	ListCmd   = "list"
	GetCmd    = "get"
	DeleteCmd = "delete"
	NewCmd    = "new"
	EditCmd   = "edit"
)

type ItemsCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)

	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type RatingItemsCommander struct {
	bot           *tgbotapi.BotAPI
	itemsService  service.ItemsService
	currentCursor uint64
	currentLimit  uint64
}

func NewItemsCommander(bot *tgbotapi.BotAPI, service service.ItemsService) ItemsCommander {

	return &RatingItemsCommander{
		bot:          bot,
		itemsService: service,
	}
}

func (c *RatingItemsCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case ListCmd:
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf(UnknownCmd, callbackPath.CallbackName)
	}
}

func (c *RatingItemsCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case HelpCmd:
		c.Help(msg)
	case ListCmd:
		c.List(msg)
	case GetCmd:
		c.Get(msg)
	case DeleteCmd:
		c.Delete(msg)
	case NewCmd:
		c.New(msg)
	case EditCmd:
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}
