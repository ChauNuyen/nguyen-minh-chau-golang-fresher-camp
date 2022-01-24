package main

import (
	"fmt"
	"time"
)

func crawlData(gr int, ch chan int) {
	for i := range ch {
		fmt.Printf("GoRountine number %v crawled data from %v\n", gr+1, i)
	}
}

func main() {
	urlList := make(chan int, 10000)

	for i := 0; i < 5; i++ {
		go crawlData(i, urlList)
	}

	for i := 0; i < 2000; i++ {
		time.Sleep(10 * time.Millisecond)
		urlList <- i
	}

	fmt.Println("Finished crawl data")
}
