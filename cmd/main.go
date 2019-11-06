package main

import (
	"fmt"
	"os"
)

func main() {

	env.factList = make(map[string]*fact)

	if len(os.Args) == 1 { // dynamic ruleset
		parseDynamic()
	} else if len(os.Args) == 2 { // file ruleset
		parseFile(os.Args[1])
	} else { // error
		fmt.Println("Error. Retry later ...")
		os.Exit(1)
	}
	for _, tree := range env.trees {
		fmt.Printf("\nROOT : \n----------------------------\n")
		printNode(&tree, 4)
	}
	engine()
}
