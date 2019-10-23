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
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err) // Log and exit
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Trim(strings.Split(scanner.Text(), com)[0], " \t\n")
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, factDeclar) {
			env.initialFacts = strings.Split(strings.TrimPrefix(line, factDeclar), "")
		} else if strings.HasPrefix(line, queryDeclar) {
			env.queries = strings.Split(strings.TrimPrefix(line, queryDeclar), "")
		} else {
			env.rules = append(env.rules, line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
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
 * Initialize env.allFacts all mentioned facts from file statements
 */
func initAllFacts() {
	// list from initial facts
	env.allFacts = make([]string, len(env.initialFacts))
	copy(env.allFacts, env.initialFacts)

	// list from query facts
	for _, query := range env.queries {
		if !stringInSlice(query, env.allFacts) {
			env.allFacts = append(env.allFacts, query)
		}
	}

	// list from statement facts
	for _, rule := range env.rules {
		for _, token := range rule {
			if !stringInSlice(string(token), env.allFacts) && stringInSlice(string(token), strings.Split(factSymbol, "")) {
				env.allFacts = append(env.allFacts, string(token))
			}
		}
	}
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
		} else if line == "run" {
			break
		}
		line = strings.Trim(strings.Split(strings.Trim(line, " "), com)[0], " \t\n")
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, factDeclar) {
			env.initialFacts = strings.Split(strings.TrimPrefix(line, factDeclar), "")
		} else if strings.HasPrefix(line, queryDeclar) {
			env.queries = strings.Split(strings.TrimPrefix(line, queryDeclar), "")
		} else {
			env.rules = append(env.rules, line)
		}
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
