package main

import (
	"context"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"sync"
	"time"
	"os/signal"

	"github.com/Grifonhard/les/device_api/internal/devices"
	"github.com/Grifonhard/les/device_api/internal/logger"
)


func main () {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	
	//start logging
	logg, file := logger.Start("client")
	defer file.Close()
	slog.SetDefault(logg)

	wg.Add(1)
	go processPost(&wg, ctx, logg) 

	wg.Add(1)
	go processGet(&wg, ctx, logg)

	
	wg.Wait()
	cancel()
}

func processPost (wg *sync.WaitGroup, ctx context.Context, logg* slog.Logger) {
	client := &http.Client{}

	for {
		select {
		case <- ctx.Done():
			logg.Info("processPost stop")
			break
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
	wg.Done()
}

func processGet (wg *sync.WaitGroup, ctx context.Context, logg* slog.Logger) {
	client := &http.Client{}
	for {
		select {
		case <- ctx.Done():
			logg.Info("processGet stop")
			break
		case <- time.After(time.Second):
			//data request
			req2, err := http.NewRequest(http.MethodGet, "http://localhost:8080/device?id=000000-00000-00000-000002", nil)
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
	wg.Done()
}	
	