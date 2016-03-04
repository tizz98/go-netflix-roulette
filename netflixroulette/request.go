package netflixroulette

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	apiUrl string = "http://netflixroulette.net/api/api.php"
)

func getMediaFromUrl(requestUrl string) *Media {
	var media Media

	resp, err := http.Get(requestUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&media)

	if err != nil {
		panic(err)
	}

	return &media
}

func getMediaListFromUrl(requestUrl string) []Media {
	var media []Media

	resp, err := http.Get(requestUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&media)

	if err != nil {
		panic(err)
	}

	return media
}

func GetMediaByName(name string) *Media {
	return getMediaFromUrl(fmt.Sprintf("%s?title=%s", apiUrl, url.QueryEscape(name)))
}

func GetMediaByNameAndYear(name string, year int) *Media {
	return getMediaFromUrl(fmt.Sprintf("%s?title=%s&year=%s", apiUrl, url.QueryEscape(name), string(year)))
}

func GetDirectorByName(name string) *Director {
	director := Director{Name: name}
	media := getMediaListFromUrl(fmt.Sprintf("%s?director=%s", apiUrl, url.QueryEscape(name)))

	director.MediaList = media

	return &director
}

func GetActorByName(name string) *Actor {
	actor := Actor{Name: name}
	media := getMediaListFromUrl(fmt.Sprintf("%s?actor=%s", apiUrl, url.QueryEscape(name)))

	actor.MediaList = media

	return &actor
}
