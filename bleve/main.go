package main

import (
	"fmt"
	"log"

	"github.com/blevesearch/bleve/v2"
)

func main() {
	// Create a new index mapping
	indexMapping := bleve.NewIndexMapping()

	// Create a document mapping for the blog posts
	blogMapping := bleve.NewDocumentMapping()

	// Create a field mapping for the title
	titleMapping := bleve.NewTextFieldMapping()
	titleMapping.Analyzer = "en"
	blogMapping.AddFieldMappingsAt("title", titleMapping)

	// Create a field mapping for the author
	authorMapping := bleve.NewTextFieldMapping()
	authorMapping.Analyzer = "keyword"
	blogMapping.AddFieldMappingsAt("author", authorMapping)

	// Create a field mapping for the body
	bodyMapping := bleve.NewTextFieldMapping()
	bodyMapping.Analyzer = "en"
	blogMapping.AddFieldMappingsAt("body", bodyMapping)

	// Create a field mapping for the date
	dateMapping := bleve.NewDateTimeFieldMapping()
	blogMapping.AddFieldMappingsAt("date", dateMapping)

	// Create a field mapping for the tags
	tagsMapping := bleve.NewTextFieldMapping()
	tagsMapping.Analyzer = "keyword"
	blogMapping.AddFieldMappingsAt("tags", tagsMapping)

	// Add the blog mapping to the index mapping
	indexMapping.AddDocumentMapping("blog", blogMapping)

	// Create a new index
	index, err := bleve.New("example.bleve", indexMapping)
	if err != nil {
		log.Fatal(err)
	}
	defer index.Close()

	// Example blog post
	blogPost := map[string]interface{}{
		"title":  "The Benefits of Functional Programming",
		"author": "John Doe",
		"body":   "Functional programming is a programming paradigm that treats computation as the evaluation of mathematical functions...",
		"date":   "2023-08-22",
		"tags":   []string{"programming", "functional programming", "software development"},
	}

	// Index the blog post
	err = index.Index("blog_1", blogPost)
	if err != nil {
		log.Fatal(err)
	}

	// Search for "functional programming"
	query := bleve.NewMatchQuery("Functional	")
	searchRequest := bleve.NewSearchRequest(query)

	searchResults, err := index.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Print the search results
	fmt.Println("Search Results:")
	for _, hit := range searchResults.Hits {
		fmt.Printf("Document ID: %s, Score: %f\n", hit.ID, hit.Score)
	}
}
