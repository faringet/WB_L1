package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое
значение счетчика.
*/

// Counter Сама структура-счетчик
type Counter struct {
	mu    sync.Mutex
	count int
}

// Increment Инкриминируем значение счетчика
func (c *Counter) Increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func main() {
	counter := Counter{}
	var wg sync.WaitGroup
	numWorkers := 999999

	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter.count)
}
