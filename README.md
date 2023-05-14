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



