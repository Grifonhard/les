package main

import (
	"fmt"
	"os"
	"regexp"
	"bufio"
	"strings"
	"sort"
	"time"
	
)

var logs []string						//сохранение всех логов
var index = make(map[string]int)    	//сохранение мест чтения 
var ind int								//последняя записанная строка из логов

func main() {
	var sDir string  						//директория
	var input string						//для штатной остановки

	if len(os.Args)== 1 {
		fmt.Println("You not enter path and files")
		return
	}
	if len(os.Args) == 2 {
		fmt.Println("You not enter files")
		return
	}		

	for i, s := range os.Args{   //Первое чтение директории, списка файлов и содержимого файлов
		if i == 0 {
			continue
		}
		if i == 1 {
			b, err := regexp.MatchString("^./", s)
			if err != nil{
				fmt.Println(err)
			}
			if !b {
				fmt.Println(s + " wrong directory")
				return
			}
			sDir = s
			continue
		}
		
		readLog(s)	
	}
	
	sDir, b := strings.CutPrefix(sDir, "./")		//удаление префикса ./
	if !b {
		fmt.Println("Ups")
	}
	writeLogs(sDir)				//Первая запись

	fmt.Println("programm log-collector is running")
	for input != "stop log-collector" {											//цикличное чтение раз в минуту
		time.Sleep(1*time.Minute)
		for i, s := range os.Args {
			if i == 0 || i == 1 {
				continue
			}
			readLog(s)
			writeLogs(sDir)
		}
	}
}

func readLog(fileName string) {	
	var i int

	file, err := os.OpenFile(fileName, os.O_RDONLY, os.FileMode(0444))  //открытие файла
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()


	scanner := bufio.NewScanner(file)									//добавление новых строк в log
	for scanner.Scan() {
		i++
		if i > index[fileName] {
		    b, err := regexp.MatchString("error", scanner.Text())		
			if err != nil{
				fmt.Println(err)
			}
			if b {
				logs = append(logs, scanner.Text())
			}
		}
	}
	
	index[fileName] = i                                                //сохранение места остановки чтения
}

func writeLogs(path string) {
	
	
	if ind == len(logs) - 1 {
		return
	}
	
	
	sort.Strings(logs)


	var str string										//для создания файла
	var splitStr []string								//для создания файла
	var indexWrite = make(map[string]struct{})			//для проверки открыт ли нужный файл
	var file *os.File                                   //для закрытия файла перед открытием нового
	var idx int											//для первого прохода, чтоб не закрывать пустой file


	for i, s := range logs {
		
		if i <= ind {																	//начало с места остановки
			continue
		}		
		ind++

		
		reg, err := regexp.Compile("\\d\\d\\d\\d-\\d\\d-\\d\\d \\d\\d")     			//регулярное выражение          
		if err != nil {
			fmt.Println(err)
		}
		str = reg.FindString(s)													//поиск и сохранение найденного регулярного выражения
		
		
		_, ok := indexWrite[str]		
		if ok {
			_, err = file.WriteString(string(s+"\n"))
			if err != nil {
				fmt.Println(err)
			}
			continue
		}


		if idx != 0 {																	//закрытие ранее открытого файла
			err = file.Close()
			if err != nil {
				fmt.Println(err)
			}
		}


		indexWrite[str] = struct{}{}

		splitStr = strings.Split(str, " ")
		str = path + "/" + "Error_" + splitStr[0] + "_" + splitStr[1] + ".log"
		file, err = os.OpenFile(str, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.FileMode(0666)) 	//создание/открытие файла	
		if err != nil {
			fmt.Println(err)
		}


		_, err = file.WriteString(string(s+"\n"))
		if err != nil {
			fmt.Println(err)
		}


		idx++
	}

	ind = len(logs) - 1
}