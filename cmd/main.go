package main

import (
	"fmt"
	"log"

	"github.com/blevesearch/bleve/v2"
)

func main() {
	index, err := bleve.Open("example.bleve")
	if err != nil {
		log.Fatal(err)	
	}

	query := bleve.NewMatchQuery("cupiditate")
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, err := index.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	for _, hit := range searchResult.Hits {
		fmt.Printf("ID: %s, Score: %f\n", hit.ID, hit.Score)
		for field, value := range hit.Fields {
			fmt.Printf("%s: %v\n", field, value)
		}
	}

}
