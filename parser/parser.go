package parser

import (
	"encoding/csv"
	"encoding/xml"
	"io"
	"log"
	"strings"
)

type Map struct {
	Version      string             `xml:"version,attr" json:"version"`
	TiledVersion string             `xml:"tiledversion,attr" json:"tiledVersion"`
	Orientation  string             `xml:"orientation,attr" json:"orientation"`
	RenderOrder  string             `xml:"renderorder,attr" json:"renderOrder"`
	Width        int                `xml:"width,attr" json:"width"`
	Height       int                `xml:"height,attr" json:"height"`
	TileWidth    int                `xml:"tilewidth,attr" json:"tileWidth"`
	TileHeight   int                `xml:"tileheight,attr" json:"tileHeight"`
	Infinite     int                `xml:"infinite,attr" json:"infinite"`
	NextLayerID  int                `xml:"nextlayerid,attr" json:"nextLayerID"`
	NextObjectID int                `xml:"nextobjectid,attr" json:"nextObjectID"`
	Tilesets     []TilesetReference `xml:"tileset" json:"tilesets"`
	Layers       []Layer            `xml:"layer" json:"layers"`
}

type TilesetReference struct {
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

func removeEmptyStrings(strings []string) []string {
	var result []string
	for _, str := range strings {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}

// Utility function to flatten out the Content array.
func Flatten(array [][]string) []string {
	var result []string
	for _, subArray := range array {
		result = append(result, subArray...)
	}
	return result
}

// Accepts an array of bytes that represent the .tmx file,
// and returns the resulting data.
func GetMapData(gameBytes []byte) Map {
	var gameMap Map
	err := xml.Unmarshal(gameBytes, &gameMap)
	if err != nil {
		log.Fatal("Error unmarshaling TMX file:", err)
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
				log.Fatal("Error reading CSV data:", err)
				break
			}
			csvData = append(csvData, removeEmptyStrings(record))
		}
		gameMap.Layers[i].Data.Content = csvData
	}

	return gameMap
}

type Tileset struct {
	Version      string `xml:"version,attr" json:"version"`
	TiledVersion string `xml:"tiledversion,attr" json:"tiledVersion"`
	Name         string `xml:"name,attr" json:"name"`
	TileWidth    int    `xml:"tilewidth,attr" json:"tileWidth"`
	TileHeight   int    `xml:"tileheight,attr" json:"tileHeight"`
	TileCount    int    `xml:"tilecount,attr" json:"tileCount"`
	Columns      int    `xml:"columns,attr" json:"columns"`
	Image        Image  `xml:"image" json:"image"`
	Tiles        []Tile `xml:"tile" json:"tiles"`
}

type Image struct {
	Source string `xml:"source,attr" json:"source"`
	Width  int    `xml:"width,attr" json:"width"`
	Height int    `xml:"height,attr" json:"height"`
}

type Tile struct {
	ID         int        `xml:"id,attr" json:"id"`
	Properties []Property `xml:"properties>property" json:"properties"`
}

type Property struct {
	Name  string `xml:"name,attr" json:"name"`
	Type  string `xml:"type,attr" json:"type"`
	Value string `xml:"value,attr" json:"value"`
}

func GetTilesetData(tileBytes []byte) Tileset {
	var tileset Tileset
	err := xml.Unmarshal(tileBytes, &tileset)
	if err != nil {
		log.Fatal("Error unmarshaling TMX file:", err)
	}

	return tileset
}
