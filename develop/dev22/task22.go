package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a, b, значение которых > 2^20.
*/

/*
Так как int64 представляет целое число от –9 223 372 036 854 775 808 до 9 223 372 036 854 775 807,
воспользуюсь пакетом math/big
*/

func main() {
	// 2^70 и 2^80, очевидно, не вмещаются в int64. Решение - представить такое большое число в виде строки
	a, _ := new(big.Int).SetString("1180591620717411303424", 10)    // 2^70
	b, _ := new(big.Int).SetString("1208925819614629174706176", 10) // 2^80

	multiplication := new(big.Int).Mul(a, b)
	division := new(big.Int).Quo(b, a)
	adding := new(big.Int).Add(a, b)
	subtraction := new(big.Int).Sub(a, b)

	fmt.Printf("Вводные данные:\n a=%d \n b=%d \n a*b=%d \n b/a=%d \n a+b=%d \n a-b=%d", a, b, multiplication, division, adding, subtraction)
}
