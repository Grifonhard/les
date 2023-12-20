package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
	//"strings"
)


//getting the path and name of file to save logs
func getPathAndName () string {
	//getting the path to save logs
	pathToExec, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	pathToDir := pathToExec[:len(pathToExec)-15]	
	path := filepath.Join(pathToDir, "logs")

	//getting the name of log file
	
	fileName := ""
	pathAndFileName := filepath.Join(path, "logs.log")

}


func Start () string {
	//getting the path to save logs
	pathToExec, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	pathToDir := pathToExec[:len(pathToExec)-15]	
	path := filepath.Join(pathToDir, "logs")
	pathAndFileName := filepath.Join(path, "logs.log")

	//create/open file
	_, err = os.OpenFile(pathAndFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.FileMode(0666)) 	//создание/открытие файла	
	if err != nil {
		log.Fatal(err)
	}
}