package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Grifonhard/les/http/device"
	"github.com/Grifonhard/les/http/storage"
)

func main() {
	//обработка запроса
	http.HandleFunc("/device", func (wrtr http.ResponseWriter, req *http.Request){
		switch req.Method {
		case "POST":
			djson := make([]byte, req.ContentLength)
			_, err := req.Body.Read(djson)
			if err != nil && err != io.EOF {
				fmt.Println(err)
			}
			var d device.Device
			err = json.Unmarshal(djson, &d)
			if err != nil {
				fmt.Println(err)
			}
			if !storage.Push(d) {
				wrtr.Write([]byte("девайс с таким id уже в хранилище"))
				break
			}
			wrtr.Write([]byte("девайс успешно добавлен"))
		case "GET":
			var d device.Device
			b := storage.TakeFrom(req.URL.Query().Get("id"), &d)
			if b == false {
				wrtr.Write([]byte("девайс не найден"))
				break
			}
			djson, err := json.Marshal(d)
			if err != nil {
				fmt.Println(err)
			}
			wrtr.Write([]byte(djson))
	}
	})
	
	//сервер
	log.Fatal(http.ListenAndServe(":8080", nil))	
}
