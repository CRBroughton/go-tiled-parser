package parser_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/crbroughton/go-tiled-parser/parser"
	"github.com/stretchr/testify/assert"
)

func TestItConvertsAMapFileToJSON(t *testing.T) {
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
		Tilesets: []parser.TilesetReference{
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
					Content: []parser.DataTile{
						{
							Tile: "0",
							X:    0,
							Y:    0,
						}, {
							Tile: "0",
							X:    1,
							Y:    0,
						}, {
							Tile: "0",
							X:    2,
							Y:    0,
						}, {
							Tile: "0",
							X:    3,
							Y:    0,
						}, {
							Tile: "0",
							X:    4,
							Y:    0,
						}, {
							Tile: "62",
							X:    0,
							Y:    1,
						}, {
							Tile: "0",
							X:    1,
							Y:    1,
						}, {
							Tile: "0",
							X:    2,
							Y:    1,
						}, {
							Tile: "0",
							X:    3,
							Y:    1,
						}, {
							Tile: "0",
							X:    4,
							Y:    1,
						}, {
							Tile: "0",
							X:    0,
							Y:    2,
						}, {
							Tile: "0",
							X:    1,
							Y:    2,
						}, {
							Tile: "62",
							X:    2,
							Y:    2,
						}, {
							Tile: "0",
							X:    3,
							Y:    2,
						}, {
							Tile: "0",
							X:    4,
							Y:    2,
						}, {
							Tile: "0",
							X:    0,
							Y:    3,
						}, {
							Tile: "62",
							X:    1,
							Y:    3,
						}, {
							Tile: "0",
							X:    2,
							Y:    3,
						}, {
							Tile: "62",
							X:    3,
							Y:    3,
						}, {
							Tile: "0",
							X:    4,
							Y:    3,
						}, {
							Tile: "0",
							X:    0,
							Y:    4,
						}, {
							Tile: "0",
							X:    1,
							Y:    4,
						}, {
							Tile: "0",
							X:    2,
							Y:    4,
						}, {
							Tile: "0",
							X:    3,
							Y:    4,
						}, {
							Tile: "0",
							X:    4,
							Y:    4,
						}, {
							Tile: "0",
							X:    0,
							Y:    5,
						}, {
							Tile: "3",
							X:    1,
							Y:    5,
						}, {
							Tile: "0",
							X:    2,
							Y:    5,
						}, {
							Tile: "0",
							X:    3,
							Y:    5,
						}, {
							Tile: "54",
							X:    4,
							Y:    5,
						},
					},
					Raw: string("\n0,0,0,0,0,\n62,0,0,0,0,\n0,0,62,0,0,\n0,62,0,62,0,\n0,0,0,0,0,\n0,3,0,0,54\n"),
				},
			},
		},
	}
	mapData := parser.GetMapData(file)

	assert.Equal(t, expectation, mapData)
}

func TestItConvertsATilesetFileToJSON(t *testing.T) {
	file, err := os.ReadFile("../mocks/tileset.tsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	expectation := parser.Tileset{
		Version:      "1.10",
		TiledVersion: "1.10.2",
		Name:         "Grass",
		TileWidth:    16,
		TileHeight:   16,
		TileCount:    77,
		Columns:      11,
		Image: parser.Image{
			Source: "../assets/Tilesets/Grass.png",
			Width:  176,
			Height: 112,
		},
		Tiles: []parser.Tile{
			{
				ID: 36,
				Properties: []parser.Property{
					{
						Name:  "collide",
						Type:  "bool",
						Value: "true",
					},
				},
			},
		},
	}
	tileData := parser.GetTilesetData(file)

	assert.Equal(t, expectation, tileData)
}
