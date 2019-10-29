package main

import (
	"fmt"
	"strings"
)

/*
 * Check any kind of input line
 * Returns true if all good
 */
func lexer(line string) bool {
	if line == "" {
		return true
	} else if strings.HasPrefix(line, factDeclar) {
		return factLexer(line[1:])
	} else if strings.HasPrefix(line, queryDeclar) {
		return queryLexer(line[1:])
	} else {
		return ruleLexer(line)
	}
}

/*
 * Check initial facts declared after symbol '='
 */
func factLexer(line string) bool {
	for _, c := range line {
		if !(charInString(c, factSymbol)) {
			fmt.Printf("Fact initialization contains illegal character '%c'.\n", c)
			return false
		}
	}
	return true
}

/*
 * Check query facts declared after symbol '?'
 */
func queryLexer(line string) bool {
	for _, c := range line {
		if !(charInString(c, factSymbol)) {
			fmt.Printf("Query contains illegal character '%c'.\n", c)
			return false
		}
	}
	return true
}

/*
 * Check default rule declaration
 */
func ruleLexer(line string) bool {
	// Check characters are all legal
	for _, c := range line {
		if !(charInString(c, openBra+closeBra+not+and+or+xor+imp+ioi+factSymbol)) {
			fmt.Printf("Rule contains illegal character '%c'.\n", c)
			return false
		}
	}

	// Check rule is divided left and right for imp and ioi symbols
	splitImp := strings.Split(line, imp)
	splitIoi := strings.Split(line, ioi)
	if len(splitImp) != 2 && len(splitIoi) != 2 {
		fmt.Printf("Rule doesn't respect synthax around => or <=> symbol.\n")
		return false
	}

	// Check each symbol is followed by a legal symbol
	for i := 0; i < len(line)-1; i++ {
		// for open bracket
		if string(line[i]) == openBra && !charInString(rune(line[i+1]), closeBra+factSymbol+not) {
			fmt.Printf("Rule contains illegal character '%c' after '%c'.\n", line[i+1], line[i])
			return false
		}

		// for close bracket
		if string(line[i]) == closeBra && !charInString(rune(line[i+1]), not+and+or+xor+imp+ioi+factSymbol) {
			fmt.Printf("Rule contains illegal character '%c' after '%c'.\n", line[i+1], line[i])
			return false
		}

		// for not
		if string(line[i]) == not && !charInString(rune(line[i+1]), openBra+factSymbol) {
			fmt.Printf("Rule contains illegal character '%c' after '%c'.\n", line[i+1], line[i])
			return false
		}

		// for operators and, or, xor
		operators := and + or + xor
		if charInString(rune(line[i]), operators) && !charInString(rune(line[i+1]), openBra+not+factSymbol) {
			fmt.Printf("Rule contains illegal character '%c' after '%c'.\n", line[i+1], line[i])
			return false
		}

		// for imp and ioi
		if strings.HasPrefix(line[i:], ioi) && !charInString(rune(line[i+3]), openBra+not+factSymbol) {
			fmt.Printf("Rule contains illegal character '%c' after '%c'.\n", line[i+3], line[i])
			return false
		} else if strings.HasPrefix(line[i:], imp) && !charInString(rune(line[i+2]), openBra+not+factSymbol) {
			fmt.Printf("Rule contains illegal character '%c' after '%c'.\n", line[i+2], line[i])
			return false
		}

		// for factSymbol
		if charInString(rune(line[i]), factSymbol) && !charInString(rune(line[i+1]), (openBra+closeBra+and+or+xor+imp+ioi)) {
			fmt.Printf("Rule contains illegal character '%c' after '%c'.\n", line[i+1], line[i])
			return false
		}
	}
	return true
}
