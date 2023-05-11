package main

import (
	"fmt"
	"sync"
)

/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать
набор из N воркеров, которые читают произвольные данные из канала и
выводят в stdout. Необходима возможность выбора количества воркеров при
старте.
*/

func main() {
	// Переменная для хранения пользовательского ввода
	var numWorkers int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scanln(&numWorkers)

	// Создаем канал для передачи данных
	dataChan := make(chan int)

	// Создаем WaitGroup для ожидания завершения всех воркеров
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Запускаем воркеров
	for i := 0; i < numWorkers; i++ {
		go worker(i+1, dataChan, &wg)
	}

	// Отправляем данные (int от 1 до кол-ва, который юзер ввел) в канал
	for i := 1; i <= numWorkers; i++ {
		dataChan <- i
	}

	// Закрываем канал, чтобы воркеры могли завершиться
	close(dataChan)

	// Ожидаем завершения всех воркеров
	wg.Wait()

	fmt.Println("Главный поток завершился")
}

func worker(workerID int, dataChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Читаем данные из канала до его закрытия
	for data := range dataChan {
		fmt.Printf("Воркер %d: %d\n", workerID, data)
	}

	fmt.Printf("Воркер %d завершился\n", workerID)
}
