package main

import (
	"log"
	"os"
)

func initEnv() {
	env.factList = make(map[string]*fact)
	env.initialFacts = nil
	env.trees = nil
}

func main() {

	if len(os.Args) == 1 { // dynamic ruleset
		for {
			initEnv()
			parseDynamic()

			initAllFacts()
			buildTree()

			// printNode(env.tree, 8)
			engine()
			printNode(env.tree, 8)
		}
	} else if len(os.Args) == 2 { // file ruleset
		initEnv()
		parseFile(os.Args[1])

		initAllFacts()
		buildTree()

		// printNode(env.tree, 8)
		engine()
		printNode(env.tree, 8)
	} else { // error
		log.Fatal("Error. Retry later ...\n")
		os.Exit(1)
	}
}
