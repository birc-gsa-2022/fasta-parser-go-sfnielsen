package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	//########################################################
	sequences := make([]string, 0)
	fileScanner := bufio.NewScanner(fastaFile)

	//scan file line by line
	for fileScanner.Scan() {
		line := fileScanner.Text()

		if len(line) == 0 {
			continue
		}

		//handle 'name of sequence' cases
		if line[0:1] == ">" {
			sequences = append(sequences, "")
			sequences[len(sequences)-1] = strings.TrimSpace(line[1:])

			//handle 'sequence' cases
		} else {
			sequences[len(sequences)-1] = sequences[len(sequences)-1] + strings.TrimSpace(line)
		}
	}

	// #######################################

	var coordFile = os.Stdin
	if len(os.Args) == 3 && os.Args[2] != "-" {
		coordFile, err = os.Open(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err.Error())
			os.Exit(1)
		}
		defer coordFile.Close()
	}

	//######################

	output := ""
	fileScannerCoord := bufio.NewScanner(coordFile)

	for fileScannerCoord.Scan() {
		lineCoord := fileScannerCoord.Text()
		if len(lineCoord) == 0 {
			continue
		}
		split_str := strings.Fields(lineCoord)
		if len(split_str) != 3 {
			fmt.Println("Error, length of coordinates")
		}
		for i, v := range split_str {
			split_str[i] = strings.TrimSpace(v)
		}

		for _, v := range sequences {

			if strings.HasPrefix(v, split_str[0]) {

				basepair := strings.TrimSpace(strings.TrimPrefix(v, split_str[0]))
				low, er1 := strconv.Atoi(split_str[1])
				high, er2 := strconv.Atoi(split_str[2])
				if er1 != nil || er2 != nil {
					fmt.Printf("Error")

				}
				output = output + basepair[low-1:high-1] + "\n"
				break
			}

		}
	}
	fmt.Print(output)

	//#########################################

}
