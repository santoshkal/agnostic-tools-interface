package main

import (
	"encoding/json"
	"fmt"

	"github.com/meilisearch/meilisearch-go"
)

// Wrapper struct around SearchResult
type SearchResultWrapper struct {
	Result *meilisearch.SearchResponse
}

func main() {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://localhost:7700",
		APIKey: "aSampleMasterKey",
	})

	// Add Schema to DB

	// jsonFile, _ := os.Open("apps.json")
	// defer jsonFile.Close()

	// // byteValue, _ := io.ReadAll(jsonFile)
	// var appMap map[string]interface{}
	// err := json.NewDecoder(jsonFile).Decode(&appMap)
	// if err != nil {
	// 	fmt.Println("Error decoding JSON:", err)
	// 	return
	// }

	// apps := []map[string]interface{}{appMap}

	// resp, err := client.Index("apps").AddDocuments(apps)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Document apps Added: %#v\n", resp.Status)
	// fmt.Printf("Document apps Added: %#v\n", resp.TaskUID)
	// fmt.Printf("Document apps Added: %#v\n", resp.IndexUID)
	// fmt.Printf("Document apps Added: %#v\n", resp.Type)

	s, err := client.Index("apps").Search("", &meilisearch.SearchRequest{
		HitsPerPage: 100,
		Page:        1,
		// AttributesToRetrieve: []string{"type", "description"},
		AttributesToRetrieve: []string{"components.schemas.io.k8s.api.apps.v1.*.type",
			"components.schemas.io.k8s.api.apps.v1.*.description"},
		// Filter: [][]string{
		// 	[]string{"name = \"limit\""},
		// 	[]string{"description: limit is a maximum number of responses to return for a list call"},
		// },
	})
	if err != nil {
		fmt.Printf("Failed to search: %v", err)
	}
	// Wrap the SearchResult
	wrappedResult := SearchResultWrapper{Result: s}
	jsonData, err := json.MarshalIndent(wrappedResult, "", "    ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v", err)
		return
	}
	if err != nil {
		fmt.Printf("Failed to marshal %v search results: %v", jsonData, err)
	}
	fmt.Println(s.Hits)
	fmt.Printf("Search results: %v\n", s.Hits)

	fmt.Printf("Search Results in JSON: %v\n", string(jsonData))

}
