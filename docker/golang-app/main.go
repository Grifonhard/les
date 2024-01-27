package main

import (
	"fmt"
	"os"
)

func main() {
	var y, m int
	fmt.Println("Добро пожаловать в календарь!!!!!!!!!")
	fmt.Print("Пожалуйста введите год: ")
	fmt.Fscan(os.Stdin, &y)
	fmt.Print("Введите номер любого месяца: ")
	fmt.Fscan(os.Stdin, &m)
	fmt.Printf("year: %d, month: %d\n", y, m)
	fmt.Println("Всего хорошего!")
}
