package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

func main() {
	edges := make(chan string)
	unprocessed := make(chan string)
	var root string
	flag.StringVar(&root, "root", "http://www.androidpolice.com", "root url")
	flag.Parse()
	res, err := url.Parse(root)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(res.Hostname())
	go crawler(edges)
	go urlExtractor(unprocessed)
	edges <- root
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
}

func crawler(channel <-chan string) {
	for url := range channel {
		fmt.Println("Now processing:", url)
		resp, err := http.Get(url)
		if err != nil {
			log.Panicln(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Panicln(err)
		}
		strbody := string(body[:len(body)])
		fmt.Println(strbody)
	}
}

func urlExtractor(channel <-chan byte[]) {
	r, err := regexp.Compile("/[-a-zA-Z0-9@:%_\\+.~#?&//=]{2,256}\\.[a-z]{2,4}\b(\\/[-a-zA-Z0-9@:%_\\+.~#?&//=]*)?/gi")
	if err != nil {
		log.Panicln(err)
	}
	for htmlBody := range channel {
		matches, _ := r.FindAll(htmlBody, -1)

	}
}
