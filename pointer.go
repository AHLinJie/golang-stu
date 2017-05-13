package main

import "fmt"

type MyInt struct {
	n int
}

type MyIntInterface interface {
	Increase() int
	Decrease() int
}

func (myInt *MyInt) Increase() {
	//指针方法
	myInt.n++
}

func (myInt MyInt) Decrease() {
	//值方法
	myInt.n--
}


type Dog struct {
	name string
	age  uint8
}

type Pet interface {
	Name() string

	Age() uint8

	Grow()
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Age() uint8 {
	return dog.age
}

func (dog *Dog) Grow() {
	dog.age++
}

func main() {
	mi := MyInt{}
	fmt.Printf("%v\n", mi)
	mi.Decrease()
	mi.Increase()
	fmt.Printf("%v\n", mi.n == 1)
	myDog := Dog{"Little D", 3}
	pt1, ok1 := interface{}(&myDog).(Pet)
	pt2, ok2 := interface{}(myDog).(Pet)
	fmt.Printf("%v, %v\n", ok1, ok2)
	fmt.Printf("%d, %d\n", pt1.Age(), pt2.Age())
	fmt.Printf("%d, %d\n", pt1.age, pt2.Age())
}