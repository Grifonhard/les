package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)



func logWriter(closeCh chan string, path string, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()


	var line []byte
	var err error
	lineCh := make(chan string, 10)															//канал для передачи прочитанных строк функции записи


	file, err := os.OpenFile(path, os.O_RDONLY, os.FileMode(0444))							//открытие файла для чтения
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()


	go checkAndWrightLog(lineCh)															//горутина для записи в файлы


	reader := bufio.NewReader(file)															//первый проход
	for{
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error while reading file ", path , ": ", err)
		}
		lineCh <- string(line)
	}


	for{																					//цикл
		select{
		
		case closeRoutine := <-closeCh:														//закрытие 
			lineCh <- closeRoutine
			wg.Done()				

		default:																			//попытка чтения
			line, _, err = reader.ReadLine()
			if err != nil {
				if err == io.EOF {
					time.Sleep(1*time.Second)
					continue
				}
				fmt.Println("error while reading file ", path , ": ", err)
			}
			lineCh <- string(line)
		}
	}
}



func checkAndWrightLog(ch chan string) {
	var line string
	var hour string
	var pathAndFileName string
	var file *os.File


	for {
		line = <- ch


		if strings.EqualFold(line, "stop") {															//проверка на необходимость выйти
			err := file.Close()
			if err != nil {
				fmt.Println("Error while closing file ", pathAndFileName, " ", err)
			}
			return
		}


		reg1, err := regexp.Compile("error")     														//проверка на наличие слова error в строке         
			if err != nil {
				fmt.Println(err)
			}
		b := reg1.Match([]byte(line))
		if !b {
			continue
		}		
	
	
		reg2, err := regexp.Compile("\\d\\d\\d\\d-\\d\\d-\\d\\d \\d\\d")     							//поиск и сохранение даты и часа          
			if err != nil {
				fmt.Println(err)
			}
		dateAndHourRow := reg2.FindString(line)
		dateAndHour := strings.Split(dateAndHourRow, " ")
		

		if strings.EqualFold(pathAndFileName, "") {														//создание имени файла в первый раз
			pathAndFileName = "file/Error_" + dateAndHour[0] + "_" + dateAndHour[1] + ".log"
			hour = dateAndHour[1]
		}


		if !strings.EqualFold(dateAndHour[1], hour) && !strings.EqualFold(hour, "") {					//проверка на необходимость обновления имени файла и закрытия старого файла
			hour = dateAndHour[1]

			err = file.Close()
			if err != nil {
				fmt.Println("Error while closing file ", pathAndFileName, " ", err)
			}

			pathAndFileName = "file/Error_" + dateAndHour[0] + "_" + dateAndHour[1] + ".log"
		}




		file, err = os.OpenFile(pathAndFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.FileMode(0666)) 	//создание/открытие файла	
		if err != nil {
			fmt.Println(err)
		}

		_, err = file.WriteString(string(line+"\n"))													//запись в файл
		if err != nil {
			fmt.Println("error while writing in file ", pathAndFileName, " ", err)
		}
	}
}



func main(){
	wg := sync.WaitGroup{}
	closeCh := make(chan string)

	flag.Parse()
	for _, path := range flag.Args() {																		//передача файлов в горутины для записи логов
		go logWriter(closeCh, path, &wg)
	}
	

	scanner := bufio.NewScanner(os.Stdin)																	//чтение консоли в ожидании ввода stop
	fmt.Println("For quit enter \"stop\"")
	for scanner.Scan() {
		word := scanner.Text()
		if strings.EqualFold(word, "stop"){
			for i :=0; i < len(flag.Args()); i++ {
				closeCh <- word
			}
			break
		}
	}

	wg.Wait()
	fmt.Println("Bye!")
}