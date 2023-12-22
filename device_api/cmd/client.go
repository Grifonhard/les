package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Grifonhard/les/device_api/internal/devices"
	"github.com/Grifonhard/les/device_api/internal/logger"
)


func main () {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	c := make(chan os.Signal, 1)
 	
	//start logging
	logg, file := logger.Start("client")
	defer file.Close()
	slog.SetDefault(logg)

	wg.Add(1)
	go processPost(&wg, ctx, logg) 

	wg.Add(1)
	go processGet(&wg, ctx, logg)

	signal.Notify(c)
	_ = <-c
	cancel()
	wg.Wait()
}

func processPost (wg *sync.WaitGroup, ctx context.Context, logg* slog.Logger) {
	defer wg.Done()
	client := &http.Client{}
	for {
		select {
		case <- ctx.Done():
			logg.Info("processPost stop")
			return
		case <- time.After(time.Second):
			//receiving a device for POST
			djson, err := json.Marshal(devices.DeviceGet())
			if err != nil {
				slog.Error(fmt.Sprint(err))
			}
			//Send request
			req1, err := http.NewRequest(http.MethodPost, "http://localhost:8080/device", bytes.NewBuffer(djson))
			if err != nil {
				slog.Error(fmt.Sprint(err))
			}	
			//response processing
			resp1, err := client.Do(req1)
			if err != nil {
				slog.Error(fmt.Sprint(err))
			}
			data1 := make ([]byte, resp1.ContentLength)
			_, err = resp1.Body.Read(data1)
			if err != nil && err != io.EOF {
				slog.Error(fmt.Sprint(err))
			}
			slog.Info(string(data1))
		}
	}
}

func processGet (wg *sync.WaitGroup, ctx context.Context, logg* slog.Logger) {
	defer wg.Done()	
	client := &http.Client{}
	for {
		select {
		case <- ctx.Done():
			logg.Info("processGet stop")
			return
		case <- time.After(time.Second):
			//receiving id for GET
			requstURL := "http://localhost:8080/device?id=" + devices.IdGet()
			//data request
			req2, err := http.NewRequest(http.MethodGet, requstURL, nil)
			if err != nil {
				slog.Error(fmt.Sprint(err))
			}	
			resp2, err := client.Do(req2)
			if err != nil {
				slog.Error(fmt.Sprint(err))
			}
			//processing the response with the received data
			data2 := make ([]byte, resp2.ContentLength)
			_, err = resp2.Body.Read(data2)
			if err != nil && err != io.EOF {
				slog.Error(fmt.Sprint(err))
			}
			var d devices.Device
			err = json.Unmarshal(data2, &d)
			if err != nil {
				slog.Error(fmt.Sprint(err))
			}
			slog.Info("device get from server: ", slog.String("ID", d.Id), slog.String("Language", d.Language), slog.Float64("Latitude", d.Latitude), slog.Float64("Longitude", d.Longitude), slog.String("Os_name", d.Os_name), slog.String("Os_version", d.Os_version))
		}
	}
}	
	