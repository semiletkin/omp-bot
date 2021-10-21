package rating

import "fmt"

type Item struct {
	ID    uint64 //автоинкрементный идентификатор
	Title string //полезная нагрузка
}

//метод интерфейса Stringer
func (i Item) String() string {
	return fmt.Sprintf("Item (%#v)", i)
}
