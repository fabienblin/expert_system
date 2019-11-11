/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   engine.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/10/30 17:51:53 by jmonneri          #+#    #+#             */
/*   Updated: 2019/11/11 18:40:04 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"errors"
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
			changedThings, err := backwardChaining(tree)
			if err != nil {
				return false, err
			}
			changed = changedThings
		}
	}
	return true, nil
}

func backwardChaining(node *infTree) (bool, error) {
	// Ici on va pouvoir gérer quand on n'a pas de rules
	if node == nil {
		return false, errors.New("Qu'est-ce que je fous la? mon node est nil")
	}
	// Cela permet de ne pas descendre trop loin dans l'arbre si la branche fille est déjà trouvée
	if node.fact.isKnown {
		fmt.Printf("%3s %2d => Head\n", node.fact.op, node.fact.value)
		return false, nil
	}
	// Ici on return car l'on est sur un caractere
	if strings.Contains(factSymbol, node.fact.op) {
		fmt.Printf("%3s %2d => Head\n", node.fact.op, node.fact.value)
		return node.fact.value != defaultF, nil
	}
	// On lance a gauche puis a droite et on recupere les valeurs de retour
	fmt.Printf("%3s %2d => Left\n", node.fact.op, node.fact.value)
	changedL, _ := backwardChaining(node.left)
	fmt.Printf("%3s %2d => Right\n", node.fact.op, node.fact.value)
	changedR, _ := backwardChaining(node.right)
	fmt.Printf("%3s %2d => Function\n", node.fact.op, node.fact.value)
	// On lance la fonction de l'operateur
	opeFunc[node.fact.op](node)

	return changedL || changedR, errors.New("Bad rules") // !! ligne a refaire
}
