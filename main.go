package main

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Map struct {
	Version      string    `xml:"version,attr" json:"version"`
	TiledVersion string    `xml:"tiledversion,attr" json:"tiledVersion"`
	Orientation  string    `xml:"orientation,attr" json:"orientation"`
	RenderOrder  string    `xml:"renderorder,attr" json:"renderOrder"`
	Width        int       `xml:"width,attr" json:"width"`
	Height       int       `xml:"height,attr" json:"height"`
	TileWidth    int       `xml:"tilewidth,attr" json:"tileWidth"`
	TileHeight   int       `xml:"tileheight,attr" json:"tileHeight"`
	Infinite     int       `xml:"infinite,attr" json:"infinite"`
	NextLayerID  int       `xml:"nextlayerid,attr" json:"nextLayerID"`
	NextObjectID int       `xml:"nextobjectid,attr" json:"nextObjectID"`
	Tilesets     []Tileset `xml:"tileset" json:"tilesets"`
	Layers       []Layer   `xml:"layer" json:"layers"`
}

type Tileset struct {
	FirstGID int    `xml:"firstgid,attr" json:"firstGID"`
	Source   string `xml:"source,attr" json:"source"`
}

type Layer struct {
	ID     int    `xml:"id,attr" json:"id"`
	Name   string `xml:"name,attr" json:"name"`
	Width  int    `xml:"width,attr" json:"width"`
	Height int    `xml:"height,attr" json:"height"`
	Data   Data   `xml:"data" json:"data"`
}

type Data struct {
	Encoding string     `xml:"encoding,attr" json:"encoding"`
	Content  [][]string `xml:"-" json:"content"`
	Raw      string     `xml:",chardata" json:"-"`
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <path_to_map.xml> <output_json_path>")
		return
	}

	filePath := os.Args[1]
	mapFile, err := os.Open(filePath)
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

	for i, layer := range gameMap.Layers {
		csvReader := csv.NewReader(strings.NewReader(layer.Data.Raw))
		csvReader.FieldsPerRecord = -1
		var csvData [][]string
		for {
			record, err := csvReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error reading CSV data:", err)
				break
			}
			csvData = append(csvData, record)
		}
		gameMap.Layers[i].Data.Content = csvData
	}

	jsonData, err := json.MarshalIndent(gameMap, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	outputFile := os.Args[2]
	err = os.WriteFile(outputFile, jsonData, 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	fmt.Printf("JSON data successfully written to %s\n", outputFile)
}
