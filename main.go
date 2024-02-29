package main

import (
	// Compile preprocessing removes unused portions of imports
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"slices"
	"strings"
)

// Deserializes JSON to Go var
func decodeJSON(filePath string) map[string]interface{} {
	file, fileOpenError := os.Open(filePath)
	if fileOpenError != nil {
		log.Fatal("Error opening file: ", fileOpenError)
	}
	var fileData map[string]interface{}
	jsonDecodeError := json.NewDecoder(file).Decode(&fileData)
	if jsonDecodeError != nil {
		log.Fatal("Error decoding JSON: ", jsonDecodeError)
	}
	return fileData
}

// Crawls DBT manifest.json and checks for nodes and their dependencies
func createNodeMap(fileData map[string]interface{}) map[string]interface{} {
	nodeMap := make(map[string]interface{})
	for key := range fileData {
		if strings.Contains(key, "child_map") {
			nodeMap[key] = fileData[key]
		}
	}
	return nodeMap
}

// Removes any duplicate values from the nodeMap
func removeNodeDuplicates(nodeMap map[string]interface{}) map[string]interface{} {

	for key := range nodeMap {
		for model := range nodeMap[key].(map[string]interface{}) {
			var uniqueValues []string
			if reflect.TypeOf(nodeMap[key].(map[string]interface{})[model]) != nil {
				for _, element := range nodeMap[key].(map[string]interface{})[model].([]interface{}) {
					if !slices.Contains(uniqueValues, element.(string)) {
						uniqueValues = append(uniqueValues, element.(string))
					} else {
						continue
					}
				}
			}
			nodeMap[key].(map[string]interface{})[model] = uniqueValues
		}
	}
	return nodeMap
}

func main() {
	filePath := flag.String("file", "", "Path to manifest.json")
	flag.Parse()
	fileData := decodeJSON(*filePath)
	nodeMap := createNodeMap(fileData)
	nodeMap = removeNodeDuplicates(nodeMap)
	fmt.Println(nodeMap)
}
