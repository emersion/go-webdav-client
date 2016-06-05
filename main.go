// +build ignore

package main

import (
	"log"

	"github.com/emersion/go-webdav-client"
)

func main() {
	c := client.New("http://127.0.0.1:5232/")

	f, err := c.OpenFile("")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(f.Stat())
}
