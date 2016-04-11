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

Outputs:
```
The Avengers is a Movie
Quentin Tarantino has 7 media:
	Kill Bill: Vol. 2 (2004) - The Bride has three left on her rampage list: Budd, Elle Driver and Bill himself. But when she arrives at Bill's house, she's in for a surprise.
	Kill Bill: Vol. 1 (2003) - An assassin is shot by her ruthless employer, Bill, and other members of their assassination circle. But she lives -- and plots her vengeance.
	Pulp Fiction (1994) - Weaving together three stories featuring a burger-loving hit man, his philosophical partner and a washed-up boxer, Quentin Tarantino influenced a generation of filmmakers with this crime caper's stylized, over-the-top violence and dark comic spirit.
	Jackie Brown (1997) - Jackie Brown is an aging flight attendant who smuggles cash on the side. But when she's busted and pressured to help with an investigation, she plans to play the opposing forces against each other and walk away with the dough.
	Reservoir Dogs (1992) - Quentin Tarantino's directorial debut is raw, violent, often mimicked -- and unforgettable. A botched robbery indicates a police informant, and the pressure mounts in the aftermath at a warehouse. Crime begets violence as the survivors unravel.
	Four Rooms (1995) - One mad New Year's Eve, an overwhelmed bellboy copes with witches and diabolical children, gets caught in the middle of a sour relationship and settles a bloody bet for members of a superstar's entourage.
	Django Unchained (2012) - Accompanied by a German bounty hunter, a freed slave named Django travels across America to free his wife from a sadistic plantation owner.

Nicolas Cage has 12 media:
	It Could Happen to You (1994) - In this charming romantic comedy based on a true story, a coffee-shop waitress gets a life-changing tip when a beat cop comes up short on pocket change and promises her half of his potential winnings from a lottery ticket.
	The Croods (2013) - When an earthquake obliterates their cave, an unworldly prehistoric family is forced to journey through unfamiliar terrain in search of a new home. But things for pessimistic dad Grug go from bad to worse when his daughter meets a clever cave boy.
	Christmas Carol: The Movie (2001) - Nicolas Cage, Kate Winslet and Simon Callow provide voices for this animated version of the classic Charles Dickens tale about the miserable Ebenezer Scrooge, who learns the true meaning of Christmas from three ghosts who confront him in the night. The film uses live-action sequences to bookend the story, and a mouse in the animated portion helps younger viewers follow the story. Winslet sings the film's main theme.
	Stolen (2012) - Master thief Will Montgomery is ready to leave his criminal past behind. But when his daughter is kidnapped, he has no choice but to reunite with his old partner in crime and pull off one last heist.
	Seeking Justice (2011) - After his wife is brutally raped, English teacher Will Gerard is approached by a man from a vigilante group who offers to exact revenge on the perpetrator. But once the deed is done, Will discovers that the group expects a certain favor in return.
	Trespass (2011) - A husband and wife find themselves pushed to their absolute limit when they're held for ransom by brutal thugs who invade their home. As tensions escalate and shocking revelations emerge, the couple is forced to take ever-more desperate measures.
	Frozen Ground (2013) - In this fact-based thriller, an Alaska state trooper pursuing a serial killer teams with a 17-year-old-prostitute who escaped the predator's clutches.
	Face/Off (1997) - An antiterrorism agent goes under the knife to acquire the likeness of a terrorist and gather details about a bombing plot.
	Rage (2014) - When his daughter is abducted, a respectable businessman with a violent past rounds up his old crew to help him find her -- by any means necessary.
	Moonstruck (1987) - In this slice-of-life comedy about the Italian American residents of a Brooklyn neighborhood, a strong widow falls in love with a one-handed baker.
	Joe (2013) - When ex-con Joe hires 15-year-old Gary to help clear trees for a lumber company, he doesn't expect to become a father figure for the abused boy.
	Leaving Las Vegas (1995) - An alcoholic moves to Las Vegas to drink himself to death and meets a prostitute who comes to love him, without trying to upset his nihilistic agenda.
```
