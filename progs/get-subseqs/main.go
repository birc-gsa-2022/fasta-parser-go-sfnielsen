package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s fasta\n", os.Args[0])
		os.Exit(1)
	}

	fastaFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}
	defer fastaFile.Close()

	fmt.Println("I should be processing the fasta file now!")

	var coordFile = os.Stdin
	if len(os.Args) == 3 && os.Args[2] != "-" {
		coordFile, err = os.Open(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err.Error())
			os.Exit(1)
		}
		defer coordFile.Close()
	}

	fmt.Println("I should be processing the coordinate stream now!")
}
