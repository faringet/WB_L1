package main

import (
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации.

var justString string - неуместное использование глобальной переменной
func someFunc() {
v := createHugeString(1 << 10) - потенциальное переполнение памяти,
								создается строка размером 1024 байт, если представить что это как-то попадет в цикл
								и будет там крутится - недалеко и до паники.
justString = v[:100]           - если строка v принимает в себя символы отличные от латиницы, то каждый символ уже
								будет занимать 2 байта, а это не то что мы ожидаем. Также в justString сохраняется
								только первые 100 символов. Остальные так и болтаются в памяти
								+ ко всему строки иммутабельны, а значит каждый раз создается новая.
}
func main() {
someFunc()
}
*/

func main() {
	justString := someFunc()
	fmt.Println(justString)

	fmt.Println("_______________")

	first100 := []rune(justString[:100]) // слайс рун поможет правильно считать все данные, даже если у нас будут символы, занимающие больше одного байта
	fmt.Println(string(first100))
}

func generateRandomChar() byte {
	rand.Seed(time.Now().UnixNano())

	randomInt := rand.Intn(256)
	return byte(randomInt)
}

func createHugeString(writer io.Writer) {
	for i := 0; i < 1024; i++ {
		writer.Write([]byte{generateRandomChar()}) // пишем рандомные символы в writer
	}
}

func someFunc() string {
	var builder strings.Builder
	builder.Grow(100) // назначаем начальную емкость builder

	createHugeString(&builder)

	return builder.String()
}
