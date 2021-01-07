package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func sumLen(txt string) chan string {
	t := url.QueryEscape(txt)
	encodeurl := "http://" + API_URL + "/?txt=" + t
	ch := make(chan string)
	go fetch(encodeurl, ch)
	return ch
}

func reset() chan string {
	encodeurl := "http://" + API_URL + "/reset"
	ch := make(chan string)
	go fetch(encodeurl, ch)
	return ch
}

func fetch(url string, ch chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	ch <- fmt.Sprintf("%s", body)
}
