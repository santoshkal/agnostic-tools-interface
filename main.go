package main

import (
	"fmt"

	"github.com/meilisearch/meilisearch-go"
)

func main() {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://localhost:7700",
		APIKey: "aSampleMasterKey",
	})

	_, err := client.CreateIndex(&meilisearch.IndexConfig{
		Uid:        "apps",
		PrimaryKey: "id",
	})
	if err != nil {
		fmt.Printf("Failed to create index: %v", err)
	}

	// jsonFile, _ := os.Open("apps.json")
	// defer jsonFile.Close()

	// // byteValue, _ := io.ReadAll(jsonFile)
	// var appMap map[string]interface{}
	// err = json.NewDecoder(jsonFile).Decode(&appMap)
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

	fa := []string{"type", "description"}
	client.Index("apps").UpdateFilterableAttributes(&fa)

	// i, err := client.GetStats("apps")
	// if err != nil {
	// 	fmt.Printf("Failed to get Index: %v", err)
	// }

	f, err := client.Index("apps").GetFilterableAttributes()
	if err != nil {
		fmt.Printf("Error fetching filterable attributes: %v", err)
	}

	s, err := client.Index("apps").Search("DeploymentList", &meilisearch.SearchRequest{
		AttributesToRetrieve: []string{"type", "description"},
		// Filter: [][]string{
		// 	[]string{"name = \"limit\""},
		// 	[]string{"description: limit is a maximum number of responses to return for a list call"},
		// },
	})
	if err != nil {
		fmt.Printf("Failed to search: %v", err)
	}
	// t, err := client.GetTask(3)
	// if err != nil {
	// 	fmt.Printf("Failed to get task: %v", err)
	// }
	// fmt.Printf("Found Index: %v\n", i.UID)
	fmt.Printf("Found Search: %v\n", s.Hits)

	fmt.Printf("Found Attributes: %v\n", f)
	// fmt.Printf("Found Items: %v\n", a.IndexUID)

}
