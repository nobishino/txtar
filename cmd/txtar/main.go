package main

import (
	"log"
	"os"

	"github.com/nobishino/txtar"
)

func main() {
	if err := txtar.Write(os.Args[1:], os.Stdout); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
