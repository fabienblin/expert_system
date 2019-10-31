/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   parser.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/10/30 17:52:26 by jmonneri          #+#    #+#             */
/*   Updated: 2019/10/30 17:53:00 by jmonneri         ###   ########.fr       */
/*                                                                            */
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

	env.allFacts = make(map[string]int)

	// list from query facts
	for _, f := range env.queries {
		env.allFacts[f] = unknownF
	}

	// list from statement facts
	for _, rule := range env.rules {
		for _, f := range rule {
			if charInString(f, factSymbol) {
				env.allFacts[string(f)] = unknownF
			}
		}
	}

	// list from initial facts
	for _, f := range env.initialFacts {
		env.allFacts[f] = trueF
	}
}
