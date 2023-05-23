package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Taiki130/lexi-search/internal/dictionary"
)

func main() {
	word := flag.String("define", "", "Specify a word to get its definition")
	flag.Parse()

	if *word == "" {
		fmt.Println("Please specify a word using the -define flag")
		return
	}

	definition, err := dictionary.GetDefinition(*word)
	if err != nil {
		log.Fatalf("Failed to get definition: %v", err)
	}

	fmt.Printf("Definition of %s:\n%s\n", *word, definition)
}
