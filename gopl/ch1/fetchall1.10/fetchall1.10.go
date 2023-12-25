package main

import (
	"fmt"
	"net/http"
	"io"
	"os"
	"time"
)

func main() {
	ch := make(chan string)
	file, err := os.OpenFile("youtube.txt", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			fmt.Println(err)
		}
	for i:=0; i <5; i++ {
		start := time.Now()
		for _, url := range os.Args[1:] {
			go fetch(url, ch)
		}
		for range os.Args[1:] {
			fmt.Println(<-ch)
		}
		file.Write([]byte(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())))
		time.Sleep(time.Second)
	}
	file.Close()
}
func fetch (url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err !=nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}