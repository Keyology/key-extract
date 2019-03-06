package extract

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func GetLink(link string) {

	apiURL := "https://api.microlink.io?url="

	var payload Schema

	resp, err := http.Get(apiURL + link)

	if err != nil {

		log.Fatalf("Error retrieving data: %s\n", err)

	}

	json.NewDecoder(resp.Body).Decode(&payload)
	json.NewEncoder(os.Stdout).Encode(payload)

	resp.Body.Close()

}

type Schema struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    *Data  `json:"data,omitempty"`
}

type Data struct {
	Author      string    `json:"author,omitempty"`
	Lang        string    `json:"lang,omitempty"`
	Date        string    `json:"date,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Publisher   string    `json:"publisher,omitempty"`
	URL         string    `json:"url"`
	Image       *ImageURL `json:"image,omitempty"`
	Video       *ImageURL `json:"video,omitempty"`
	Logo        *ImageURL `json:"logo,omitempty"`
	Screenshot  *ImageURL `json:"screenshot,omitempty"`
}

type ImageURL struct {
	Width            int      `json:"width,omitempty"`
	Height           int      `json:"height,omitempty"`
	Size             int      `json:"size,omitempty"`
	SizePretty       string   `json:"size_pretty,omitempty"`
	Duration         int      `json:"duration,omitempty"`
	DurationPretty   string   `json:"duration_pretty,omitempty"`
	Type             string   `json:"type,omitempty"`
	URL              string   `json:"url"`
	Palette          []string `json:"palette,omitempty"`
	BackgroundColor  string   `json:"background_color,omitempty"`
	AlternativeColor string   `json:"alternative_color,omitempty"`
	Color            string   `json:"color,omitempty"`
}
