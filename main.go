package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

	fmt.Println(globalEmotes)

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
	img, _ := os.Create(imageName)
	defer img.Close()

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	b, _ := io.Copy(img, resp.Body)
	fmt.Println("File size: ", b)
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
