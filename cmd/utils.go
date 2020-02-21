/* ************************************************************************** */
/*                                                          LE - /            */
/*                                                              /             */
/*   utils.go                                         .::    .:/ .      .::   */
/*                                                 +:+:+   +:    +:  +:+:+    */
/*   By: jojomoon <jojomoon@student.le-101.fr>      +:+   +:    +:    +:+     */
/*                                                 #+#   #+    #+    #+#      */
/*   Created: 2019/10/30 17:52:29 by jmonneri     #+#   ##    ##    #+#       */
/*   Updated: 2020/02/05 23:42:08 by jojomoon    ###    #+. /#+    ###.fr     */
/*                                                         /                  */
/*                                                        /                   */
/* ************************************************************************** */
package main

import (
	"fmt"
	"strings"
)

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
 * Print infTree with variable indentation
 */
func printNode(node *infTree, indent int, factCibled *infTree) {
	if node == nil {
		return
	}
	printNode(node.right, indent+4, factCibled)

	for i := 0; i < indent; i++ {
		fmt.Printf(" ")
	}
	if factCibled == node {
		fmt.Printf("\033[2m%v\033[0m\n", node.fact.op)
	} else if node.fact != nil {
		if node.fact.value == trueF {
			if node.fact.fixed {
				fmt.Printf("\033[32m")
			} else {
				fmt.Printf("\033[2;49;32m")
			}
		} else if node.fact.value == falseF {
			if node.fact.fixed {
				fmt.Printf("\033[31m")
			} else {
				fmt.Printf("\033[2;49;31m")
			}
		} else if node.fact.value == unknownF {
			fmt.Printf("\033[33m")
		}
		fmt.Printf("%v\033[0m\n", node.fact.op)
	}
	printNode(node.left, indent+4, factCibled)
}

func getNode(node *infTree, indent int, factCibled *infTree) string {
	if node == nil {
		return ""
	}
	var ret = getNode(node.right, indent + 4, factCibled)

	for i := 0; i < indent; i++ {
		ret += " "
	}
	if factCibled == node {
		ret += "\033[1;49;"
	} else if !node.fact.fixed && node.fact.value != unknownF{
		ret += "\033[2;49;"
	} else {
		ret += "\033[0;49;"
	}
	if node.fact.value == trueF {
		ret += "32m"
	} else if node.fact.value == falseF {
		ret += "31m"
	} else if node.fact.value == unknownF {
		ret += "33m"
	} else if node.fact.value == errorF {
		ret += "95m"
	} else {
		ret += "0m"
	}
	ret += node.fact.op + "\033[0m\n"
	return ret + getNode(node.left, indent+4, factCibled)
}

func nodeToStr(node *infTree) string {
	if node == nil {
		return ""
	} else if strings.Contains(factSymbol, node.fact.op) {
		return node.fact.op
	} else if node.fact.op == not {
		return "!" + nodeToStr(node.right)
	} else if node.fact.op == and {
		return "(" + nodeToStr(node.left) + " + " + nodeToStr(node.right) + ")"
	} else if node.fact.op == or {
		return "(" + nodeToStr(node.left) + " | " + nodeToStr(node.right) + ")"
	} else if node.fact.op == xor {
		return "(" + nodeToStr(node.left) + " ^ " + nodeToStr(node.right) + ")"
	} else if node.fact.op == imp {
		return nodeToStr(node.left) + " => " + nodeToStr(node.right)
	} else if node.fact.op == ioi {
		return nodeToStr(node.left) + " <=> " + nodeToStr(node.right)
	}
	return ""
}

func getFalse(node1 *infTree, node2 *infTree) *infTree {
	if node1 != nil && node1.fact.value == falseF {
		return node1
	} else if node2 != nil && node2.fact.value == falseF {
		return node2
	}
	return nil
}

func getTrue(node1 *infTree, node2 *infTree) *infTree {
	if node1 != nil && node1.fact.value == trueF {
		return node1
	} else if node2 != nil && node2.fact.value == trueF {
		return node2
	}
	return nil
}

func getContextRule(node *infTree) string {
	if node.fact.op == imp || node.fact.op == ioi {
		return "In the rule:\n" + getNode(node, 2, nil) + "\n"
	}
	return getContextRule(node.head)
}

func getContextRule2(node *infTree) string {
	if node.fact.op == imp || node.fact.op == ioi {
		return getNode(node, 2, nil)
	}
	return getContextRule2(node.head)
}