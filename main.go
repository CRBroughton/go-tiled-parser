package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/crbroughton/go-tiled-parser/parser"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <path_to_map.xml> <output_json_path>")
		return
	}

	// TODO - Refactor into generic reader
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

	gameMap := parser.GetMapData(byteValue)

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
