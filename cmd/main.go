package main

import (
	"fmt"
	"log"
	"os"
)

func initEnv() {
	env.rules = nil
	env.initialFacts = nil
	env.queries = nil
	env.trees = nil
	env.factList = make(map[string]*fact)
	env.tree = nil
}

func main() {
	initEnv()

	if len(os.Args) == 1 { // dynamic ruleset
		fmt.Printf("Using dynamic mode. \nPlease write the rules followed by initial facts then your query.\nType 'exit' to stop.\nType 'run' to run inference engine.\n")
		for {
			parseDynamic()
			printNode(env.tree, 4)
			engine()
			fmt.Printf("You may redefine known facts to retry the query.\n")
			env.initialFacts = nil
		}
	} else if len(os.Args) == 2 { // file ruleset
		parseFile(os.Args[1])
		printNode(env.tree, 4)
		engine()
	} else { // error
		log.Fatal("Error. Retry later ...\n")
		os.Exit(1)
	}
}
