package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Grifonhard/les/http/device"
)

func main() {
	client := &http.Client{}
	

	//получение девайса для POST
	djson, err := json.Marshal(device.DeviceGet())
	if err != nil {
		fmt.Println(err)
	}


	//Отправка данных
	req1, err := http.NewRequest(http.MethodPost, "http://localhost:8080/device", bytes.NewBuffer(djson))
	if err != nil {
		fmt.Println(err)
	}
	
	//обработка ответа после отправления данных
	resp1, err := client.Do(req1)
	if err != nil {
		fmt.Println(err)
	}
	data1 := make ([]byte, resp1.ContentLength)
	_, err = resp1.Body.Read(data1)
	if err != nil && err != io.EOF {
		fmt.Println(err)
	}
	fmt.Println(string(data1))


	//запрос на получение данных
	req2, err := http.NewRequest(http.MethodGet, "http://localhost:8080/device?id=000000-00000-00000-000001", nil)
	if err != nil {
		fmt.Println(err)
	}	
	resp2, err := client.Do(req2)
	if err != nil {
		fmt.Println(err)
	}

	//обработка ответа с полученными данными
	data2 := make ([]byte, resp2.ContentLength)
	_, err = resp2.Body.Read(data2)
	if err != nil && err != io.EOF {
		fmt.Println(err)
	}
	var d device.Device
	err = json.Unmarshal(data2, &d)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(d)
	
}