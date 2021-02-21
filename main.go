package main

import (
	"context"
	"log"

	"github.com/ho-wan/go-client-practice/github"
)

func main() {
	gc := github.NewClient()

	ctx := context.Background()

	err := gc.GetList(ctx)
	if err != nil {
		log.Fatalln("failed to get list", err)
	}
}
