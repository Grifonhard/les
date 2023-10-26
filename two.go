package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"strconv"
	"math/rand"
)

func main() {
	fmt.Print("Enter the programm running time in seconds: ")				//получение времени работы
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')					
		if err != nil {
			fmt.Println(err)
		}
	input = strings.TrimSuffix(input, "\r\n")
	secondsOfWork, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	
	source := rand.NewSource(time.Now().UnixMicro())						//отправляемые данные			
	sRand := rand.New(source)


	ch := make(chan int)
	go showInt(ch)


	begin := time.Now().Unix()
	for secondsOfWork > time.Now().Unix() - begin {
		ch <- sRand.Int()
	}
}

func showInt (ch chan int) {
	for {
		fmt.Println(<- ch)
	}
}