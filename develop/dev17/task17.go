package main

import (
	"fmt"
	"sort"
)

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func main() {
	// Пусть исходный массив будет:
	arr := []int{9, 7, 5, 100, 12, 2, 14, 3, 10, 6, 14, 2}

	// Отсортируем, так как бинарный поиск работает только с отсортированным массивом
	sort.Ints(arr)
	fmt.Println(arr)

	// Элемент, который будем искать (его индекс)
	target := 100

	ans := handmadeBinarySearch(arr, target)
	fmt.Println(ans)
}

func handmadeBinarySearch(array []int, value int) int {
	low := 0
	high := len(array) - 1

	// Проверка, что в исходном массиве больше чем 1 элемент (как в 16 таске)
	for low < high {
		// Поиск элемента в середине массива
		mid := (low + high) / 2
		sug := array[mid]
		if sug == value {
			return mid
		} else if sug > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if array[low] == value {
		return low
	}
	// Если не удалось найти элемент
	return -1
}
