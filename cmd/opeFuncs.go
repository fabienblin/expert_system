/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   opeFuncs.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/11/11 14:34:50 by jmonneri          #+#    #+#             */
/*   Updated: 2019/12/19 19:12:32 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"errors"
	"fmt"
)

func setToTrueF(node *infTree) error {
	if node.fact.value == falseF {
		return errors.New("Contradiction dans les données")
	} else if !node.fact.isKnown {
		node.fact.isKnown = true
		node.fact.value = trueF
	}
	return nil
}

func setToFalseF(node *infTree) error {
	if node.fact.value == trueF {
		return errors.New("Contradiction dans les données")
	} else if !node.fact.isKnown {
		node.fact.isKnown = true
		node.fact.value = falseF
	}
	return nil
}

func setToUnknownF(node *infTree) error {
	if !node.fact.isKnown && node.fact.value != unknownF {
		node.fact.value = unknownF
	}
	return nil
}

func notFunc(node *infTree, checked []string) error {
	fmt.Printf("notFunc\n")

	if node.fact.isKnown {
		if node.fact.value == trueF {
			return setToFalseF(node.right)
		}
		return setToTrueF(node.right)
	}
	if node.right.fact.isKnown {
		if node.right.fact.value == trueF {
			return setToFalseF(node)
		}
		return setToTrueF(node)
	}
	return nil
}

func andFunc(node *infTree, checked []string) error {
	fmt.Printf("andFunc\n")
	leftFact := node.left.fact
	rightFact := node.right.fact

	if node.fact.value == trueF {
		err := setToTrueF(node.left)
		if err == nil {
			err = setToTrueF(node.right)
		}
		return err
	}
	if node.fact.value == falseF {
		if rightFact.value == trueF {
			return setToFalseF(node.left)
		} else if leftFact.value == trueF {
			return setToFalseF(node.right)
		}
	}
	if leftFact.value == falseF || rightFact.value == falseF {
		return setToFalseF(node)
	}
	if leftFact.value == trueF && rightFact.value == trueF {
		return setToTrueF(node)
	}
	return nil
}

func orFunc(node *infTree, checked []string) error {
	fmt.Printf("orFunc\n")
	leftFact := node.left.fact
	rightFact := node.right.fact

	if leftFact.value == trueF || rightFact.value == trueF {
		return setToTrueF(node)
	} else if node.fact.value == falseF {
		err := setToFalseF(node.left)
		if err == nil {
			err = setToFalseF(node.right)
		}
		return err
	} else if node.fact.value == trueF {
		if rightFact.value == falseF {
			return setToTrueF(node.left)
		} else if leftFact.value == falseF {
			return setToTrueF(node.right)
		}
		setToUnknownF(node.left)
		setToUnknownF(node.right)
	}
	return nil
}

func xorFunc(node *infTree, checked []string) error {
	fmt.Printf("xorFunc\n")
	leftFact := node.left.fact
	rightFact := node.right.fact

	if node.fact.value == trueF {
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
	} else if node.fact.value == falseF {
		if leftFact.isKnown {
			if leftFact.value == falseF {
				return setToFalseF(node.right)
			}
			return setToTrueF(node.right)
		} else if rightFact.isKnown {
			if rightFact.value == falseF {
				return setToFalseF(node.left)
			}
			return setToTrueF(node.left)
		}
		setToUnknownF(node.left)
		setToUnknownF(node.right)
	}
	if leftFact.isKnown && (leftFact.value == rightFact.value) {
		return setToFalseF(node)
	} else if leftFact.isKnown && rightFact.isKnown {
		return setToTrueF(node)
	}
	return nil
}

func impFunc(node *infTree, checked []string) (bool, error) {
	fmt.Printf("impFunc\n")

	if !node.left.fact.isKnown {
		if err := resolve(node.left, node, checked); err != nil {
			return false, err
		}
	}
	if node.left.fact.value == trueF {
		return false, setToTrueF(node.right)
	}
	return false, nil
}

func ioiFunc(node *infTree, checked []string) (bool, error) {
	fmt.Printf("ioiFunc\n")

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
			return false, setToTrueF(node.right)
		}
		return false, setToFalseF(node.right)
	}
	if node.right.fact.isKnown {
		if node.right.fact.value == trueF {
			return false, setToTrueF(node.left)
		}
		return false, setToFalseF(node.left)
	}
	return false, nil
}
