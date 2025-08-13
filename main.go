package main

import (
	"docker-history/src/container/history"
	"fmt"
	"log"
	//"github.com/moby/moby/api/types"
	"github.com/moby/moby/client"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Error creating Docker client: %v", err)
	}

	imageName := "mcp/memory"
	Known, History, err := history.IsContainerKnown(imageName)

	if err != nil {
		log.Fatal(err)
	}

	if !(Known) {
		log.Printf("Container %s is unknown", imageName)
		History, err = history.AnalyseContainer(err, cli, imageName) //send to async process to analyse layer
	}

	fmt.Print(History)
}
