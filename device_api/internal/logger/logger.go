package logger

import (
	"log"
	"log/slog"
	"os"
	"path/filepath"
	//"time"
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
	
	
	pathAndFileName := filepath.Join(path, "logs.log")
	return pathAndFileName
}


func Start () (*slog.Logger, *os.File){
	//create/open file
	file, err := os.OpenFile(getPathAndName(), os.O_RDWR|os.O_CREATE|os.O_APPEND, os.FileMode(0666)) 	//создание/открытие файла	
	if err != nil {
		log.Fatal(err)
	}

	//create logger
	logg := slog.New(slog.NewJSONHandler(file, nil))
	return logg, file
}

