package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	api "github.com/luadns/luadns-go"
)

const (
	baseURL = "https://api.luadns.com/v1"
)

var email string
var key string
var url string

func main() {
	flag.StringVar(&email, "email", "joe@example.com", "your email address")
	flag.StringVar(&key, "key", "", "your API key")
	flag.StringVar(&url, "url", baseURL, "base URL")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	c := api.NewClient(context.Background(), email, key, api.SetBaseURL(url))
	user, err := c.Me()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("email:  ", user.Email)
	fmt.Println("name:   ", user.Name)
	fmt.Println("package:", user.Package)

	zones, err := c.ListZones()
	if err != nil {
		log.Fatalln(err)
	}

	for _, z := range zones {
		fmt.Println("===> zone", z.Name)
		records, err := c.ListRecords(z)
		if err != nil {
			log.Fatalln(err)
		}

		for _, r := range records {
			fmt.Println("    ", r.Name, r.Type, r.Content, r.TTL)
		}
	}
}
