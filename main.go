package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Property struct {
	SchemaName   string `json:"schema"`
	PropertyName string `json:"property"`
	Type         string `json:"type"`
	Description  string `json:"description"`
}

func main() {
	// Read JSON data from file
	sb, err := os.ReadFile("./apps1.json")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Define map to unmarshal JSON data into
	var OpenApiSchema map[string]interface{}

	// Unmarshal JSON data into map
	err = json.Unmarshal(sb, &OpenApiSchema)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Access the schemas map
	schemasMap, ok := OpenApiSchema["components"].(map[string]interface{})["schemas"].(map[string]interface{})
	if !ok {
		fmt.Println("Error: Unable to access schemas map")
		return
	}

	// Create a slice to hold the properties
	var properties []Property

	// Iterate over the schemas
	for schemaName, schema := range schemasMap {
		// Access the schema map
		schemaMap, ok := schema.(map[string]interface{})
		if !ok {
			fmt.Println("Error: Unable to assert schema to map[string]interface{}")
			continue
		}

		// Process properties
		processProperties(schemaName, schemaMap, &properties, OpenApiSchema)
	}

	// Marshal properties into JSON
	jsonData, err := json.MarshalIndent(properties, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling properties to JSON:", err)
		return
	}

	// Print JSON data
	fmt.Println(string(jsonData))
}

// Function to process properties
func processProperties(schemaName string, schemaMap map[string]interface{}, properties *[]Property, OpenApiSchema map[string]interface{}) {
	// Access the properties map within the schema
	propertiesMap, ok := schemaMap["properties"].(map[string]interface{})
	if !ok {
		fmt.Printf("Error: Unable to access properties map %v\n", schemaMap["properties"])
		return
	}

	// Iterate over the properties
	for propName, prop := range propertiesMap {
		// Access the property map
		propMap, ok := prop.(map[string]interface{})
		if !ok {
			fmt.Println("Error: Unable to assert property to map[string]interface{}")
			continue
		}

		// Create a Property object
		property := Property{
			SchemaName:   schemaName,
			PropertyName: propName,
			Type:         fmt.Sprintf("%v", propMap["type"]),
			Description:  fmt.Sprintf("%v", propMap["description"]),
		}

		// If type is not defined and $ref is present, follow the reference
		if property.Type == "" {
			ref, ok := propMap["$ref"].(string)
			if ok {
				// Follow the reference
				processReference(schemaName, ref, properties, OpenApiSchema)
			}
		}

		// Append property to properties slice
		*properties = append(*properties, property)
	}
}

// Function to process reference
func processReference(schemaName string, ref string, properties *[]Property, OpenApiSchema map[string]interface{}) {
	// Extract the path from the reference
	refParts := strings.Split(ref, "/")
	if len(refParts) < 3 {
		fmt.Println("Error: Invalid reference format")
		return
	}
	path := refParts[2]

	// Access the referenced schema
	refSchema, ok := OpenApiSchema[path].(map[string]interface{})
	if !ok {
		fmt.Println("Error: Unable to access referenced schema")
		return
	}

	// Process properties of the referenced schema
	processProperties(path, refSchema, properties, OpenApiSchema)
}
