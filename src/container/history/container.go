package history

import (
	"context"
	"fmt"
	"github.com/moby/moby/client"
	"log"
)

type containerHistory struct {
	imageName string
}

func IsContainerKnown(image string) (bool, containerHistory, error) {
	var StoredContainer containerHistory

	return false, StoredContainer, nil
}

func AnalyseContainer(err error, cli *client.Client, imageName string) (containerHistory, error) {

	var thisContainer containerHistory
	//TODO:flag on name

	imageHistory, err := cli.ImageHistory(context.Background(), imageName)
	if err != nil {
		log.Fatalf("Error getting image history: %v", err)
	}

	for _, layer := range imageHistory {
		Known, layerDetails, err := IsLayerKnown(layer)

		if err != nil {
			log.Fatalf("Error checking if layer is known: %v", err)
		}

		//its known and not an MCP
		if layerDetails.OK {
			continue
		} else {
			//todo:log to list of MCP issues found
		}

		if !(Known) {
			log.Printf("Layer %s is unknown", layer.ID)

			//send to async process to analyse layer
			layerDetails, err = AnalyseLayer(layer)
			if err != nil {
				log.Fatalf("Error analysing layer %s: %v", layer.ID, err)
			}
		}

		//todo:do something with new layer details

		fmt.Printf("ID: %s\nCreatedBy: %s\nSize: %d bytes\nComment: %s\n\n",
			layer.ID, layer.CreatedBy, layer.Size, layer.Comment)
	}

	//todo store details for IsContainerKnown

	return thisContainer, nil
}
