package rating

import "fmt"

type Items struct {
	ID    uint64
	Title string
}

func (i Items) String() string {
	return fmt.Sprintf("Item (%#v)", i)
}
