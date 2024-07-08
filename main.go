package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Map struct {
	TMXName      xml.Name  `xml:"map"`
	Version      string    `xml:"version,attr"`
	TiledVersion string    `xml:"tiledversion,attr"`
	Orientation  string    `xml:"orientation,attr"`
	RenderOrder  string    `xml:"renderorder,attr"`
	Width        int       `xml:"width,attr"`
	Height       int       `xml:"height,attr"`
	TileWidth    int       `xml:"tilewidth,attr"`
	TileHeight   int       `xml:"tileheight,attr"`
	Infinite     int       `xml:"infinite,attr"`
	NextLayerID  int       `xml:"nextlayerid,attr"`
	NextObjectID int       `xml:"nextobjectid,attr"`
	Tilesets     []Tileset `xml:"tileset"`
	Layers       []Layer   `xml:"layer"`
}

type Tileset struct {
	FirstGID int    `xml:"firstgid,attr"`
	Source   string `xml:"source,attr"`
}

type Layer struct {
	ID     int    `xml:"id,attr"`
	Name   string `xml:"name,attr"`
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
	Data   Data   `xml:"data"`
}

type Data struct {
	Encoding string `xml:"encoding,attr"`
	Content  string `xml:",chardata"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <path_to_map.tmx>")
		return
	}

	mapFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer mapFile.Close()

	byteValue, err := io.ReadAll(mapFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var gameMap Map
	err = xml.Unmarshal(byteValue, &gameMap)
	if err != nil {
		fmt.Println("Error unmarshaling TMX file:", err)
		return
	}

	for _, layer := range gameMap.Layers {
		fmt.Printf("Layer ID: %d, Name: %s, Width: %d, Height: %d\n", layer.ID, layer.Name, layer.Width, layer.Height)
		fmt.Println("Data content (CSV):")

		csvReader := csv.NewReader(strings.NewReader(layer.Data.Content))
		csvReader.FieldsPerRecord = -1
		for {
			record, err := csvReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error reading CSV data:", err)
				break
			}
			fmt.Println(record)
		}
		fmt.Println()
	}
}
