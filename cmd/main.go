package main

import (
	"fmt"
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
		for {
			parseDynamic()
			printNode(env.tree, 4)
			engine()
			initEnv()
		}
	} else if len(os.Args) == 2 { // file ruleset
		parseFile(os.Args[1])
		printNode(env.tree, 4)
		engine()
	} else { // error
		fmt.Println("Error. Retry later ...")
		os.Exit(1)
	}
}
