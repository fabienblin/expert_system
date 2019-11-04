/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   engine.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/10/30 17:51:53 by jmonneri          #+#    #+#             */
/*   Updated: 2019/10/31 05:59:27 by jmonneri         ###   ########.fr       */
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
			changedThings, _, err := backwardChaining(&tree)
			if err != nil {
				return false, err
			} else if changedThings {
				changed = true
			}
		}
	}
	return true, nil
}

func backwardChaining(tree *infTree) (bool, int, error) {
	if tree == nil {
		return false, 0, errors.New("Qu'est-ce que je fous la? mon tree est nil")
	}
	// Cela permet de ne pas descendre trop loin dans l'arbre si la branche fille est déjà trouvée
	if tree.isTrue >= 0 {
		return true, tree.isTrue, nil
	}
	// Ici on return car l'on est sur un caractere
	if strings.Contains(factSymbol, tree.operator) {
		return false, tree.isTrue, nil
	}
	// On lance a gauche puis a droite et on recupere les valeurs de retour
	valueIsKnownL, valueL, _ := backwardChaining(tree.left)
	valueIsKnownR, valueR, _ := backwardChaining(tree.right)
	// On lance la fonction de l'operateur        !!! Il faut faire le tableau de pointeur sur fonction
	opeFunc[tree.operator](valueIsKnownL, valueL, valueIsKnownR, valueR)

	return true, 0, errors.New("Bad rules")
}
