package main

import "fmt"

type Animal interface {
	Grow()
	Move(string) string
}

type Cat struct {
	Name     string
	Age      uint8
	Location string
}

func (cat *Cat) Grow() {
	cat.Age++
}

func (cat *Cat) Move(new string) string {
	old := cat.Location
	cat.Location = new
	return old
}

func main() {
	myCat := Cat{"Little C", 2, "In the house"}
	myCat.Grow()
	animal, ok := interface{}(&myCat).(Animal)//&取址　先转换为空{}　然后在断言是Animal
	fmt.Printf("%v, %v, now age is %d\n", ok, animal, myCat.Age)
}
