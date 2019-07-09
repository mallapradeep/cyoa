package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gophercises/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	filename := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	//opening up that file
	f, err := os.Open(*filename)
	// chk if there is any error
	if err != nil {
		panic(err)
	}

	// we r gonna use this Line a lot so put it as a func n call it later

	// //decode that file, we pass in a file to that decoder
	// d := json.NewDecoder(f)
	// var story cyoa.Story
	// if err := d.Decode(&story); err != nil {
	// 	panic(err)
	// }

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handler))
}
