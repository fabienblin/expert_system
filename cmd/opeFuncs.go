/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   opeFuncs.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/11/11 14:34:50 by jmonneri          #+#    #+#             */
/*   Updated: 2019/12/19 12:56:49 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"errors"
	"fmt"
)

func setToTrueF(node *infTree) (bool, error) {
	if node.fact.value == falseF {
		return false, errors.New("Contradiction dans les données")
	} else if !node.fact.isKnown {
		node.fact.isKnown = true
		node.fact.value = trueF
		return true, nil
	}
	return false, nil
}

func setToFalseF(node *infTree) (bool, error) {
	if node.fact.value == trueF {
		return false, errors.New("Contradiction dans les données")
	} else if !node.fact.isKnown {
		node.fact.isKnown = true
		node.fact.value = falseF
		return true, nil
	}
	return false, nil
}

func setToUnknownF(node *infTree) (bool, error) {
	if node.fact.isKnown {
		return false, errors.New("Contradiction dans les données") //!! a enlever
	} else if node.fact.value != unknownF {
		node.fact.value = unknownF
		return true, nil
	}
	return false, nil
}

func notFunc(node *infTree, checked []string) (bool, error) {
	fmt.Printf("notFunc\n")
	var err error = nil
	if !node.right.fact.isKnown {
		if _, err = forwardChaining(node.right, checked); err != nil {
			return false, err
		}
	}
	// On sépare les 2 if car la ligne du dessus risque de changer node.right.fact.isKnown
	if node.right.fact.isKnown && err == nil {
		if node.right.fact.value == trueF {
			return setToFalseF(node)
		}
		return setToTrueF(node)
	}
	return false, err
}

func andFunc(node *infTree, checked []string) (bool, error) {
	fmt.Printf("andFunc\n")
	leftFact := node.left.fact
	rightFact := node.right.fact

	// Dans les cas ci-dessous on a besoin d' informations donc on relance le forward pour aller les chercher
	if !leftFact.isKnown && !rightFact.isKnown {
		if _, err := forwardChaining(node.right, checked); err != nil {
			return false, err
		}
		if _, err := forwardChaining(node.left, checked); err != nil {
			return false, err
		}
	} else if !leftFact.isKnown && rightFact.value == trueF {
		if _, err := forwardChaining(node.left, checked); err != nil {
			return false, err
		}
	} else if !rightFact.isKnown && leftFact.value == trueF {
		if _, err := forwardChaining(node.right, checked); err != nil {
			return false, err
		}
	}
	// Traitement des donnees recues
	if leftFact.value == falseF || rightFact.value == falseF {
		return setToFalseF(node)
	}
	if leftFact.value == trueF && rightFact.value == trueF {
		return setToTrueF(node)
	}
	return false, nil
}

func orFunc(node *infTree, checked []string) (bool, error) {
	fmt.Printf("orFunc\n")
	leftFact := node.left.fact
	rightFact := node.right.fact

	// Recherche d' informations manquantes
	if !leftFact.isKnown && rightFact.value != trueF {
		if _, err := forwardChaining(node.left, checked); err != nil {
			return false, err
		}
	}
	if !rightFact.isKnown && leftFact.value != trueF {
		if _, err := forwardChaining(node.right, checked); err != nil {
			return false, err
		}
	}
	// Traitement des informations
	if leftFact.value == trueF || rightFact.value == trueF {
		return setToTrueF(node)
	} else if leftFact.value == falseF && rightFact.value == falseF {
		return setToFalseF(node)
	}
	return false, nil
}

func xorFunc(node *infTree, checked []string) (bool, error) {
	fmt.Printf("xorFunc\n")
	leftFact := node.left.fact
	rightFact := node.right.fact

	// Seek
	if !leftFact.isKnown {
		if _, err := forwardChaining(node.left, checked); err != nil {
			return false, err
		}
	}
	if !rightFact.isKnown {
		if _, err := forwardChaining(node.right, checked); err != nil {
			return false, err
		}
	}
	// Treatment
	if leftFact.isKnown && (leftFact.value == rightFact.value) {
		return setToFalseF(node)
	} else if leftFact.isKnown && rightFact.isKnown && (leftFact.value != rightFact.value) {
		return setToTrueF(node)
	}
	return false, nil
}

func impFunc(node *infTree, checked []string) (bool, error) {
	fmt.Printf("impFunc\n")

	// SEEK
	if !node.left.fact.isKnown {
		if _, err := forwardChaining(node.left, checked); err != nil {
			return false, err
		}
	}
	// TREAT
	if node.left.fact.value == trueF {
		return setToTrueF(node.right)
	}
	return false, nil
}

func ioiFunc(node *infTree, checked []string) (bool, error) {
	fmt.Printf("ioiFunc\n")

	// SEEK
	if !node.left.fact.isKnown && !node.right.fact.isKnown {
		if _, err := forwardChaining(node.left, checked); err != nil {
			return false, err
		}
		if _, err := forwardChaining(node.right, checked); err != nil {
			return false, err
		}
	}
	// TREAT
	if node.left.fact.isKnown {
		if node.left.fact.value == trueF {
			return setToTrueF(node.right)
		}
		return setToFalseF(node.right)
	}
	if node.right.fact.isKnown {
		if node.right.fact.value == trueF {
			return setToTrueF(node.left)
		}
		return setToFalseF(node.left)
	}
	return false, nil
}
