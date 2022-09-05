package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s fasta\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}
	defer f.Close()

	//##################implementation###############

	output := ""
	fileScanner := bufio.NewScanner(f)

	//scan file line by line
	for fileScanner.Scan() {
		line := fileScanner.Text()

		if len(line) == 0 {
			continue
		}

		//handle 'name of sequence' cases
		if line[0:1] == ">" {
			if output != "" {
				output = output + "\n"
			}
			output = output + strings.TrimSpace(line[1:]) + "	"

			//handle 'sequence' cases
		} else {
			output = output + line
		}
	}
	fmt.Println(output)

}
