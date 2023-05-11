package main

import "fmt"

/*
Дана переменная int64. Разработать программу, которая устанавливает i-й бит в
1 или 0.
*/

func main() {

	// number - дана по условию
	number := int64(55)

	temp := setBitToOne(number, 1)
	temp2 := setBitToZero(number, 3)

	fmt.Printf("Первоначальная переменная %d [%b]\n", number, number)
	fmt.Printf("%b\n", temp)
	fmt.Printf("%b\n", temp2)

}

// Маска с единицей в позиции pos + побитовое ИЛИ с n
func setBitToOne(n int64, pos uint) int64 {
	return n | 1<<(pos-1)

}

// Маска с единицей в позиции pos + ее инвертирование + побитовое и с n
func setBitToZero(n int64, pos uint) int64 {
	return n & ^(1 << (pos - 1))
}
