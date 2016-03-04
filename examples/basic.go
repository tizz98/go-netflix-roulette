package main

import (
	"fmt"
	"github.com/tizz98/go-netflix-roulette/netflixroulette"
)

func main() {
	media := netflixroulette.GetMediaByName("The Avengers")
	fmt.Printf("%s is a %s\n", media.ShowTitle, media.GetReadableMediaType())

	director := netflixroulette.GetDirectorByName("Quentin Tarantino")
	fmt.Printf("%s has %d media:\n", director.Name, len(director.MediaList))

	for _, media := range director.MediaList {
		fmt.Printf("\t%s (%s) - %s\n", media.ShowTitle, media.ReleaseYear, media.Summary)
	}
	fmt.Println("")

	actor := netflixroulette.GetActorByName("Nicolas Cage")
	fmt.Printf("%s has %d media:\n", actor.Name, len(actor.MediaList))

	for _, media := range actor.MediaList {
		fmt.Printf("\t%s (%s) - %s\n", media.ShowTitle, media.ReleaseYear, media.Summary)
	}
}
