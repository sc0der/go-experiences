package main

import (
	"fmt"
)

type Movie struct {
	Actors      []string
	Rating      float32
	ReleaseYear int
	Title       string
	SayHi       func()
}

var episodeIV = Movie{
	Title:       "Star Wars: A New Hope",
	Rating:      5.0,
	ReleaseYear: 1977,
	SayHi:       func() { fmt.Println("Salom") },
}

func main() {
	fmt.Println(episodeIV.Title)       // Outputs "Int:
	fmt.Println(episodeIV.Rating)      // Outputs "Int:
	fmt.Println(episodeIV.ReleaseYear) // Outputs "Int:
	episodeIV.SayHi()                  // Outputs "Int:"
}
