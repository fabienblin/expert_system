/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   engine.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/10/30 17:51:53 by jmonneri          #+#    #+#             */
/*   Updated: 2019/12/19 18:46:16 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"strings"
)

func engine() {
	fmt.Printf("Engine:\n")
	worked, err := searchQueries(env.queries)
	if !worked {
		outputError(err)
	} else {
		output()
	}
}

func searchQueries(queries []string) (bool, error) {
	for _, query := range queries {
		if err := backwardChaining(env.factList[query], []string{}); err != nil {
			return false, err
		}
	}
	return true, nil
}

func backwardChaining(query *fact, checked []string) error {
	// On check que le fact n'ait pas déjà été demandé (anti-boucle).
	if stringInSlice(query.op, checked) {
		return nil
	}
	checked = append(checked, query.op)
	// On trouve les règles définissant la query
	for _, rule := range env.trees {
		if node := digInRule(query, rule); node != nil {
			err := resolve(node, node, checked) // !! TESTER QUE CELA A FONCTIONNE
		}
	}
	return nil
}

func digInRule(fact *fact, node *infTree) *infTree {
	if strings.Contains(factSymbol, node.fact.op) {
		if node.fact.op == fact.op {
			return node
		}
		return nil
	}
	if node.fact.op != imp && node.left != nil {
		digInRule(fact, node.left)
	}
	return digInRule(fact, node.right)
}

func resolve(node *infTree, from *infTree, checked []string) error {
	var err error = nil

	if node == nil {
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
		err = resolve(node.left, node, checked)
		err = resolve(node.right, node, checked)
	}
	// On lance la fonction de l'operateur
	err = opeFunc[node.fact.op](node)
	if err != nil {
		return err
	}
	return err
}
