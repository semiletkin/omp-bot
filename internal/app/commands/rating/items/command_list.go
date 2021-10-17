package items

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

// List обработка сообщений с запросом на выдачу списка объектов
func (c *RatingItemsCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the items: \n"

	//если это первый вывод - выводим по максимальное число)
	if c.currentLimit == 0 {
		c.currentLimit = 3
	}

	outputMsgText = outputMsgText + fmt.Sprintf("(Items on page: %d )\n\n", c.currentLimit)

	//формируем список для текущей позиции курсора и количества объектов на странице
	entities, err := c.itemsService.List(c.currentCursor, c.currentLimit)
	if err != nil {
		log.Println("RatingItemsCommander.List ", err, c.currentCursor, c.currentLimit)
		if err.Error() == "cursor out of index" {
			//если курсор вышел за границы - уменьшим его на виличину объектов на странице
			c.currentCursor = c.currentCursor - c.currentLimit
			//и сгенерируем список заново
			c.List(inputMessage)
		}
		return
	}

	//формируем итоговое сообщение
	for _, entity := range entities {
		outputMsgText += fmt.Sprintf("%s\n", entity)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	//генерируем клавиатуру для управлением постраничного вывода
	//данные для кнопки, задающий лимит объектов 1
	limit1, _ := json.Marshal(CallbackListData{
		Cursor: 0,
		Limit:  1,
	})
	//данные для кнопки, задающий лимит объектов 2
	limit2, _ := json.Marshal(CallbackListData{
		Cursor: 0,
		Limit:  2,
	})
	//данные для кнопки, задающий лимит объектов 3
	limit3, _ := json.Marshal(CallbackListData{
		Cursor: 0,
		Limit:  3,
	})
	//данные для кнопки следующей страницы
	nextPage, _ := json.Marshal(CallbackListData{
		Cursor: c.currentCursor + c.currentLimit,
		Limit:  c.currentLimit,
	})
	//данные для кнопки предыдущей страницы
	prevPage, _ := json.Marshal(CallbackListData{
		Cursor: c.currentCursor - c.currentLimit,
		Limit:  c.currentLimit,
	})

	//задаем обратный вызов для кнопки 1
	callbackPathLimit1 := path.CallbackPath{
		Domain:       "rating",
		Subdomain:    "items",
		CallbackName: "list",
		CallbackData: string(limit1),
	}
	//задаем обратный вызов для кнопки 2
	callbackPathLimit2 := path.CallbackPath{
		Domain:       "rating",
		Subdomain:    "items",
		CallbackName: "list",
		CallbackData: string(limit2),
	}
	//задаем обратный вызов для кнопки 3
	callbackPathLimit3 := path.CallbackPath{
		Domain:       "rating",
		Subdomain:    "items",
		CallbackName: "list",
		CallbackData: string(limit3),
	}
	//задаем обратный вызов для кнопки предыдущей страницы
	callbackPathPrevPage := path.CallbackPath{
		Domain:       "rating",
		Subdomain:    "items",
		CallbackName: "list",
		CallbackData: string(prevPage),
	}
	//задаем обратный вызов для кнопки следующей страницы
	callbackPathNextPage := path.CallbackPath{
		Domain:       "rating",
		Subdomain:    "items",
		CallbackName: "list",
		CallbackData: string(nextPage),
	}

	//генерируемстрочку с клавишами навигации
	rowNav := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Previous page", callbackPathPrevPage.String()),
		tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPathNextPage.String()),
	)

	//если на первой странице - доступна только клавиша следующей страницы
	if c.currentCursor == 0 {
		rowNav = rowNav[1:]
	}

	//генерируем финальную клавиатуру
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1", callbackPathLimit1.String()),
			tgbotapi.NewInlineKeyboardButtonData("2", callbackPathLimit2.String()),
			tgbotapi.NewInlineKeyboardButtonData("3", callbackPathLimit3.String()),
		),
		rowNav,
	)

	c.bot.Send(msg)
}
