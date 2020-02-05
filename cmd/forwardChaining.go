/* ************************************************************************** */
/*                                                          LE - /            */
/*                                                              /             */
/*   forwardChaining.go                               .::    .:/ .      .::   */
/*                                                 +:+:+   +:    +:  +:+:+    */
/*   By: jojomoon <jojomoon@student.le-101.fr>      +:+   +:    +:    +:+     */
/*                                                 #+#   #+    #+    #+#      */
/*   Created: 2020/01/11 01:57:46 by jmonneri     #+#   ##    ##    #+#       */
/*   Updated: 2020/01/22 11:25:02 by jojomoon    ###    #+. /#+    ###.fr     */
/*                                                         /                  */
/*                                                        /                   */
/* ************************************************************************** */
package main

import (
	"fmt"
	"log"
	"strings"
)

func computeTrees() {
	changed := true
	for changed {
		changed = false
		for _, tree := range env.trees {
			changedThings, err := forwardChaining(tree)
			if err != nil {
				log.Fatal(err)
			}
			changed = changedThings || changed
		}
	}
	retry := false
	for _, fact := range env.factList {
		if !stringInSlice(fact.op, env.queries) && fact.value == defaultF {
			fact.value = falseF
			retry = true
		}
	}
	if retry {
		computeTrees()
	}
	return
}

func forwardChaining(node *infTree) (bool, error) {
	// Ici on va pouvoir gérer quand on n'a pas de rules
	if node == nil {
		return false, nil
	}
	// Ici on return car l'on est sur un caractere
	if strings.Contains(factSymbol, node.fact.op) {
		fmt.Printf("%3s %2d => Head\n", node.fact.op, node.fact.value)
		return false, nil
	}
	var err error
	var changedL, changedR, changedNode bool
	// On lance a gauche puis a droite et on recupere les valeurs de retour
	fmt.Printf("%3s %2d => Left\n", node.fact.op, node.fact.value)
	if changedL, err = forwardChaining(node.left); err != nil {
		return false, err
	}
	fmt.Printf("%3s %2d => Right\n", node.fact.op, node.fact.value)
	if changedR, err = forwardChaining(node.right); err != nil {
		return false, err
	}
	fmt.Printf("%3s %2d => Function\n", node.fact.op, node.fact.value)
	// On lance la fonction de l'operateur
	if changedNode, err = opeFuncFor[node.fact.op](node); err != nil {
		return false, err
	}

	return changedL || changedR || changedNode, nil
}
