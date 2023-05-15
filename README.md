# WB_L1

## Теоречиские вопросы

### 1. Какой самый эффективный способ конкатенации строк?
Строки в Go — это иммутабельные массивы байт.
При конкатенации двух строк происходит выделение новой памяти.
Для простых случаев, когда производительность не является проблемой возможно использовать fmt.Sprintf или "+"

```go
s1 := fmt.Sprintf("Size: %d MB.", 85)
s2 := "Size: " + "85" + " MB."
```
output:
```
Size: 85 MB.
```

Самый эффективный способ конкатенации строк - через strings.Builder

`strings.Builder` минимизирует количество аллокаций памяти при построении строк. В отличие от обычной конкатенации, которая создает новую строку на каждой итерации, strings.Builder выделяет память заранее с определенным буфером и автоматически увеличивает его размер при необходимости. Это снижает нагрузку на сборщик мусора и улучшает производительность.
```go
var str strings.Builder
	for i := 0; i < 5; i++ {
		str.WriteString("sb demo ")
	}
	fmt.Println(str.String())
```
output:
```
sb demo sb demo sb demo sb demo sb demo 
```


### 3. Чем отличаются RWMutex от Mutex?
Mutex означает mutual exclusion(взаимное исключение) и является способом защиты critical section(критическая секция) программы.
В качестве примера приведу функцию из [**18 задания**](https://github.com/faringet/WB_L1/blob/master/develop/dev18/task18.go):
```go
// Increment Инкриминируем значение счетчика
func (c *Counter) Increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}
```

RWMutex концептуально то же самое, что и Mutex: он защищает доступ к памяти. Тем не менее, RWMutex дает немного больше контроля над памятью. Мы можем запросить блокировку для чтения, и в этом случае будет предоставлен доступ, если блокировка не удерживается для записи.

Как пример:
```go
func (c *counter) CountV1() int {
   c.Lock()
   defer c.Unlock()
   return c.count
}
func (c *counter) CountV2() int {
   c.RLock()
   defer c.RUnlock()
   return c.count
}
```
CountV2 не блокирует count если не было блокировок на запись.


### 4. Чем отличаются буферизированные и не буферизированные каналы?
Буферизированные каналы:

* имеют встроенный буфер, который временно хранит данные
* отправитель может записывать данные в буферизированный канал, даже если получатель ещё не готов принять их
* если буфер канала заполнен, то отправитель будет заблокирован до освобождения места в буфере
* получатель может прочитать данные из буфера канала, даже если отправитель ещё не записал новые данные
* если буфер канала пуст, то получатель будет заблокирован до появления данных в буфере

Не буферизированные каналы:

* сам go не аллоцирует буфер, поэтому в структуре hchan параметр buf будет пустым, как и размерность буфера(dataqsiz)
* при передаче данных отправитель и получатель должны быть готовы к операции чтения и записи соответственно
* если отправитель пытается записать данные, а получатель не готов прочитать их, то отправитель будет заблокирован до тех пор, пока получатель не будет готов принять данные
* если получатель пытается прочитать данные, а отправитель ещё не записал данные, получатель будет заблокирован до тех пор, пока данные не будут доступны

Наглядно посмотрим на структуру канала:

```go
ch := make(chan int, 2)
ch <- 42
// заполненная структура hchan
ch = {chan int} 
 qcount = {uint} 1
 dataqsiz = {uint} 2
 *buf = {*[2]int} len:2
 elemsize = {uint16} 8
 closed = {uint32} 0
 *elemtype = {*runtime._type} 
 sendx = {uint} 1
 recvx = {uint} 0
 recvq = {waitq<int>} 
 sendq = {waitq<int>} 
 lock = {runtime.mutex}
 ```
 Интересная подробная статья на medium в тему:
 
 [**Часть 1**](https://medium.com/@victor_nerd/%D0%BF%D0%BE%D0%B4-%D0%BA%D0%B0%D0%BF%D0%BE%D1%82%D0%BE%D0%BC-golang-%D0%BA%D0%B0%D0%BA-%D1%80%D0%B0%D0%B1%D0%BE%D1%82%D0%B0%D1%8E%D1%82-%D0%BA%D0%B0%D0%BD%D0%B0%D0%BB%D1%8B-%D1%87%D0%B0%D1%81%D1%82%D1%8C-1-e1da9e3e104d)
 
 [**Часть 2**](https://medium.com/@victor_nerd/golang-channel-internal-part2-b4e37ad9a118)
 
 
 ### 5. Какой размер у структуры struct{}{}?
 Пустая структура struct{}{} не занимает память. Она не имеет полей, и поэтому ее размер равен нулю - 0 байт.
 
 
 ### 6. Есть ли в Go перегрузка методов или операторов?
 В Go отсутвует перегрузка методов и операторов.
 
 


