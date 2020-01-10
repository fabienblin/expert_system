/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   utils.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/10/30 17:52:29 by jmonneri          #+#    #+#             */
/*   Updated: 2019/10/30 17:53:01 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

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
