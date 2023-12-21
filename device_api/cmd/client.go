package main

import (
	//"context"
	"log/slog"
	//"time"

	"github.com/Grifonhard/les/device_api/internal/logger"
)


func main () {
	//start logging
	logg, file := logger.Start()
	defer file.Close()
	slog.SetDefault(logg)
	
}
