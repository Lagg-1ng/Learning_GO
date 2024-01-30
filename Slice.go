package main

import (
	"fmt"
)

func main() {
	// Создаем слайс и заполняем его числами от 1 до 100
	slice := make([]int, 100)
	for i := 0; i < 100; i++ {
		slice[i] = i + 1
	}

	// Оставляем в слайсе первые и последние 10 элементов
	slice = append(slice[:10], slice[len(slice)-10:]...)

	// Разворачиваем слайс в обратном порядке
	reverseSlice(slice)

	// Выводим результат
	fmt.Println(slice)
}

// Функция для разворота слайса в обратном порядке
func reverseSlice(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
