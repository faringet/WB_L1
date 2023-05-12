package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
собственное множество.
*/

func main() {
	// дано по условию
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// пустая мапа для множества
	ans := make(map[string]struct{})

	// каждая строку пишется в ans как ключи (так как ключи в мапе уникальны - решен вопрос с множеством)
	for _, str := range sequence {
		ans[str] = struct{}{} // пустая структура - заглушка для значений мапы
	}
	for str := range ans {
		fmt.Println(str)
	}
}
