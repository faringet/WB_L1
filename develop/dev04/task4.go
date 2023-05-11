package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать
набор из N воркеров, которые читают произвольные данные из канала и
выводят в stdout. Необходима возможность выбора количества воркеров при
старте.
Программа должна завершаться по нажатию Ctrl+C.
*/

func main() {
	// Переменная для хранения пользовательского ввода
	var numWorkers int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scanln(&numWorkers)

	// Канал для передачи данных
	dataChan := make(chan int)

	// Канал для получения сигнала завершения
	doneChan := make(chan struct{})

	// WaitGroup для ожидания завершения всех воркеров
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Запуск воркеров
	for i := 0; i < numWorkers; i++ {
		go worker(i+1, dataChan, &wg, doneChan)
	}

	// Отправка данных (int от 1 до кол-ва, который юзер ввел) в канал
	for i := 1; i <= numWorkers; i++ {
		dataChan <- i
	}

	// Канал, который ловит сигналы ОС
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Ожидание когда юзер завершит программу (Ctrl+C)
	<-sigChan

	// Закрываем канал, чтобы воркеры могли завершиться
	close(dataChan)
	close(doneChan)

	// Ожидаем завершения всех воркеров
	wg.Wait()

	fmt.Println("main завершился")
}

func worker(workerID int, dataChan <-chan int, wg *sync.WaitGroup, doneChan <-chan struct{}) {
	defer wg.Done()

	// Чтение данных из канала до его закрытия || получения сигнала завершения
	for {
		select {
		case data, ok := <-dataChan:
			if !ok {
				// Канал данных закрыт, завершение работы
				fmt.Printf("Воркер %d завершился\n", workerID)
				return
			}
			fmt.Printf("Воркер %d: %d\n", workerID, data)

		case <-doneChan:
			// Получен сигнал завершения программы, завершаем работу воркера
			fmt.Printf("Воркер %d завершился\n", workerID)
			return
		}
	}
}
