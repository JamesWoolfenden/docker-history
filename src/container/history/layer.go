package history

import "github.com/moby/moby/api/types/image"

type Layer struct {
	MCP     string
	History image.HistoryResponseItem
	OK      bool
}

func LayerHistory() {

}

func IsLayerKnown(layer image.HistoryResponseItem) (ok bool, fullLayer Layer, err error) {
	var StoredLayer Layer
	//todo: get layer record is sha found otherwise false

	return false, StoredLayer, nil
}

func AnalyseLayer(history image.HistoryResponseItem) (Layer, error) {
	var newLayer Layer
	newLayer.History = history

	//todo do something to look at layer
	//todo: add record to layerdb
	return newLayer, nil
}
