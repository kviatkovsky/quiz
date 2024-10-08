package main

import (
	"log"

	"github.com/kviatkovsky/quiz/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("could not start CLI: %v\n", err)
	}
}
