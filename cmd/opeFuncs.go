/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   opeFuncs.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/11/11 14:34:50 by jmonneri          #+#    #+#             */
/*   Updated: 2019/11/11 21:07:21 by jmonneri         ###   ########.fr       */
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
		return false, errors.New("Contradiction dans les données")
	} else if node.fact.value != unknownF {
		node.fact.value = unknownF
		return true, nil
	}
	return false, nil
}

func notFunc(node *infTree) (bool, error) {
	fmt.Printf("notFunc\n")
	if node.right.fact.isKnown {
		if node.right.fact.value == trueF {
			return setToFalseF(node)
		}
		return setToTrueF(node)
	} else if node.fact.isKnown {
		if node.fact.value == trueF {
			return setToFalseF(node.right)
		}
		return setToTrueF(node.right)
	}
	return false, nil
}

func andFunc(node *infTree) (bool, error) {
	fmt.Printf("andFunc\n")
	leftFact := node.left.fact
	rightFact := node.right.fact
	if leftFact.value == trueF && rightFact.value == trueF {
		return setToTrueF(node)
	} else if node.fact.value == trueF {
		changedR := false
		changedL, err := setToTrueF(node.left)
		if err == nil {
			changedR, err = setToTrueF(node.right)
		}
		return changedR || changedL, err
	} else if leftFact.value == falseF || rightFact.value == falseF {
		return setToFalseF(node)
	} else if node.fact.value == falseF {
		if rightFact.value == trueF {
			return setToFalseF(node.left)
		} else if leftFact.value == trueF {
			return setToFalseF(node.right)
		}
	}
	return false, nil
}

func orFunc(node *infTree) (bool, error) {
	fmt.Printf("orFunc\n")
	leftFact := node.left.fact
	rightFact := node.right.fact
	if leftFact.value == trueF || rightFact.value == trueF {
		return setToTrueF(node)
	} else if node.fact.value == falseF {
		changedR := false
		changedL, err := setToFalseF(node.left)
		if err == nil {
			changedR, err = setToFalseF(node.right)
		}
		return changedR || changedL, err
	} else if node.fact.value == trueF {
		if rightFact.value == falseF {
			return setToTrueF(node.left)
		} else if leftFact.value == falseF {
			return setToTrueF(node.right)
		}
		if !leftFact.isKnown {
			setToUnknownF(node.left)
		}
		if !rightFact.isKnown {
			setToUnknownF(node.right)
		}
	}
	return false, nil
}

func xorFunc(node *infTree) (bool, error) { // J' en suis ici
	fmt.Printf("xorFunc\n")
	leftFact := node.left.fact
	rightFact := node.right.fact
	if leftFact.isKnown && (leftFact.value == rightFact.value) {
		return setToFalseF(node)
	} else if node.fact.value == trueF {
		if leftFact.isKnown {
			if leftFact.value == falseF {
				return setToTrueF(node.right)
			}
			return setToFalseF(node.right)
		} else if rightFact.isKnown {
			if rightFact.value == falseF {
				return setToTrueF(node.left)
			}
			return setToFalseF(node.left)
		}
		setToUnknownF(node.left)
		setToUnknownF(node.right)
	} else if leftFact.value == trueF || rightFact.value == trueF {
		return setToTrueF(node)
	} else if node.fact.value == falseF {

	}
	return false, nil
}

func ioiFunc(node *infTree) (bool, error) {
	fmt.Printf("ioiFunc\n")
	printNode(node, 4)
	return false, nil
}

func impFunc(node *infTree) (bool, error) {
	fmt.Printf("impFunc\n")
	printNode(node, 4)
	return false, nil
}
