package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// _ = "breakpoint"
	resp, err := http.Get("https://twitchemotes.com/api_cache/v2/global.json")

	if err != nil {
		fmt.Println(err)
	}

	if resp.Body == nil {
		panic("ahhh")
	}

	var globalEmotes global

	dec := json.NewDecoder(resp.Body)

	err2 := dec.Decode(&globalEmotes)

	if err2 != nil {
		fmt.Println(err)
	}

	for key, _ := range globalEmotes.Emotes {
		// fmt.Println(globalEmotes.Emotes[key])
		emoteId := globalEmotes.Emotes[key].ImageID
		finalUrl := strings.Replace(globalEmotes.Template.Large, "{image_id}", strconv.Itoa(emoteId), 1)
		fmt.Println(finalUrl)
		saveImage(finalUrl, fmt.Sprintf("/Users/nma/go/src/github.com/nicolas-martin/twitchImgGetter/emotes/%s.png", key))
	}

}

func getGlobalEmotes() *global {
	resp, err := http.Get("https://twitchemotes.com/api_cache/v2/global.json")

	if err != nil {
		fmt.Println(err)
	}

	if resp.Body == nil {
		panic("ahhh")
	}

	var globalEmotes global

	dec := json.NewDecoder(resp.Body)

	err2 := dec.Decode(&globalEmotes)

	if err2 != nil {
		fmt.Println(err)
	}

	return &globalEmotes
}

func saveImage(url string, imageName string) {
	// url := "http://i.imgur.com/m1UIjW1.jpg"

	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}

	defer response.Body.Close()

	//open a file for writing
	file, err := os.Create(imageName)
	if err != nil {
		log.Fatal(err)
	}
	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	fmt.Println("Success!")
}

type global struct {
	Meta struct {
		GeneratedAt time.Time `json:"generated_at"`
	} `json:"meta"`
	Template struct {
		Small  string `json:"small"`
		Medium string `json:"medium"`
		Large  string `json:"large"`
	} `json:"template"`
	Emotes map[string]Emote
}

type Emote struct {
	Description string      `json:"description"`
	ImageID     int         `json:"image_id"`
	FirstSeen   interface{} `json:"first_seen"`
}
