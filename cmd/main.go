package main

import (
	"flag"
	"log"
	"os"
)

func initEnv() {
	env.factList = make(map[string]*fact)
	env.initialFacts = nil
	env.trees = nil
}

// flags f et v boolens

func main() {

	flagVerbose := flag.Bool("v", false, "verbose mode")
	flagForward := flag.Bool("f", false, "forward mode")

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 { // dynamic ruleset
		for {
			initEnv()
			parseDynamic()

			initAllFacts()
			buildTree()

			// printNode(env.tree, 8)
			engine(*flagVerbose, *flagForward)
			printNode(env.tree, 8)
		}
	} else if len(args) == 1 { // file ruleset

		initEnv()
		parseFile(args[0])

		initAllFacts()
		buildTree()

		printNode(env.tree, 8)
		engine(*flagVerbose, *flagForward)
		printNode(env.tree, 8)
	} else { // error
		log.Fatal("Error. Retry later ...\n")
		os.Exit(1)
	}
}
