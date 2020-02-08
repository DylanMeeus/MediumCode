package main

import "fmt"

type A interface{
	DoX()
	DoY()
}

type B interface{
	DoX()
}

type AB interface {
	A
	B
}

type S struct {}
func (s S) DoX(){}
func (s S) DoY(){}

func main() {
	s := S{}
	fmt.Printf("%v\n",s)
}
