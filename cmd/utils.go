package main

import "fmt"

/*
 * Find string in a string list
 */
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

/*
 * Find character in a string
 */
func charInString(c rune, str string) bool {
	for _, current := range str {
		if current == c {
			return true
		}
	}
	return false
}

/*
 * Print infTree with variable indetation
 */
func printNode(node *infTree, indent int) {
	if node == nil {
		return
	}
	printNode(node.right, indent+4)

	for i := 0; i < indent; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("%v [%v|%v]\n", node.fact.op, node.fact.isTrue, node.fact.isKnown)
	printNode(node.left, indent+4)
}
