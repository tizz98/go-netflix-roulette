# Go Netflix Roulette
A wrapper for the [Netflix Roulette Api](http://netflixroulette.net/api/), written in [Go](https://golang.org).

## Install
`go get github.com/tizz98/go-netflix-roulette`


## Example
```go
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
```
