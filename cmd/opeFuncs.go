/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   opeFuncs.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/11/11 14:34:50 by jmonneri          #+#    #+#             */
/*   Updated: 2020/01/10 18:17:03 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"errors"
	"fmt"
)

func setToTrueF(node *infTree, toPrint string) error {
	if node.fact.value == falseF {
		return errors.New("Error: Contradiction in the facts")
	} else if !node.fact.isKnown {
		if verbose {
			fmt.Println(toPrint)
		}
		node.fact.isKnown = true
		node.fact.value = trueF
	}
	return nil
}

func setToFalseF(node *infTree, toPrint string) error {
	if node.fact.value == trueF {
		return errors.New("Error: Contradiction in the facts")
	} else if !node.fact.isKnown {
		if verbose {
			fmt.Println(toPrint)
		}
		node.fact.isKnown = true
		node.fact.value = falseF
	}
	return nil
}

func setToUnknownF(node *infTree, toPrint string) error {
	if !node.fact.isKnown && node.fact.value != unknownF {
		if verbose {
			fmt.Println(toPrint)
		}
		node.fact.value = unknownF
	}
	return nil
}

func getOtherSide(node *infTree, firstSide *infTree) *infTree {
	if firstSide == node.left {
		return node.right
	}
	return node.left
}

func notFunc(node *infTree, from *infTree, checked []string) error {
	if from == node.head {
		if node.right.fact.isKnown {
			if node.right.fact.value == trueF {
				return setToFalseF(node, getContextRule(node)+"We know that "+nodeToStr(node.right)+" is true so "+nodeToStr(node)+" is false")
			}
			return setToTrueF(node, getContextRule(node)+"We know that "+nodeToStr(node.right)+" is false so "+nodeToStr(node)+" is true")
		}
	} else {
		if node.fact.isKnown {
			if node.fact.value == trueF {
				return setToFalseF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is true so "+nodeToStr(node.right)+" is false")
			}
			return setToTrueF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is false so "+nodeToStr(node)+" is true")
		}
	}
	return nil
}

func andFunc(node *infTree, from *infTree, checked []string) error {
	leftFact := node.left.fact
	rightFact := node.right.fact

	if from == node.head {
		if leftFact.value == falseF || rightFact.value == falseF {
			return setToFalseF(node, getContextRule(node)+"We know that "+nodeToStr(getFalse(node.left, node.right))+" is false so "+nodeToStr(node)+" is false")
		}
		if leftFact.value == trueF && rightFact.value == trueF {
			return setToTrueF(node, getContextRule(node)+"We know that"+nodeToStr(node.left)+" and "+nodeToStr(node.right)+" are true so "+nodeToStr(node)+" is true")
		}
	} else {
		if node.fact.value == trueF {
			err := setToTrueF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is true so "+nodeToStr(node.right)+" and "+nodeToStr(node.left)+" are true")
			if err == nil {
				err = setToTrueF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is true so "+nodeToStr(node.right)+" and "+nodeToStr(node.left)+" are true")
			}
			return err
		}
		if node.fact.value == falseF {
			if rightFact.value == trueF {
				return setToFalseF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is false and "+nodeToStr(node.right)+" is true so "+nodeToStr(node.left)+" is false")
			} else if leftFact.value == trueF {
				return setToFalseF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is false and "+nodeToStr(node.left)+" is true so "+nodeToStr(node.right)+" is false")
			}
		}
	}
	return nil
}

func orFunc(node *infTree, from *infTree, checked []string) error {
	leftFact := node.left.fact
	rightFact := node.right.fact

	if from == node.head {
		if leftFact.value == trueF || rightFact.value == trueF {
			return setToTrueF(node, getContextRule(node)+"We know that "+nodeToStr(getTrue(node.left, node.right))+" is true so "+nodeToStr(node)+" is true")
		} else if leftFact.value == falseF && rightFact.value == falseF {
			return setToFalseF(node, getContextRule(node)+"We know that"+nodeToStr(node.left)+" and "+nodeToStr(node.right)+" are false so "+nodeToStr(node)+" is false")
		}
	} else {
		if node.fact.value == falseF {
			err := setToFalseF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is false so "+nodeToStr(node.right)+" and "+nodeToStr(node.left)+" are false")
			if err == nil {
				err = setToFalseF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is false so "+nodeToStr(node.right)+" and "+nodeToStr(node.left)+" are false")
			}
			return err
		} else if node.fact.value == trueF {
			if rightFact.value == falseF {
				return setToTrueF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is true and "+nodeToStr(node.right)+" is false so "+nodeToStr(node.left)+" is true")
			} else if leftFact.value == falseF {
				return setToTrueF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is true and "+nodeToStr(node.left)+" is false so "+nodeToStr(node.right)+" is true")
			}
			setToUnknownF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is true but we don't know anyting for childs so "+nodeToStr(node.left)+" and "+nodeToStr(node.right)+" are undetermined")
			setToUnknownF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is true but we don't know anyting for childs so "+nodeToStr(node.left)+" and "+nodeToStr(node.right)+" are undetermined")
		} else if node.fact.value == unknownF {
			setToUnknownF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is undetermined so "+nodeToStr(node.left)+" and "+nodeToStr(node.right)+" are undetermined")
			setToUnknownF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is undetermined so "+nodeToStr(node.left)+" and "+nodeToStr(node.right)+" are undetermined")
		}
	}
	return nil
}

func xorFunc(node *infTree, from *infTree, checked []string) error {
	leftFact := node.left.fact
	rightFact := node.right.fact

	if from == node.head {
		if leftFact.isKnown && (leftFact.value == rightFact.value) {
			return setToFalseF(node, getContextRule(node)+"We know that "+nodeToStr(node.left)+" are both of the same value so "+nodeToStr(node)+" is false")
		} else if leftFact.isKnown && rightFact.isKnown {
			return setToTrueF(node, getContextRule(node)+"We know that "+nodeToStr(getTrue(node.left, node.right))+" is true and "+nodeToStr(getFalse(node.left, node.right))+" is false so "+nodeToStr(node)+" is true")
		}
	} else {
		if node.fact.value == trueF {
			if leftFact.isKnown {
				if leftFact.value == falseF {
					return setToTrueF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is true and "+nodeToStr(node.left)+" is false so "+nodeToStr(node.right)+" is true")
				}
				return setToFalseF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is true and "+nodeToStr(node.left)+" is true so "+nodeToStr(node.right)+" is false")
			} else if rightFact.isKnown {
				if rightFact.value == falseF {
					return setToTrueF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is true and "+nodeToStr(node.right)+" is false so "+nodeToStr(node.left)+" is true")
				}
				return setToFalseF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is true and "+nodeToStr(node.left)+" is true so "+nodeToStr(node.right)+" is false")
			}
			setToUnknownF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is true but we don't know anyting for childs so "+nodeToStr(node.left)+" and "+nodeToStr(node.right)+" are undetermined")
			setToUnknownF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is true but we don't know anyting for childs so "+nodeToStr(node.left)+" and "+nodeToStr(node.right)+" are undetermined")
		} else if node.fact.value == falseF {
			if leftFact.isKnown {
				if leftFact.value == falseF {
					return setToFalseF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is false and "+nodeToStr(node.left)+" is false so "+nodeToStr(node.right)+" is true")
				}
				return setToTrueF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is false and "+nodeToStr(node.left)+" is true so "+nodeToStr(node.right)+" is true")
			} else if rightFact.isKnown {
				if rightFact.value == falseF {
					return setToFalseF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is false and "+nodeToStr(node.right)+" is false so "+nodeToStr(node.left)+" is true")
				}
				return setToTrueF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is false and "+nodeToStr(node.right)+" is true so "+nodeToStr(node.left)+" is true")
			}
			setToUnknownF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is false but we don't know anyting for childs so "+nodeToStr(node.left)+" and "+nodeToStr(node.right)+" are undetermined")
			setToUnknownF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is false but we don't know anyting for childs so "+nodeToStr(node.left)+" and "+nodeToStr(node.right)+" are undetermined")
		}
	}
	return nil
}

func impFunc(node *infTree, from *infTree, checked []string) error {
	if !node.left.fact.isKnown {
		if err := resolve(node.left, node, checked); err != nil {
			return err
		}
	}
	if node.left.fact.value == trueF {
		return setToTrueF(node.right, getContextRule(node)+"We know that "+nodeToStr(node.left)+" is true so "+nodeToStr(node.right)+" is true")
	}
	return nil
}

func ioiFunc(node *infTree, from *infTree, checked []string) error {
	var to *infTree = getOtherSide(node, from)

	if !to.fact.isKnown {
		if err := resolve(to, node, checked); err != nil {
			return err
		}
	}
	if to.fact.value == falseF {
		return setToFalseF(from, getContextRule(node)+"We know that "+nodeToStr(to)+" is false so "+nodeToStr(from)+" is false")
	} else if to.fact.value == trueF {
		return setToTrueF(from, getContextRule(node)+"We know that "+nodeToStr(to)+" is true so "+nodeToStr(from)+" is true")
	}
	return nil
}
