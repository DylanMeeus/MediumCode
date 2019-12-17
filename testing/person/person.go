package person

import "fmt"

type Person struct {
	Name     string
	Lastname string
	Age      int
}

func (p Person) String() string {
	return fmt.Sprintf("%v %v %v", p.Name, p.Lastname, p.Age)
}
