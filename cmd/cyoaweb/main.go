package main

import (
	"flag"
	"fmt"
	"log"
	"mymodule/golang/yourownadventure"
	"net/http"
	"os"
)

func main() {

	port := flag.Int("port", 3000, "the port to start the server")
	fileName := flag.String("file", "gopher.json", "json file where the story is stored.")
	flag.Parse()

	os.ReadFile(*fileName)

	f, err := os.Open(*fileName)
	if err != nil {
		fmt.Println("Couldnt open the provided file.")
	}

	story, err := yourownadventure.JsonStory(f)
	if err != nil {
		fmt.Println("Couldnt decode the story provided.")
	}

	h := yourownadventure.NewHandler(story, nil)
	fmt.Printf("Starting server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
