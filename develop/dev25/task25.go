package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep
*/

func sleep(seconds int) {
	// преобразуем инпут функции в тип time.Duration чтобы дальше скормить это тикеру
	duration := time.Duration(seconds) * time.Second

	// тикер содержит канал, который поставляет "тики" часов с интервалами
	ticker := time.NewTicker(duration)

	// останавливаем тикер при выходе из функции
	defer ticker.Stop()

	// блокируемся до получения значения и ждем указанное время (wantToSleep)
	<-ticker.C
}

func main() {
	wantToSleep := 5

	fmt.Printf("Старт программы! Будем спать %d секунд \n\n", wantToSleep)
	sleep(wantToSleep) // Вызов собственной функции sleep, ожидание 2 секунды
	fmt.Printf("Завершение программы после %d секунд сна", wantToSleep)
}
