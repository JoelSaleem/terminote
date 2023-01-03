package main

import (
	"fmt"
	"log"

	"terminote/internal/notes_loader.go"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	NotesDir string `default:"/Users/joelsaleem/.terminote"`
}

func main() {
	var cfg Config
	err := envconfig.Process("terminote", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = notes_loader.LoadNotes(cfg.NotesDir)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("done")
}
