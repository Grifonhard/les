package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"io"
	"sort"
	"time"
)

func main() {
	fmt.Printf("Enter path to file with array: ")	                   		   //чтение пути к файлу из консоли
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input, _ = strings.CutSuffix(input, "\r\n")

	fileArr, err := os.OpenFile(input, os.O_RDONLY, os.FileMode(0444))			//открытие файла
	if err != nil {
		fmt.Println(err)
	}
	defer fileArr.Close()

	data := make([]byte, 1000000)												//чтение в байтовый срез
	n, err := fileArr.Read(data)
	if err != nil && err != io.EOF {
		fmt.Println(err)
	}

	
	arrString := strings.Split(string(data[:n]), " ")							//преобразование в массив чисел
	arr := make([]int, len(arrString))
	for i, aS := range arrString {
		ar, err := strconv.ParseInt(aS, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		arr[i] = int(ar)
	}

	sort.IntSlice(arr).Sort()													//демонстрация отсортированного массива и выбор необходимого числа
	fmt.Println("Array for binary search: ", arr)
	fmt.Printf("select a number to search: ")
	input, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input, _ = strings.CutSuffix(input, "\r\n")
	rN, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	reqNum := int(rN)

	begin := time.Now()															//поиск
	var low, high, mid, found int
	low = 0
	high = len(arr) - 1
	for low <= high {
		mid = ( low + high ) / 2
		if arr[mid] > reqNum {
			high = mid - 1
		} else if arr[mid] == reqNum {
			found = reqNum
			break
		} else {
			low = mid + 1
		}
	}
	end := time.Now()
	result := end.UnixNano() - begin.UnixNano()
	fmt.Printf("время выполнения бинарного поиска: %d ns, %d(request number) = %d(found)\n", result, reqNum, found)

}