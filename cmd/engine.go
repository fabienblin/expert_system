/* ************************************************************************** */
/*                                                          LE - /            */
/*                                                              /             */
/*   engine.go                                        .::    .:/ .      .::   */
/*                                                 +:+:+   +:    +:  +:+:+    */
/*   By: jojomoon <jojomoon@student.le-101.fr>      +:+   +:    +:    +:+     */
/*                                                 #+#   #+    #+    #+#      */
/*   Created: 2019/10/30 17:51:53 by jmonneri     #+#   ##    ##    #+#       */
/*   Updated: 2020/01/07 15:32:07 by jojomoon    ###    #+. /#+    ###.fr     */
/*                                                         /                  */
/*                                                        /                   */
/* ************************************************************************** */
package main

import (
	"fmt"
	"strings"
)

var i int = 0 // !!!! a enlever c du debug

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
		fmt.Printf("# solution %s = %d\n", env.factList[query].op, env.factList[query].value)
	}
	return true, nil
}

func backwardChaining(query *fact, checked []string) error {
	i++
	fmt.Printf("%*sBackward Chaining:\n", i, " ")
	// On check que le fact n'ait pas déjà été demandé (anti-boucle).
	fmt.Printf("%*sBackward Chaining: query searched: %s\n", i, " ", query.op)
	if stringInSlice(query.op, checked) {
		fmt.Printf("%*sBackward Chaining: abort because of already checked\n", i, " ")
		i--
		return nil
	}
	checked = append(checked, query.op)
	// On trouve les règles définissant la query
	for _, rule := range env.trees {
		if node := digInRule(query, rule); node != nil {
			fmt.Printf("%*sBackward Chaining: rule found: \n", i, " ")
			printNode(rule, 4)
			err := resolve(node, node, checked) // !! TESTER QUE CELA A FONCTIONNE
			if err != nil {
				i--
				return err
			}
			printNode(rule, 4)
		}
	}
	i--
	return nil
}

func digInRule(fact *fact, node *infTree) *infTree {
	i++
	if strings.Contains(factSymbol, node.fact.op) {
		if node.fact == fact {
			i--
			return node
		}
		i--
		return nil
	}
	if node.fact.op != imp && node.left != nil {
		if node := digInRule(fact, node.left); node != nil {
			i--
			return node
		}
	}
	i--
	return digInRule(fact, node.right)
}

func resolve(node *infTree, from *infTree, checked []string) error {
	i++
	fmt.Printf("%*sResolve:\n", i, " ")
	var err error = nil

	if node == nil || node.fact.isKnown {
		fmt.Printf("%*sResolve: node is nil or known\n", i, " ")
		i--
		return nil
	}
	fmt.Printf("%*sResolve: node = %s = %d\n", i, " ", node.fact.op, node.fact.value)
	if from != node.head && !(node.fact.op == imp || node.fact.op == ioi) {
		err = resolve(node.head, node, checked)
		if node == from || err != nil {
			i--
			return err
		}
	}
	if strings.Contains(factSymbol, node.fact.op) {
		if !node.fact.isKnown {
			return backwardChaining(node.fact, checked)
		}
		i--
		return nil
	}
	if from == node.head {
		err := resolve(node.left, node, checked)
		if err == nil {
			err = resolve(node.right, node, checked)
		}
		if err != nil {
			i--
			return err
		}
	}
	// On lance la fonction de l'operateur
	i--
	return opeFunc[node.fact.op](node, from, checked)
}
