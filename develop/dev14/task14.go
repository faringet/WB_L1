package main

import "fmt"

/*
Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.
*/

func main() {
	whatType(1)
	whatType("test")
	whatType(true)
	whatType(make(chan int))

	whatType(make(map[int]int64))

}

// Эту задачу отлично решает type switch statement:
func whatType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Int")
	case string:
		fmt.Println("String")
	case bool:
		fmt.Println("Boolean")
	case chan int:
		fmt.Println("Channel")

	default:
		fmt.Println("Another type")
	}
}
