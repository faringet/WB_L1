package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.
Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func main() {
	fmt.Println(uniqChecker("abcd"))
	fmt.Println(uniqChecker("abCdefAaf"))
	fmt.Println(uniqChecker("aabcd"))
	fmt.Println(uniqChecker("aAcbfg"))
}

func uniqChecker(str string) (string, bool) {

	// переводим все в нижний регистр, чтобы быть регистронезависимой
	str = strings.ToLower(str)

	// классический двойной цикл для сравнения элементов
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {

			// есть совпадения - выбрасываем false
			if str[i] == str[j] {
				ans := str + " -"
				return ans, false
			}
		}
	}
	// если нет - все хорошо - true
	ans := str + " -"
	return ans, true
}
