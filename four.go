package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

func quickSorting(ar [] float64) []float64 {
	if len(ar) == 1 {
		return ar
	}


	result := make([]float64, 0)

	
	if len(ar) == 2 {
		if ar[0] > ar[1] {
			result = append(result, ar[1])
			result = append(result, ar[0])
			return result
		} else {
			return ar
		}
	}

	m := make([]float64, 0)
	l := make([]float64, 0)
	j := len(ar)/2


	for i, a := range ar {
		if i == j {
			continue
		}
		if a >= ar[j] {
			m = append(m, a)
		} else {
			l = append(l, a)
		}
	}

	if len(l) > 0 {
		ls := quickSorting(l)
		for _, la := range ls {
			result = append(result, la)
		}
	}
	result = append(result, ar[j])
	if len(m) > 0 {
		ms := quickSorting(m)
		for _, ma := range ms {
			result = append(result, ma)
		}
	}
	
	return result
}

func main() {
	fmt.Println("Enter an array of numbers in the following form: \"23.13 14.21 7 64 0.14 ..\"")
	fmt.Printf("Your array: ")	
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input, _ = strings.CutSuffix(input, "\r\n")
	arrString := strings.Split(input, " ")
	
	arr := make([]float64, len(arrString))
	for i, aS := range arrString {
		ar, err := strconv.ParseFloat(aS, 64)
		if err != nil {
			fmt.Println(err)
		}
		arr[i] = ar
	}


	fmt.Println("Sorted array: ", quickSorting(arr))
}