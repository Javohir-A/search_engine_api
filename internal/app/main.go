package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"

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
	db, err := sql.Open("postgres", "your_connection_string_here")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the movies table
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS movies (
            id SERIAL PRIMARY KEY,
            title TEXT,
            director TEXT,
            release_year INTEGER,
            genre TEXT,
            plot TEXT,
            actors TEXT
        )
    `)
	if err != nil {
		log.Fatal(err)
	}

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
