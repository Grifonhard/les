package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"log/slog"
	"fmt"

	"github.com/Grifonhard/les/device_api/internal/devices"
	"github.com/Grifonhard/les/device_api/internal/logger"
	"github.com/Grifonhard/les/device_api/internal/storage"
	"github.com/caarlos0/env/v10"
)



func main() {
	//start logging
	logg, file := logger.Start("server")
	defer file.Close()
	slog.SetDefault(logg)

	//request processing
	http.HandleFunc("/device", func (wrtr http.ResponseWriter, req *http.Request){
		switch req.Method {
		case "POST":
			djson := make([]byte, req.ContentLength)
			_, err := req.Body.Read(djson)
			if err != nil && err != io.EOF {
				slog.Error(fmt.Sprint(err))
			}
			var d devices.Device
			err = json.Unmarshal(djson, &d)
			if err != nil {
				slog.Error(fmt.Sprint(err))
			}
			if !storage.Push(&d, logg) {
				wrtr.Write([]byte("device with this id is already in the storage"))
				break
			}
			wrtr.Write([]byte("device added"))
		case "GET":
			var d devices.Device
			b := storage.TakeFrom(req.URL.Query().Get("id"), &d)
			if !b {
				wrtr.Write([]byte("device not found"))
				break
			}
			djson, err := json.Marshal(d)
			if err != nil {
				slog.Error(fmt.Sprint(err))
			}
			wrtr.Write([]byte(djson))
	}
	})
	
	//server
	log.Fatal(http.ListenAndServe(":8080", nil))	
}
