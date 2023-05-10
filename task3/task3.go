/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(2^2+3^2+4^2...) с использованием конкурентных вычислений.
Данное задание выполнено на основе task2
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Тут храним суммы квадратов
	var ans int

	arr := [5]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	wg.Add(len(arr))

	// Мьютекс для синхронизации доступа (к переменной ans)
	var mux sync.Mutex

	for _, num := range arr {
		go square(num, &wg, &mux, &ans)
	}

	wg.Wait()
	fmt.Println(ans)
}

func square(num int, wg *sync.WaitGroup, mux *sync.Mutex, ans *int) *int {
	mux.Lock()
	*ans += num * num
	mux.Unlock()
	wg.Done()
	return ans
}
