package main

import (
	"fmt"
	"log"
	"math/rand"
	"search_engine/config"
	"search_engine/pkg/postgres"

	"github.com/bxcodec/faker/v3"
	_ "github.com/lib/pq"
)

type Movie struct {
	Title       string
	Director    string
	ReleaseYear int
	Genre       string
	Plot        string
	Actors      string
}

func main() {
	cnf := config.NewConfig()
	if err := cnf.Load(); err != nil {
		log.Fatal(err)
	}
	// Connect to your PostgreSQL database
	db, err := postgres.ConnectDB(cnf.Database)
	if err != nil {
		log.Fatal("Connection failed: ", err)
	}
	defer db.Close()

	// Generate and insert mock data
	for i := 0; i < 10000; i++ { // Generate 10,000 mock movies
		movie := generateMockMovie()
		_, err := db.Exec(`
            INSERT INTO movies (title, director, release_year, genre, plot, actors)
            VALUES ($1, $2, $3, $4, $5, $6)
        `, movie.Title, movie.Director, movie.ReleaseYear, movie.Genre, movie.Plot, movie.Actors)
		if err != nil {
			log.Printf("Error inserting movie: %v", err)
		}
		if i%100 == 0 {
			fmt.Printf("Inserted %d movies\n", i)
		}
	}

	fmt.Println("Mock data generation complete")
}

func generateMockMovie() Movie {
	return Movie{
		Title:       faker.Sentence(),
		Director:    faker.Name(),
		ReleaseYear: rand.Intn(50) + 1970, // Random year between 1970 and 2019
		Genre:       faker.Word(),
		Plot:        faker.Paragraph(),
		Actors:      faker.Name() + ", " + faker.Name() + ", " + faker.Name(),
	}
}
