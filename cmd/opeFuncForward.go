/* ************************************************************************** */
/*                                                          LE - /            */
/*                                                              /             */
/*   opeFuncForward.go                                .::    .:/ .      .::   */
/*                                                 +:+:+   +:    +:  +:+:+    */
/*   By: jojomoon <jojomoon@student.le-101.fr>      +:+   +:    +:    +:+     */
/*                                                 #+#   #+    #+    #+#      */
/*   Created: 2020/01/11 01:58:59 by jmonneri     #+#   ##    ##    #+#       */
/*   Updated: 2020/01/22 11:25:02 by jojomoon    ###    #+. /#+    ###.fr     */
/*                                                         /                  */
/*                                                        /                   */
/* ************************************************************************** */
package main

import (
	"errors"
	"fmt"
)

func setToTrueFFor(node *infTree) (bool, error) {
	if node.fact.value == falseF {
		return false, errors.New("Contradiction dans les données")
	} else if !node.fact.isKnown {
		node.fact.isKnown = true
		node.fact.value = trueF
		return true, nil
	}
	return false, nil
}

func setToFalseFFor(node *infTree) (bool, error) {
	if node.fact.value == trueF {
		return false, errors.New("Contradiction dans les données")
	} else if !node.fact.isKnown {
		node.fact.isKnown = true
		node.fact.value = falseF
		return true, nil
	}
	return false, nil
}

func setToUnknownFFor(node *infTree) (bool, error) {
	if node.fact.isKnown {
		return false, errors.New("Contradiction dans les données") //!! a enlever
	} else if node.fact.value != unknownF {
		node.fact.value = unknownF
		return true, nil
	}
	return false, nil
}

func notFuncFor(node *infTree) (bool, error) {
	fmt.Printf("notFunc\n")
	if node.right.fact.isKnown {
		if node.right.fact.value == trueF {
			return setToFalseFFor(node)
		}
		return setToTrueFFor(node)
	} else if node.fact.isKnown {
		if node.fact.value == trueF {
			return setToFalseFFor(node.right)
		}
		return setToTrueFFor(node.right)
	}
	return false, nil
}

func andFuncFor(node *infTree) (bool, error) {
	fmt.Printf("andFunc\n")
	leftFact := node.left.fact
	rightFact := node.right.fact
	if leftFact.value == trueF && rightFact.value == trueF {
		return setToTrueFFor(node)
	} else if node.fact.value == trueF {
		changedR := false
		changedL, err := setToTrueFFor(node.left)
		if err == nil {
			changedR, err = setToTrueFFor(node.right)
		}
		return changedR || changedL, err
	} else if leftFact.value == falseF || rightFact.value == falseF {
		return setToFalseFFor(node)
	} else if node.fact.value == falseF {
		if rightFact.value == trueF {
			return setToFalseFFor(node.left)
		} else if leftFact.value == trueF {
			return setToFalseFFor(node.right)
		}
	}
	return false, nil
}

func orFuncFor(node *infTree) (bool, error) {
	fmt.Printf("orFunc\n")
	leftFact := node.left.fact
	rightFact := node.right.fact
	if leftFact.value == trueF || rightFact.value == trueF {
		return setToTrueFFor(node)
	} else if node.fact.value == falseF {
		changedR := false
		changedL, err := setToFalseFFor(node.left)
		if err == nil {
			changedR, err = setToFalseFFor(node.right)
		}
		return changedR || changedL, err
	} else if node.fact.value == trueF {
		if rightFact.value == falseF {
			return setToTrueFFor(node.left)
		} else if leftFact.value == falseF {
			return setToTrueFFor(node.right)
		}
		if !leftFact.isKnown {
			setToUnknownFFor(node.left)
		}
		if !rightFact.isKnown {
			setToUnknownFFor(node.right)
		}
	}
	return false, nil
}

func xorFuncFor(node *infTree) (bool, error) {
	fmt.Printf("xorFunc\n")
	leftFact := node.left.fact
	rightFact := node.right.fact
	if leftFact.isKnown && (leftFact.value == rightFact.value) {
		return setToFalseFFor(node)
	} else if node.fact.value == trueF {
		if leftFact.isKnown {
			if leftFact.value == falseF {
				return setToTrueFFor(node.right)
			}
			return setToFalseFFor(node.right)
		} else if rightFact.isKnown {
			if rightFact.value == falseF {
				return setToTrueFFor(node.left)
			}
			return setToFalseFFor(node.left)
		}
		setToUnknownFFor(node.left)
		setToUnknownFFor(node.right)
	} else if leftFact.isKnown && rightFact.isKnown && (leftFact.value != rightFact.value) {
		return setToTrueFFor(node)
	} else if node.fact.value == falseF {
		if leftFact.isKnown {
			if leftFact.value == trueF {
				return setToTrueFFor(node.left)
			}
			return setToFalseFFor(node.left)
		} else if rightFact.isKnown {
			if rightFact.value == trueF {
				return setToTrueFFor(node.right)
			}
			return setToFalseFFor(node.right)
		}
	}
	return false, nil
}

func impFuncFor(node *infTree) (bool, error) {
	fmt.Printf("impFunc\n")
	if node.left.fact.value == trueF {
		return setToTrueFFor(node.right)
	}
	return false, nil
}

func ioiFuncFor(node *infTree) (bool, error) {
	fmt.Printf("ioiFunc\n")
	if node.left.fact.isKnown {
		if node.left.fact.value == trueF {
			return setToTrueFFor(node.right)
		}
		return setToFalseFFor(node.right)
	}
	if node.right.fact.isKnown {
		if node.right.fact.value == trueF {
			return setToTrueFFor(node.left)
		}
		return setToFalseFFor(node.left)
	}
	return false, nil
}
