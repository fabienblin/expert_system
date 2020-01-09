/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   utils.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/10/30 17:52:29 by jmonneri          #+#    #+#             */
/*   Updated: 2020/01/09 19:53:16 by jmonneri         ###   ########.fr       */
/*                                                                            */
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
func printNode(node *infTree, indent int, factCibled *infTree) string {
	if node == nil {
		return ""
	}
	printNode(node.right, indent+4, factCibled)

	for i := 0; i < indent; i++ {
		fmt.Printf(" ")
	}
	if factCibled == node {
		fmt.Printf("\033[2m%v\033[0m\n", node.fact.op)
	} else {
		fmt.Printf("%v\n", node.fact.op)
	}
	printNode(node.left, indent+4, factCibled)
	return ""
}

func nodeToStr(node *infTree) string {
	if node == nil {
		return ""
	} else if strings.Contains(node.fact.op, factSymbol) {
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
