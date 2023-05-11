package main

/*
Реализовать все возможные способы остановки выполнения горутины.
1) Через контекст (также можно через context.WithTimeout)
2) Через канал (можно писать как в примере - закрытие канала close(stop), можно stop <- true)
3) Через runtime.Goexit
4) Через обычный таймер
5) Через signal.Notify
6) просто закончить main
*/

// 1. Через контекст

//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//
//	go func() {
//		for {
//			select {
//			case <-ctx.Done():
//				return
//			default:
//				// Выполнение работы горутины
//				fmt.Println("горутина жива...")
//				time.Sleep(1 * time.Second)
//			}
//		}
//	}()
//
//	time.Sleep(5 * time.Second)
//	cancel() // Остановка горутины путем отмены контекста
//
//	fmt.Println("горутина остановлена")
//}

// 2. Через канал

//func main() {
//	stop := make(chan bool)
//
//	go func() {
//		for {
//			select {
//			case <-stop:
//				return
//			default:
//				// Выполнение работы горутины
//				fmt.Println("ex2 горутина жива...")
//				time.Sleep(1 * time.Second)
//			}
//		}
//	}()
//
//	time.Sleep(5 * time.Second)
//	close(stop) // закрыли канал - горутина остановлена
//
//	fmt.Println("ex2 горутина остановлена")
//}

// 3. Через runtime.Goexit()

//func main() {
//	go func() {
//		for {
//			// Выполнение работы горутины
//			fmt.Println("ex3 горутина жива......")
//			time.Sleep(2 * time.Second)
//			runtime.Goexit() // Остановка текущей горутины
//			fmt.Println("это сообщение горутина никогда не выведет")
//		}
//	}()
//
//	time.Sleep(2 * time.Second) // спим пару секунд чтобы горутина успела написать в консоль
//	fmt.Println("горутина остановлена")
//}

// 4. Через таймер

//func main() {
//	timer := time.NewTimer(5 * time.Second)
//
//	go func() {
//		for {
//			select {
//			case <-timer.C:
//				return
//			default:
//				// Выполнение работы горутины
//				fmt.Println("ex4 горутина жива...")
//				time.Sleep(1 * time.Second)
//			}
//		}
//	}()
//
//	<-timer.C // Ожидание истечения таймера
//
//	fmt.Println("горутина остановлена")
//}

// 5. Через signal.Notify

//func main() {
//	stop := make(chan os.Signal, 1)
//
//	done := make(chan bool)
//
//	// Уведомляем о получении сигналов SIGINT и SIGTERM
//	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
//
//	fmt.Println("Программа работает. Нажми Ctrl+C для остановки.")
//
//	go func() {
//		for {
//			select {
//			case <-stop:
//				fmt.Println("Сигнал получен, завершение горутины.")
//				done <- true
//				return
//			default:
//				// Выполнение работы горутины
//				fmt.Println("ex5 горутина жива...")
//				time.Sleep(1 * time.Second)
//			}
//		}
//	}()
//
//	// Сигнал из main
//	<-done
//
//	// Завершение программу
//	os.Exit(0)
//}
