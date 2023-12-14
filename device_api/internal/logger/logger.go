package logger

import (
	"time"
	"fmt"
	"os"
)

func Start () {
	pathAndFileName := "device_api/internal/logs/" + fmt.Fprintln()+".log"
	//создание/открытие файла
	_, err := os.OpenFile(pathAndFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.FileMode(0666)) 	//создание/открытие файла	
	if err != nil {
		fmt.Println(err)
	}
}