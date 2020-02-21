/* ************************************************************************** */
/*                                                          LE - /            */
/*                                                              /             */
/*   parser.go                                        .::    .:/ .      .::   */
/*                                                 +:+:+   +:    +:  +:+:+    */
/*   By: jojomoon <jojomoon@student.le-101.fr>      +:+   +:    +:    +:+     */
/*                                                 #+#   #+    #+    #+#      */
/*   Created: 2019/10/30 17:52:26 by jmonneri     #+#   ##    ##    #+#       */
/*   Updated: 2020/01/22 11:25:02 by jojomoon    ###    #+. /#+    ###.fr     */
/*                                                         /                  */
/*                                                        /                   */
/* ************************************************************************** */
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
 * UNUSED
 * Main parse function takes program args and defines exec mode
 */
func parse() {
	if len(os.Args) == 1 { // dynamic ruleset
		parseDynamic()
	} else if len(os.Args) == 2 { // file ruleset
		parseFile(os.Args[1])
	} else { // error
		log.Fatal("Error. Retry later ...\n")
		os.Exit(1)
	}

	initAllFacts()
	buildTree()
}

/*
 * Parse file and initialize the env global variable
 */
func parseFile(fileName string) {
	var line string

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err) // Log and exit
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
}

/*
 * Parse line dynamically
 */
func parseDynamic() {
	var line string

	scanner := bufio.NewScanner(os.Stdin)
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
		fmt.Printf("Warning : Incomplete data from input.\n")
	}

	initAllFacts()
	buildTree()
}

/*
 * Parse and lex any line
 */
func parseLine(line string) {
	line = strings.Split(line, com)[0]
	line = strings.Replace(line, " ", "", -1)
	line = strings.Replace(line, "\t", "", -1)

	// lex
	lexer(line)

	if line == "" {
		return
	}

	// parse
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

	// list from query facts
	for _, f := range env.queries {
		if _, ok := env.factList[string(f)]; !ok {
			env.factList[string(f)] = newFact()
		}
		env.factList[string(f)].op = string(f)
		env.factList[string(f)].isKnown = false
		env.factList[string(f)].value = defaultF
		env.factList[string(f)].fixed = false
	}

	// list from statement facts
	for _, rule := range env.rules {
		for _, f := range rule {
			if charInString(f, factSymbol) {
				if _, ok := env.factList[string(f)]; !ok {
					env.factList[string(f)] = newFact()
				}
				env.factList[string(f)].op = string(f)
				env.factList[string(f)].isKnown = false
				env.factList[string(f)].value = defaultF
				env.factList[string(f)].fixed = false
			}
		}
	}

	// list from initial facts
	for _, f := range env.initialFacts {
		if _, ok := env.factList[string(f)]; !ok {
			env.factList[string(f)] = newFact()
		}
		env.factList[string(f)].op = string(f)
		env.factList[string(f)].isKnown = true
		env.factList[string(f)].value = trueF
		env.factList[string(f)].fixed = true
	}
}
