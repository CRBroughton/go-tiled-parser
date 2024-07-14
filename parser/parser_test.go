package parser_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/crbroughton/go-tiled-parser/parser"
	"github.com/stretchr/testify/assert"
)

func TestItConvertsAMapFileTOJSON(t *testing.T) {

	file, err := os.ReadFile("../mocks/data.tmx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	expectation := parser.Map{
		Version:      "1.10",
		TiledVersion: "1.10.2",
		Orientation:  "orthogonal",
		RenderOrder:  "right-down",
		Width:        5,
		Height:       6,
		TileWidth:    16,
		TileHeight:   16,
		Infinite:     0,
		NextLayerID:  2,
		NextObjectID: 1,
		Tilesets: []parser.Tileset{
			{
				FirstGID: 1,
				Source:   "../../go-animal-crossing/TileSets/Grass.tsx",
			},
		},
		Layers: []parser.Layer{
			{
				ID:     1,
				Name:   "Tile Layer 1",
				Width:  5,
				Height: 6,
				Data: parser.Data{
					Encoding: "csv",
					Content: [][]string{
						{
							"0",
							"0",
							"0",
							"0",
							"0",
						},
						{
							"62",
							"0",
							"0",
							"0",
							"0",
						},
						{
							"0",
							"0",
							"62",
							"0",
							"0",
						},
						{
							"0",
							"62",
							"0",
							"62",
							"0",
						},
						{
							"0",
							"0",
							"0",
							"0",
							"0",
						},
						{
							"0",
							"0",
							"0",
							"0",
						},
					},
					Raw: string("\n0,0,0,0,0,\n62,0,0,0,0,\n0,0,62,0,0,\n0,62,0,62,0,\n0,0,0,0,0,\n0,0,0,0,0\n"),
				},
			},
		},
	}
	mapData := parser.GetMapData(file)

	assert.Equal(t, expectation, mapData)
}
