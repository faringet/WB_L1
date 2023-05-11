package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры
Human (аналог наследования).
*/

type Human struct {
	Age  int
	Name string
	Sex  string
}

// Метод структуры Human
func (h *Human) getHumanAge() {
	fmt.Println(h.Age)
}

// Метод структуры Human
func (h *Human) getHumanName() {
	fmt.Println(h.Name)
}

// Метод структуры Human
func (h *Human) getHumanSex() {
	fmt.Println(h.Sex)
}

// Print Метод структуры Human
func (h *Human) Print() {
	fmt.Println("parent")
}

// Встраивание Human в Action
type Action struct {
	Human
}

// Print Метод структуры Action
func (a *Action) Print() {
	fmt.Println("child")
}

func main() {
	// Переменные Боб и Алиса типа Action
	Bob := Action{Human{Age: 30, Name: "Bob", Sex: "male"}}
	Alice := Action{Human{Age: 25, Name: "Alice", Sex: "female"}}

	// Метод структуры Human работает на переменной типа Action
	Bob.getHumanAge()
	Bob.getHumanName()
	Bob.getHumanSex()

	// Метод структуры Human работает на переменной типа Action
	Alice.getHumanAge()
	Alice.getHumanName()
	Alice.getHumanSex()

	// + бонус - Что будет, если и в родительской и дочерней структуре есть реализация методов с одинаковым названием?
	var x Action
	// Реализация родительского метода будет переписана реализацией дочернего метода
	x.Print()

}
