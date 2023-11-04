package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	var b bool = true
	
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter string, for quit enter \"quit\"")
	for b {
		runes := make(map[rune]int)
		result := 1

		input, err := reader.ReadString('\n')					//получение ввода
		if err != nil {
			fmt.Println(err)
		}
		input = strings.TrimSpace(input)
		input = strings.ToLower(input)

		b = !strings.EqualFold(input, "quit")					//выход
		if !b {
			continue
		}

		for _, r := range input {								//анализ
			runes[r] += 1
			result *= runes[r]
		}

		if result == 1 {										//вердикт
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}

	}
	fmt.Println("Bye")
}