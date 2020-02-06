/* ************************************************************************** */
/*                                                          LE - /            */
/*                                                              /             */
/*   engine.go                                        .::    .:/ .      .::   */
/*                                                 +:+:+   +:    +:  +:+:+    */
/*   By: jojomoon <jojomoon@student.le-101.fr>      +:+   +:    +:    +:+     */
/*                                                 #+#   #+    #+    #+#      */
/*   Created: 2019/10/30 17:51:53 by jmonneri     #+#   ##    ##    #+#       */
/*   Updated: 2020/02/06 01:17:05 by jojomoon    ###    #+. /#+    ###.fr     */
/*                                                         /                  */
/*                                                        /                   */
/* ************************************************************************** */
package main
// Stocker dans le node (pas dans le fact) un "checked" que l'on vérifie au moment où l'on passe dessus au lieu de vérifier le checked déjà présent. De plus, en écrasant, il faut faire péter une erreur si on est dans un imp et continuer d'écraser si on est dans un ioi
import (
	"errors"
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
				fmt.Printf("# solution %s = %s\n", env.factList[query].op, err)
			} else {
				fmt.Printf("# solution %s = %s\n", env.factList[query].op, output[env.factList[query].value])
			}
		}
	}
	if verbose {
		fmt.Printf(getNode(env.tree, 4, nil))
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
		fmt.Printf("Searching for rules defining %s\n", query.op)
	}
	for _, rule := range env.trees {
		if err := digInRule(query, rule, checked); err != nil {
			return err
		}
	}
	if query.value == defaultF {
		if verbose {
			fmt.Printf("Set %s to false by default: not enought information\n", query.op)
		}
		query.value = falseF
		query.isKnown = true
		query.fixed = false
	}
	if verbose {
		fmt.Printf("Stop searching %s\n", query.op)
	}
	return nil
}

func digInRule(fact *fact, node *infTree, checked []string) error {
	if strings.Contains(factSymbol, node.fact.op) {
		if node.fact == fact {
			if verbose {
				fmt.Printf("Rule found for %s searched:\n%s", fact.op, getContextRule2(node))
			}
			if err := resolve(node, node, checked); err != nil {
				node.fact.value = errorF
				node.fact.isKnown = false
				return err
			}
		}
		return nil
	}
	if node.fact.op != imp && node.left != nil {
		if err := digInRule(fact, node.left, checked); err != nil {
			return err
		}
	}
	return digInRule(fact, node.right, checked)
}

func resolve(node *infTree, from *infTree, checked []string) error {
	var err error = nil
	if node == nil{
		return nil
	}
	if node.fact.value == errorF {
		return errors.New("Error : There was already a contradiction for the fact " + node.fact.op)
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

func consequensesRelaunch() error {
	for _, fact := range env.factList {
		if fact.isKnown && !fact.fixed {
			if err := backwardChaining(fact, []string{}); err != nil {
				return err
			}
		}
	}
	return nil
}

func seekForOtherSide(node *infTree, checked []string) (bool, error) {
	var otherSide *infTree = getOtherSide(node.head, node)
	if err := resolve(otherSide, node.head, checked); err != nil {
		return false, err
	}
 	if otherSide.fact.isKnown {
		return true, opeFunc[node.head.fact.op](node.head, node, checked)
	}
	return false, nil
}