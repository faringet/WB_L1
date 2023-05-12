package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {
	// Пусть первое множество будет:
	set1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// второе:
	set2 := []int{10, 11, 12, 13, 14, 15, 3, 4, 5}

	result := intersection(set1, set2)
	fmt.Println(result)
}

func intersection(set1, set2 []int) []int {
	intersect := make([]int, 0)  // слайс для пересечения
	setMap := make(map[int]bool) // мапа для элементов первого множества

	// setMap заполняется элементами из первого множества как ключи, все значения - true
	for _, num := range set1 {
		setMap[num] = true
		//fmt.Println(setMap)
	}

	// пробегаемся по второму множеству и добавляем в intersect только совпадающими элементы
	for _, num := range set2 {
		if setMap[num] {
			intersect = append(intersect, num)
			//fmt.Println(intersect)
			setMap[num] = false // как только нашли совпадение ставим false - чтобы избежать повторений в результате
		}
	}
	return intersect
}
