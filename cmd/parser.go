package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
 * Parse file and initialize the env global variable
 */
func parseFile(fileName string) {
	var line string

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		parseLine(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if !(env.initialFacts != nil && env.queries != nil && env.rules != nil) {
		log.Fatal("Incomplete data from file.\n")
		os.Exit(1)
	}
	initAllFacts()
	buildTree()
}

/*
 * Parse line dynamically
 */
func parseDynamic() {
	var line string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Using dynamic mode. \nPlease write the rules followed by initial facts then your query.\nType 'exit' to stop.\nType 'run' to run inference engine.\n")
	for scanner.Scan() {
		line = scanner.Text()
		if line == "exit" {
			os.Exit(0)
		} else if (env.initialFacts != nil && env.queries != nil && env.rules != nil) || line == "run" {
			break
		}
		parseLine(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if !(env.initialFacts != nil && env.queries != nil && env.rules != nil) {
		log.Fatal("Incomplete data from input.\n")
		os.Exit(1)
	}
	initAllFacts()
	buildTree()
}

func parseLine(line string) {
	line = strings.Split(line, com)[0]
	line = strings.Replace(line, " ", "", -1)
	line = strings.Replace(line, "\t", "", -1)

	lexer(line)

	if line == "" {
		return
	}

	if strings.HasPrefix(line, factDeclar) {
		env.initialFacts = strings.Split(strings.TrimPrefix(line, factDeclar), "")
	} else if strings.HasPrefix(line, queryDeclar) {
		env.queries = strings.Split(strings.TrimPrefix(line, queryDeclar), "")
	} else {
		env.rules = append(env.rules, line)
	}
}

/*
 * Initialize env.allFacts from all mentioned facts
 */
func initAllFacts() {
	// list from initial facts
	env.allFacts = make([]string, len(env.initialFacts))
	copy(env.allFacts, env.initialFacts)

	// list from query facts
	for _, fact := range env.queries {
		if !stringInSlice(fact, env.allFacts) {
			env.allFacts = append(env.allFacts, fact)
		}
	}

	// list from statement facts
	for _, statement := range env.rules {
		for _, stmt := range statement {
			if !stringInSlice(string(stmt), env.allFacts) && stringInSlice(string(stmt), strings.Split(factSymbol, "")) {
				env.allFacts = append(env.allFacts, string(stmt))
			}
		}
	}
}
