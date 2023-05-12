package main

import "fmt"

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
массива, во второй — результат операции x*2, после чего данные из второго
канала должны выводиться в stdout.
*/

func main() {

	// дано по условию
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	inputNumbers := make(chan int)    // в этот канал пишем значения из массива
	outputNumbersX2 := make(chan int) // а в этот - считанные данные * 2

	// пишем из массива
	go func() {
		for _, num := range input {
			inputNumbers <- num
		}
		close(inputNumbers)
	}()

	// читаем из inputNumbers и умножаем * 2
	go func() {
		for num := range inputNumbers {
			outputNumbersX2 <- num * 2
		}
		close(outputNumbersX2)
	}()

	// читаем из outputNumbersX2 и выводим в stdout
	for result := range outputNumbersX2 {
		fmt.Println(result)
	}
}
