package main

import (
    "context"
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

    imageName := "mcp/aws-kb-retrieval"
    history, err := cli.ImageHistory(context.Background(), imageName)
    if err != nil {
        log.Fatalf("Error getting image history: %v", err)
    }

    for _, layer := range history {
        fmt.Printf("ID: %s\nCreatedBy: %s\nSize: %d bytes\nComment: %s\n\n",
            layer.ID, layer.CreatedBy, layer.Size, layer.Comment)
    }
}
