/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   engine.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/10/30 17:51:53 by jmonneri          #+#    #+#             */
/*   Updated: 2019/11/27 00:48:58 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"strings"
)

func engine() {
	fmt.Printf("Engine:\n")
	worked, err := computeTrees()
	if !worked {
		outputError(err)
	} else {
		output()
	}
}

func computeTrees() (bool, error) {
	changed := true
	for changed {
		changed = false
		for _, tree := range env.trees {
			changedThings, err := forwardChaining(tree)
			if err != nil {
				return false, err
			}
			changed = changedThings || changed
		}
	}
	return true, nil
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
	if changedNode, err = opeFunc[node.fact.op](node); err != nil {
		return false, err
	}

	return changedL || changedR || changedNode, nil
}
