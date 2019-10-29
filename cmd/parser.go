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
	var rules, query, initial = false, false, false

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		line = strings.Trim(strings.Split(strings.Trim(line, " "), com)[0], " \t\n")
		line = strings.Replace(line, " ", "", -1)
		line = strings.Replace(line, "\t", "", -1)
		if !lexer(line) {
			os.Exit(1)
		}
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, factDeclar) {
			env.initialFacts = strings.Split(strings.TrimPrefix(line, factDeclar), "")
			initial = true
		} else if strings.HasPrefix(line, queryDeclar) {
			env.queries = strings.Split(strings.TrimPrefix(line, queryDeclar), "")
			query = true
		} else {
			env.rules = append(env.rules, line)
			rules = true
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if !(rules && query && initial) {
		log.Fatal("Incomplete data from file.\n")
		os.Exit(1)
	}
	initAllFacts()
	buildTree()
	fmt.Printf("rules : %q\n", env.rules)
	fmt.Printf("initialFacts : %q\n", env.initialFacts)
	fmt.Printf("queries : %q\n", env.queries)
	fmt.Printf("allFacts : %q\n", env.allFacts)
	for _, tree := range env.trees {
		fmt.Printf("\nROOT : \n----------------------------\n")
		printNode(&tree, 4)
	}
}

/*
 * Parse line dynamically
 */
func parseDynamic() {
	var line string
	var rules, query, initial = false, false, false

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Using dynamic mode. \nPlease write the rules followed by initial facts then your query.\nType 'exit' to stop.\nType 'run' to run inference engine.\n")
	for scanner.Scan() {
		line = scanner.Text()
		if line == "exit" {
			os.Exit(0)
		} else if (rules && query && initial) || line == "run" {
			break
		}
		line = strings.Trim(strings.Split(strings.Trim(line, " "), com)[0], " \t\n")
		line = strings.Replace(line, " ", "", -1)
		line = strings.Replace(line, "\t", "", -1)
		if !lexer(line) {
			os.Exit(1)
		}
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, factDeclar) {
			env.initialFacts = strings.Split(strings.TrimPrefix(line, factDeclar), "")
			initial = true
		} else if strings.HasPrefix(line, queryDeclar) {
			env.queries = strings.Split(strings.TrimPrefix(line, queryDeclar), "")
			query = true
		} else {
			env.rules = append(env.rules, line)
			rules = true
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if !(rules && query && initial) {
		log.Fatal("Incomplete data from input.\n")
		os.Exit(1)
	}
	initAllFacts()
	buildTree()
	fmt.Printf("rules : %q\n", env.rules)
	fmt.Printf("initialFacts : %q\n", env.initialFacts)
	fmt.Printf("queries : %q\n", env.queries)
	fmt.Printf("allFacts : %q\n", env.allFacts)
	for _, tree := range env.trees {
		fmt.Printf("\nROOT : \n----------------------------\n")
		printNode(&tree, 4)
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
