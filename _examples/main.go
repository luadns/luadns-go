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

	ctx := context.Background()

	c := api.NewClient(email, key, api.SetBaseURL(url))
	user, err := c.Me(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("email:  ", user.Email)
	fmt.Println("name:   ", user.Name)
	fmt.Println("package:", user.Package)

	var meta api.ListMeta
	zones, err := c.ListZones(ctx, &api.ListParams{}, api.GetListMeta(&meta))
	if err != nil {
		log.Fatalln(err)
	}

	for _, z := range zones {
		fmt.Println("===> zone", z.Name)
		records, err := c.ListRecords(ctx, z, &api.ListParams{})
		if err != nil {
			log.Fatalln(err)
		}

		for _, r := range records {
			fmt.Println("    ", r.Name, r.Type, r.Content, r.TTL)
		}
	}

	if len(zones) == 0 {
		fmt.Println("===> No zones found, skipping other tests")
	}

	zone := zones[0]
	name := "foo." + zone.Name + "."

	fmt.Println("===> create many")
	created, err := c.CreateManyRecords(ctx, zone, []*api.RR{{Name: name, Type: "TXT", Content: "foo"}})
	if err != nil {
		log.Fatalln(err)
	}
	for _, r := range created {
		fmt.Println("    ", r.Name, r.Type, r.Content, r.TTL)
	}

	fmt.Println("===> update many")
	updated, err := c.UpdateManyRecords(ctx, zone, []*api.RR{{Name: name, Type: "TXT", Content: "bar"}})
	if err != nil {
		log.Fatalln(err)
	}
	for _, r := range updated {
		fmt.Println("    ", r.Name, r.Type, r.Content, r.TTL)
	}

	fmt.Println("===> delete many")
	deleted, err := c.DeleteManyRecords(ctx, zone, []*api.RR{{Name: name, Type: "TXT"}})
	if err != nil {
		log.Fatalln(err)
	}
	for _, r := range deleted {
		fmt.Println("    ", r.Name, r.Type, r.Content, r.TTL)
	}
}
