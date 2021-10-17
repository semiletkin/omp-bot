package items

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

// CallbackListData Структура данных для функций обратного вызова
type CallbackListData struct {
	Cursor uint64 `json:"cursor"` //текущая позиция курсора
	Limit  uint64 `json:"limit"`  //количество элементов для вывода на странице
}

func (c *RatingItemsCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {

	//разбиаем пришедшие данные
	parsedData := CallbackListData{}
	errJson := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if errJson != nil {
		return
	}

	//удаляем предыдущее сообщение, чтобы не засорять экран листингом объектов
	_, err := c.bot.DeleteMessage(tgbotapi.DeleteMessageConfig{ChatID: callback.Message.Chat.ID, MessageID: callback.Message.MessageID})
	if err != nil {
		return
	}

	//устанавливаем в обработчике команд текущее положение курсора и количество элементов в списке
	c.currentCursor = parsedData.Cursor
	c.currentLimit = parsedData.Limit

	//формируем новый список объектов
	c.List(callback.Message)
}
