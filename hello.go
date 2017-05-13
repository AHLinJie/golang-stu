package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("hello, world\n")
	fmt.Println("NOW TIME IS :", time.ANSIC)
	fmt.Println("MY LOVE NUM IS", rand.Intn(10))
	fmt.Println("MY LOVE NUM IS", rand.Int())
	fmt.Printf("you hanve %g problems", math.Sqrt(7))
	var uni rune = 'a'
	var name = [...]int{1, 2, 3, 4, 5}
	var sl = name[1:4]
	sl = sl[:cap(sl)]
	var l = len(name)

	fmt.Print(len(name), len(sl), cap(sl), sl, l)
	fmt.Printf("字符 '%c' 的Unicode代码点是 %s。\n", uni, ("U+8D5E"))
	fmt.Printf("\n\n")
}
