package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func DownloadImage(url string, loc string, wg *sync.WaitGroup) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	f, err := os.Create(loc)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	io.Copy(f, res.Body)

	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	startTime := time.Now()
	go DownloadImage("https://source.unsplash.com/random", "randomImage.png", &wg)
	wg.Wait()
	diff := time.Now().Sub(startTime)
	fmt.Println(diff)
}
