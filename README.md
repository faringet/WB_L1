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


### 2. Что такое интерфейсы, как они применяются в Go?
Великолепная [**статья с хабра**](https://habr.com/ru/articles/658623/), где, все рассказывается в деталях.


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
 
 
 ### 7. В какой последовательности будут выведены элементы map[int]int?
 ```go
m[0]=1
m[1]=124
m[2]=281
```
В Go порядок элементов в мапе не определен. При каждом запуске или при изменении данных в мапе может меняться порядок элементов в самой map[int]int.

Отвечая на вопрос выше - вывод будет рандомным. 
 
 
 ### 8. В чем разница make и new?
 С помощью make(T, args) создаются сложные типы данных (слайсы, мапы, каналы). Возвращает инициализированное (не обнуленное) значение типа T(не * T)
 Причиной различия является то, что эти три типа представляют под капотом ссылки на структуры данных, которые должны быть инициализированы перед   использованием.
 
 С помощью new выделяется память (создает неименованную переменную и возваращет указатель на ее значение)
 Синтакс использования:
 
 var a = new(T)
 
 Пример:
 ```go
p := new(int) // p has *int type now
fmt.Println(*p) // "0"
каждый вызов ‘new’ возвращает новый адрес

p := new(int)
q := new(int)
fmt.Println(p == q) // false
```


 ### 9. Сколько существует способов задать переменную типа slice или map?
 * слайс - 5 способов
 * мапа - 4 способа

 ```go
slice := make([]int, 0)	// начальная длина 0
slice := make([]int, 0, 5) // начальная длина 0, емкость 5
var slice []int	// пустой слайс, не указывается ни начальная длина, ни начальная емкость. Будет иметь значение nil
slice := []int{} // пустой слайс, будет иметь нулевую длину, но уже будет выделена память для потенциальных будущих элементов. slice не будет равен nil, он будет указывать на пустую область памяти.
sice := []int{1,2} // объявленный слайс будет содержать значения 1, 2
    
//4 способа задать map	
var m map[string]int // m будет объявлена, но не будет инициализирована. Переменная m будет иметь нулевое значение для типа map[string]int, то есть будет равна nil	
m := map[string]int{ // m будет иметь тип map[string]int. Ключами мапы будут строки "1" и "2", а значениями - 55 и 66.
    "1": 55,
    "2": 66,
}	
m := make(map[string]int) // m будет содержать пустую мапу, готовую для добавления пар ключ-значение.
m := make(map[string]int, 10) // то же самое что и не строчке выше, только будет установлена изначальная емкость - 10
```


 ### 10. Что выведет данная программа и почему?
 
  ```go
func update(p *int) {
b := 2
p = &b
}

func main() {
var (
a = 1
p = &a
) 
fmt.Println(*p)
update(p)
fmt.
  ```
  output:
   ```go
   1
   1
   ```
   Если мы хотим чтобы все работало, как я предполагаю, как задумывалось, необходимо:
   * изменить func update чтобы она возвращала обновленное значение
   * писать это значение в какую-то переменною в мейне и выводить ее на экран (в старом коде у нас переменные p разные и находятся в разных областях видимости)

 Обновленный код, который выводит 1, 2:
   ```go
 func update(p *int) *int {
	b := 2
	p = &b
	return p

}
func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	newP := update(p)
	fmt.Println(*newP)
}
```


 ### 11. Что выведет данная программа и почему?
 
   ```go
func main() {
wg := sync.WaitGroup{}
for i := 0; i < 5; i++ {
wg.Add(1)

go func(wg sync.WaitGroup, i int) {
fmt.Println(i)
wg.Done()
}(wg, i)
}

wg.Wait()
fmt.Println("exit")
}
  ```
  
  Тут у нас будет deadlock! Дело в том что go func принимает в себя WitGroup по значению. wg.Done() не может "достучаться" до глобальной WaitGroup, а будет уменьшать свою копию.
  
  Чтобы все исправить просто передадим WaitGroup по указателю:
  
   ```go
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(&wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
  ```
  
  
   ### 12. Что выведет данная программа и почему?
   ```go
func main() {
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n)
}
 ```
   output:
   ```go
   0
   ```
   Внутри if блока создается новая переменная. Переменная n, которая находится в цикле инкрементируется локально, а прграмма выводит на экран совсем другую n - у них разные области видимости.
   
   
   ### 13. Что выведет данная программа и почему?
   ```go
func someAction(v []int8, b int8) {
v[0] = 100
v = append(v, b)
}
func main() {
var a = []int8{1, 2, 3, 4, 5}
someAction(a, 6)
fmt.Println(a)
}
   ```
 
  output:
   ```go
   [100 2 3 4 5]
   ```
   
   Тут у нас слайс передается в функцию по значению, однако, сам слайс содержит указатель на массив, в котором хранятся сами данные. Это означает, что при передаче слайса в функцию копируется только указатель на массив, а не все элементы самого массива. В функции someAction слайс `v` изменяется, присваивая его первому элементу значение 100. Это означает, что первый элемент слйса `a` теперь также имеет значение 100, потому что `v` и `a` ссылаются на один и тот же слайс в памяти. 
   
   Фукция append должна добавить значение `b` в конец слайса `v`, но в этой функции `v` перезаписывается и ссылается на новый слайс, который уже не связан с `a`. Поэтому шестерку мы так и не увидим в выводе.
   
   
   ### 14. Что выведет данная программа и почему?
   ```go
func main() {
slice := []string{"a", "a"}
func(slice []string) {
slice = append(slice, "a")
slice[0] = "b"
slice[1] = "b"
fmt.Print(slice)
}(slice)
fmt.Print(slice)
}
   ```
   output:
   ```go
   [b b a][a a]
   ```
   
   Почему так ? 
   * В анонимной функции происходит:
   1. к слайсу, который объявлен в мейне, аппендим "a"
   
  ```
   [a a a]
   ```
   
   2. в этом слайсе присваиваем нулевому и первому элементу значение "b" и выводим это на экран:
   
   ```
   [b b a]
   ```
   
   * Анонимная функция завершает работу
   * Мейн продолжает работу, но вот только он использует исходное значение slice, поэтому выводится:
   
  ```
   [a a]
   ```
