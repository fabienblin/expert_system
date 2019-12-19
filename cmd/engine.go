/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   engine.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/10/30 17:51:53 by jmonneri          #+#    #+#             */
/*   Updated: 2019/12/19 12:58:15 by jmonneri         ###   ########.fr       */
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
		if _, err := backwardChaining(env.factList[query], []string{}); err != nil {
			return false, err
		}
	}
	return true, nil
}

func backwardChaining(query *fact, checked []string) (bool, error) {
	// On check que le fact n'ait pas déjà été demandé (anti-boucle).
	if stringInSlice(query.op, checked) {
		return false, nil
	}
	checked = append(checked, query.op)
	// On trouve les règles définissant la query
	for _, rule := range env.trees {
		if (query, rule) {
			// On lance la résolution
			if worked, error := forwardChaining(rule, checked); error != nil {
				return false, error
			}
		}
	}
	return true, nil
}

func isDefinedBy(fact *fact, node *infTree) bool {
	defined := false
	if strings.Contains(factSymbol, node.fact.op) {
		if node.fact.op == fact.op {
			return true
		}
		return false
	}
	if node.fact.op != imp && node.left != nil {
		defined = defined || isDefinedBy(fact, node.left)
	}
	return defined || isDefinedBy(fact, node.right)
}

/* func findUnknownFact(node *infTree, checked []string) (bool, error) {
	if node == nil {
		return false, nil
	}
	if strings.Contains(factSymbol, node.fact.op) {
		if node.fact.isKnown {
			return false, nil
		}
		return backwardChaining(node.fact, checked)
	}
	var err error
	var changedL, changedR, changedNode bool
	if changedL, err = findUnknownFact(node.left, checked); err != nil {
		return false, err
	}
	if node.fact.op == imp || node.fact.op == ioi {
		if worked, err := opeFunc[node.fact.op](node); err != nil {
			return false, err
		}
	}
	if changedR, err = findUnknownFact(node.right, checked); err != nil {
		return false, err
	}

	if changedNode, err = opeFunc[node.fact.op](node); err != nil {
		return false, err
	}

	return changedL || changedR || changedNode, nil
}

func computeTree(trees []*infTree) (bool, error) {
	changed := true
	for changed {
		changed = false
		for _, tree := range trees {
			changedThings, err := forwardChaining(tree)
			if err != nil {
				return false, err
			}
			changed = changedThings || changed
		}
	}
	return true, nil
}
*/
func forwardChaining(node *infTree, checked []string) (bool, error) {
	// Ici on va pouvoir gérer quand on n'a pas de rules
	if node == nil {
		return false, nil
	}
	// Ici on return car l'on est sur un caractere
	if strings.Contains(factSymbol, node.fact.op) {
		if !node.fact.isKnown {
			return backwardChaining(node.fact, checked)
		}
		return false, nil
	}
	fmt.Printf("%3s %2d => Function\n", node.fact.op, node.fact.value)
	// On lance la fonction de l'operateur
	changed, err := opeFunc[node.fact.op](node)
	if err != nil {
		return false, err
	}
	return changed, err
}
