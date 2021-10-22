package items

import (
	"errors"
	"sync"

	"github.com/ozonmp/omp-bot/internal/model/rating"
)

const (
	OutOfBounds = "cursor out of bounds"
	NotFound    = "entity not found"
)

// ItemsService Интерфейс согласно заданию-1
type ItemsService interface {
	Describe(itemID uint64) (*rating.Item, error)
	List(cursor uint64, limit uint64) ([]rating.Item, error)
	Create(rating.Item) (uint64, error)
	Update(itemID uint64, item rating.Item) error
	Remove(itemID uint64) (bool, error)
}

// DummyItemsService макет сервиса
type DummyItemsService struct {
	lock      *sync.Mutex
	idCounter uint64
	entities  []rating.Item
}

// NewDummyItemsService конструкор с тесовыми данными
func NewDummyItemsService() *DummyItemsService {
	return &DummyItemsService{lock: &sync.Mutex{}, idCounter: 3, entities: []rating.Item{
		{ID: 1, Title: "First item"},
		{ID: 2, Title: "Second item"},
		{ID: 3, Title: "Third item"},
	}}
}

//функция возвращает индекс элемента с указанным идентификатором
func indexOfItem(id uint64, sl []rating.Item) (int, error) {
	for res, item := range sl {
		if item.ID == id {
			return res, nil
		}
	}

	return 0, errors.New(NotFound)
}

// List метод интерфейса выборки слайса объектов с текущей cursor позиции по limit элементов
func (s *DummyItemsService) List(cursor uint64, limit uint64) ([]rating.Item, error) {

	s.lock.Lock()
	defer s.lock.Unlock()

	//если текущая позиция вышла за длину слайса
	if int(cursor) >= len(s.entities) {
		return nil, errors.New(OutOfBounds)
	}

	//вычисляем конечную позицию выборкиб ограничив максимальное знаение длиной слайса
	end := int(cursor + limit)
	if end >= len(s.entities) {
		end = len(s.entities)
	}

	return s.entities[cursor:end], nil
}

// Describe метод интерфейса возврата объекта по его идентификатору
func (s *DummyItemsService) Describe(itemID uint64) (*rating.Item, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	//ищем в слайсе индекс объекта с указанным идентификатором
	idx, err := indexOfItem(itemID, s.entities)

	if err != nil {
		return nil, err
	}

	return &s.entities[idx], nil
}

// Create метод интерфейса создания нового объекта
func (s *DummyItemsService) Create(entity rating.Item) (uint64, error) {

	s.lock.Lock()
	defer s.lock.Unlock()

	//генерируем очередной идентификатор
	s.idCounter++

	//присваиваем новый идентификатор и добавляем в слайс
	entity.ID = s.idCounter
	s.entities = append(s.entities, entity)

	return entity.ID, nil
}

// Update метод интерфейса обновления данных указанного объекта
func (s *DummyItemsService) Update(itemID uint64, item rating.Item) error {

	s.lock.Lock()
	defer s.lock.Unlock()

	//ищем в слайсе индекс объекта с указанным идентификатором
	idx, err := indexOfItem(itemID, s.entities)
	if err != nil {
		return err
	}

	//Обновляем данные, оставив прежним идентификатор
	s.entities[idx] = item
	s.entities[idx].ID = itemID

	return nil
}

//метод интерфейса удаления объекта с заданным идентификатором
func (s *DummyItemsService) Remove(itemID uint64) (bool, error) {

	s.lock.Lock()
	defer s.lock.Unlock()

	//ищем в слайсе индекс объекта с указанным идентификатором
	idx, err := indexOfItem(itemID, s.entities)

	if err != nil {
		return false, err
	}

	//создаем новый слайс без выбранного индекса
	s.entities = append(s.entities[:idx], s.entities[idx+1:]...)

	return true, nil
}
