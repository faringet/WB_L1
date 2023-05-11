package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

func main() {
	numWorkers := 5
	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	ans := make(map[string]string)

	// Запускаем горутины для записи данных
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {

			// Лочим мапу
			mutex.Lock()

			// Выполняем запись данных в map
			key := fmt.Sprintf("Key%d", workerID)
			value := fmt.Sprintf("Value%d", workerID)
			ans[key] = value

			// Разлочиваемся
			mutex.Unlock()

			// Уведомляем WaitGroup
			wg.Done()
		}(i)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим в stdout
	for key, value := range ans {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}
