package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	bs, err := os.ReadFile("/IdeaProjects/gocourse/data/test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello World", string(bs))
}
