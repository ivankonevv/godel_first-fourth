package main

import (
	"first_test/first"
	"first_test/fourth"
	"first_test/second"
	"first_test/third"
	"fmt"
)

func main() {
	fmt.Println("-----1 задание-----")
	first.First()

	fmt.Println("\n-----2 задание-----")
	second.Second()

	fmt.Println("\n-----3 задание-----")
	third.Third()

	fmt.Println("\n-----4 задание-----")
	fourth.ArrOperations()
}
