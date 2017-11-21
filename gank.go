package main

import (
	"github.com/cls1991/gank.io-go/core"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	url, file := "http://gank.io", "out/gank.txt"
	list, err := core.Extract(url)
	if err != nil {
		log.Fatal(err)
	}
	dat := strings.Join(list, "\n")
	if err := ioutil.WriteFile(file, []byte(dat), 0644); err != nil {
		log.Fatal(err)
	}
}
