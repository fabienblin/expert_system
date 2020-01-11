/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   engine.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/10/30 17:51:53 by jmonneri          #+#    #+#             */
/*   Updated: 2020/01/11 02:17:18 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"strings"
)

/*
 * Run the inference engine
 */
func engine(flagForward bool) {
	if flagForward {
		computeTrees()
		for _, query := range env.queries {
			if env.factList[query].value == defaultF {
				env.factList[query].value = falseF
			}
			fmt.Printf("# solution %s = %d\n", env.factList[query].op, env.factList[query].value)
		}
	} else {
		for _, query := range env.queries {
			if err := backwardChaining(env.factList[query], []string{}); err != nil {
				fmt.Println(err)
			}
			fmt.Printf("# solution %s = %d\n", env.factList[query].op, env.factList[query].value)
		}
	}
}

func backwardChaining(query *fact, checked []string) error {
	// On check que le fact n'ait pas déjà été demandé (anti-boucle).
	if stringInSlice(query.op, checked) {
		return nil
	}
	checked = append(checked, query.op)
	// On trouve les règles définissant la query
	if verbose {
		fmt.Printf("Searching for queries defining %s\n", query.op)
	}
	for _, rule := range env.trees {
		if node := digInRule(query, rule); node != nil {
			if verbose {
				fmt.Printf("Rule found:\n")
				printNode(rule, 4, nil)
			}
			err := resolve(node, node, checked)
			if err != nil {
				return err
			}
		}
	}
	if query.value == defaultF {
		query.value = falseF
		query.isKnown = true
	}
	return nil
}

func digInRule(fact *fact, node *infTree) *infTree {
	if strings.Contains(factSymbol, node.fact.op) {
		if node.fact == fact {
			return node
		}
		return nil
	}
	if node.fact.op != imp && node.left != nil {
		if node := digInRule(fact, node.left); node != nil {
			return node
		}
	}
	return digInRule(fact, node.right)
}

func resolve(node *infTree, from *infTree, checked []string) error {
	var err error = nil
	if node == nil || node.fact.isKnown {
		return nil
	}
	if from != node.head && !(node.fact.op == imp || node.fact.op == ioi) {
		err = resolve(node.head, node, checked)
		if node == from || err != nil {
			return err
		}
	}
	if strings.Contains(factSymbol, node.fact.op) {
		if !node.fact.isKnown {
			return backwardChaining(node.fact, checked)
		}
		return nil
	}
	if from == node.head {
		err := resolve(node.left, node, checked)
		if err == nil {
			err = resolve(node.right, node, checked)
		}
		if err != nil {
			return err
		}
	}
	// On lance la fonction de l'operateur
	return opeFunc[node.fact.op](node, from, checked)
}
