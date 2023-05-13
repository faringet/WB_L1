package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке. Пример: «snow dog sun — sun dog snow».
*/

func main() {
	// дано по условию
	str := "now dog sun"

	// разделяем строку на слова
	sepStr := strings.Split(str, " ")

	// Свайп слов в массиве
	for i, j := 0, len(sepStr)-1; i < j; i, j = i+1, j-1 {
		sepStr[i], sepStr[j] = sepStr[j], sepStr[i]
	}

	// Склеиваем элементы массива в строку
	fmt.Println(strings.Join(sepStr, " "))
}
