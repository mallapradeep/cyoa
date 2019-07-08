package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/gophercises/cyoa"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	//opening up that file
	f, err := os.Open(*filename)
	// chk if there is any error
	if err != nil {
		panic(err)
	}
	//decode that file, we pass in a file to that decoder
	d := json.NewDecoder(f)
	var story cyoa.Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", story)

}
