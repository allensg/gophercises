package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type (
	// Adventure defines the struct for a json adventure
	Adventure struct {
		Segment []Segment
	}

	Segment struct {
		Title   string    `json:"title"`
		Story   []string  `json: story`
		Options []Options `json: options`
	}

	Options struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	}
)

func main() {

	// read the json file
	moduleTitle := ""

	adventure, err := decode()

	for _, segment := range adventure.Segment {
		fmt.Println(segment)
	}

	fmt.Println(adventure)
	if err != nil {

		// if error is not nil
		// print error
		fmt.Println(err)
	}

	// prompt the user
	fmt.Print("Welcome to the choose your own adventure terminal input because webapps aren't worth the time")

	fmt.Print(fmt.Sprintf("You have chosen the %s module ", moduleTitle))

	//define game loop

}

func decode() (*Adventure, error) {
	// adventure := make(map[string]interface{})
	file, err := os.ReadFile("gopher.json")
	adventure := &Adventure{}
	if err != nil {
		return adventure, err
	}
	err = json.Unmarshal(file, &adventure)

	if err != nil {
		return adventure, err
	}

	return adventure, nil
}
