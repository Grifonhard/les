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
func getPathAndName (s string) string {
	//getting the path to save logs
	pathToExec, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	pathToDir := pathToExec[:len(pathToExec)-15]	
	path := filepath.Join(pathToDir, "logs")

	//getting the name of log file
	
	fileName := s + "logs.log"
	pathAndFileName := filepath.Join(path, fileName)
	return pathAndFileName
}


func Start (s string) (*slog.Logger, *os.File){
	//create/open file
	file, err := os.OpenFile(getPathAndName(s), os.O_RDWR|os.O_CREATE|os.O_APPEND, os.FileMode(0666)) 	//создание/открытие файла	
	if err != nil {
		log.Fatal(err)
	}

	//create logger
	logg := slog.New(slog.NewJSONHandler(file, nil))
	return logg, file
}

