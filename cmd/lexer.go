/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   lexer.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/10/30 17:52:13 by jmonneri          #+#    #+#             */
/*   Updated: 2019/10/30 17:52:55 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"log"
	"os"
	"strings"
)

/*
 * Check any kind of input line
 * Returns true if all good
 */
func lexer(line string) {
	if line == "" {

	} else if strings.HasPrefix(line, factDeclar) {
		factLexer(line[1:])
	} else if strings.HasPrefix(line, queryDeclar) {
		queryLexer(line[1:])
	} else {
		ruleLexer(line)
	}
}

/*
 * Check initial facts declared after symbol '='
 */
func factLexer(line string) {
	for _, c := range line {
		if !(charInString(c, factSymbol)) {
			log.Fatalf("Fact initialization contains illegal character '%c'.\n", c)
			os.Exit(1)
		}
	}
}

/*
 * Check query facts declared after symbol '?'
 */
func queryLexer(line string) {
	for _, c := range line {
		if !(charInString(c, factSymbol)) {
			log.Fatalf("Query contains illegal character '%c'.\n", c)
			os.Exit(1)
		}
	}
}

/*
 * Check default rule declaration
 */
func ruleLexer(line string) {
	// Check characters are all legal
	for _, c := range line {
		if !(charInString(c, openBra+closeBra+not+and+or+xor+imp+ioi+factSymbol)) {
			log.Fatalf("Rule contains illegal character '%c'.\n", c)
			os.Exit(1)
		}
	}

	// Check brackets
	var bracketCount = 0
	for _, c := range line {
		if string(c) == openBra {
			bracketCount++
		}
		if string(c) == closeBra {
			bracketCount--
		}
	}
	if bracketCount != 0 {
		log.Fatal("Syntax error on brackets.\n")
		os.Exit(1)
	}

	// Check rule is divided left and right for imp and ioi symbols
	splitImp := strings.Split(line, imp)
	splitIoi := strings.Split(line, ioi)
	if len(splitImp) != 2 && len(splitIoi) != 2 {
		log.Fatal("Rule doesn't respect syntax around => or <=> symbol.\n")
		os.Exit(1)
	}

	// Check each symbol is followed by a legal symbol
	for i := 0; i < len(line)-1; i++ {
		// for open bracket
		if string(line[i]) == openBra && !charInString(rune(line[i+1]), closeBra+factSymbol+not) {
			log.Fatalf("Rule contains illegal character '%c' after '%c'.\n", line[i+1], line[i])
			os.Exit(1)
		}

		// for close bracket
		if string(line[i]) == closeBra && !charInString(rune(line[i+1]), and+or+xor+imp+ioi) {
			log.Fatalf("Rule contains illegal character '%c' after '%c'.\n", line[i+1], line[i])
			os.Exit(1)
		}

		// for not
		if string(line[i]) == not && !charInString(rune(line[i+1]), openBra+factSymbol) {
			log.Fatalf("Rule contains illegal character '%c' after '%c'.\n", line[i+1], line[i])
			os.Exit(1)
		}

		// for operators and, or, xor
		operators := and + or + xor
		if charInString(rune(line[i]), operators) && !charInString(rune(line[i+1]), openBra+not+factSymbol) {
			log.Fatalf("Rule contains illegal character '%c' after '%c'.\n", line[i+1], line[i])
			os.Exit(1)
		}

		// for imp and ioi
		if strings.HasPrefix(line[i:], ioi) && !charInString(rune(line[i+3]), openBra+not+factSymbol) {
			log.Fatalf("Rule contains illegal character '%c' after '%c'.\n", line[i+3], line[i])
			os.Exit(1)
		} else if strings.HasPrefix(line[i:], imp) && !charInString(rune(line[i+2]), openBra+not+factSymbol) {
			log.Fatalf("Rule contains illegal character '%c' after '%c'.\n", line[i+2], line[i])
			os.Exit(1)
		}

		// for factSymbol
		if charInString(rune(line[i]), factSymbol) && !charInString(rune(line[i+1]), (openBra+closeBra+and+or+xor+imp+ioi)) {
			log.Fatalf("Rule contains illegal character '%c' after '%c'.\n", line[i+1], line[i])
			os.Exit(1)
		}
	}
}
