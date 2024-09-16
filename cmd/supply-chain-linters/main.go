package main

import (
	"fmt"
	"log"
	"os"

	"github.com/illume/supply-chain-linters/pkg/linter"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s [path to scan]\n", os.Args[0])
	}

	path := os.Args[1]

	// Run the linter on the provided path
	err := linter.Run(path)
	if err != nil {
		log.Fatalf("Linter failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Linter completed successfully")
	os.Exit(0)
}
