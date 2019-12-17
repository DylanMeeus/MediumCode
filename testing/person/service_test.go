package person

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	n := Service{repo: inmemory{}}
	richard := Person{
		Name:     "richard",
		Lastname: "feynman",
		Age:      51,
	}
	n.createPerson(richard)
	p := n.lookup("richard")
	if p != richard {
		fmt.Errorf("%v does not equal %v\n", p, richard)
	}
}
