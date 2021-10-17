package items

import (
	"errors"

	"github.com/ozonmp/omp-bot/internal/model/rating"
)

// ItemsService Интерфейс согласно заданию-1
type ItemsService interface {
	Describe(itemsID uint64) (*rating.Items, error)
	List(cursor uint64, limit uint64) ([]rating.Items, error)
	Create(rating.Items) (uint64, error)
	Update(itemsID uint64, items rating.Items) error
	Remove(itemsID uint64) (bool, error)
}

// DummyItemsService макет сервиса
type DummyItemsService struct {
	idCounter uint64
	entities  []rating.Items
}

// NewDummyItemsService конструкор с тесовыми данными
func NewDummyItemsService() *DummyItemsService {
	return &DummyItemsService{3, []rating.Items{
		{1, "First item"},
		{2, "Second item"},
		{3, "Third item"},
	}}
}

//функция возвращает индекс элемента с указанным идентификатором
func indexOfItems(id uint64, sl []rating.Items) (int, error) {
	for res, item := range sl {
		if item.ID == id {
			return res, nil
		}
	}

	return 0, errors.New("entity not found")
}

// List метод интерфейса выборки слайса объектов с текущей cursor позиции по limit элементов
func (s *DummyItemsService) List(cursor uint64, limit uint64) ([]rating.Items, error) {

	//если текущая позиция вышла за длину слайса
	if int(cursor) >= len(s.entities) {
		return nil, errors.New("cursor out of index")
	}

	//вычисляем конечную позицию выборкиб ограничив максимальное знаение длиной слайса
	end := int(cursor + limit)
	if end >= len(s.entities) {
		end = len(s.entities)
	}

	return s.entities[cursor:end], nil
}

// Describe метод интерфейса возврата объекта по его идентификатору
func (s *DummyItemsService) Describe(itemsID uint64) (*rating.Items, error) {

	//ищем в слайсе индекс объекта с указанным идентификатором
	idx, err := indexOfItems(itemsID, s.entities)

	if err != nil {
		return nil, err
	}

	return &s.entities[idx], nil
}

// Create метод интерфейса создания нового объекта
func (s *DummyItemsService) Create(entity rating.Items) (uint64, error) {

	//генерируем очередной идентификатор
	s.idCounter++

	//присваиваем новый идентификатор и добавляем в слайс
	entity.ID = s.idCounter
	s.entities = append(s.entities, entity)

	return entity.ID, nil
}

// Update метод интерфейса обновления данных указанного объекта
func (s *DummyItemsService) Update(itemsID uint64, items rating.Items) error {

	//ищем в слайсе индекс объекта с указанным идентификатором
	idx, err := indexOfItems(itemsID, s.entities)
	if err != nil {
		return err
	}

	//Обновляем данные, оставив прежним идентификатор
	s.entities[idx] = items
	s.entities[idx].ID = itemsID

	return nil
}

//метод интерфейса удаления объекта с заданным идентификатором
func (s *DummyItemsService) Remove(itemsID uint64) (bool, error) {

	//ищем в слайсе индекс объекта с указанным идентификатором
	idx, err := indexOfItems(itemsID, s.entities)

	if err != nil {
		return false, err
	}

	//создаем новый слайс без выбранного индекса
	s.entities = append(s.entities[:idx], s.entities[idx+1:]...)

	return true, nil
}
