/* ************************************************************************** */
/*                                                          LE - /            */
/*                                                              /             */
/*   opeFuncs.go                                      .::    .:/ .      .::   */
/*                                                 +:+:+   +:    +:  +:+:+    */
/*   By: jojomoon <jojomoon@student.le-101.fr>      +:+   +:    +:    +:+     */
/*                                                 #+#   #+    #+    #+#      */
/*   Created: 2019/11/11 14:34:50 by jmonneri     #+#   ##    ##    #+#       */
/*   Updated: 2020/02/05 16:19:37 by jojomoon    ###    #+. /#+    ###.fr     */
/*                                                         /                  */
/*                                                        /                   */
/* ************************************************************************** */
package main

import (
	"errors"
	"fmt"
)

func setToTrueF(node *infTree, toPrint string, fixed bool) error {
	var err error = nil
	if node.fact.value == falseF && ((fixed && node.fact.fixed) || (!fixed && !node.fact.fixed)) {
		return errors.New("Error: Contradiction in the facts")
	} else if !node.fact.isKnown || (fixed && !node.fact.fixed) {
		if verbose {
			fmt.Println(toPrint)
		}
		node.fact.value = trueF
		node.fact.fixed = fixed
		if (node.fact.isKnown && node.fact.value != trueF) {
			err = consequensesRelaunch()
		}
		node.fact.isKnown = true
	}
	return err
}

func setToFalseF(node *infTree, toPrint string, fixed bool) error {
	var err error = nil
	if node.fact.value == trueF && ((fixed && node.fact.fixed) || (!fixed && !node.fact.fixed)) {
		return errors.New("Error: Contradiction in the facts")
	} else if !node.fact.isKnown  || (fixed && !node.fact.fixed) {
		if verbose {
			fmt.Println(toPrint)
		}
		node.fact.fixed = fixed
		node.fact.value = falseF
		if (node.fact.isKnown && node.fact.value != falseF) {
			err = consequensesRelaunch()
		}
		node.fact.isKnown = true
	}
	return err
}

func setToUnknownF(node *infTree, toPrint string, checked []string ) error {
	if !node.fact.isKnown && node.fact.value != unknownF {
		if ok, err := seekForOtherSide(node, checked); err != nil {
			return err
		} else if ok {
			return nil
		}
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
				return setToFalseF(node, getContextRule(node)+"We know that "+nodeToStr(node.right)+" is true so "+nodeToStr(node)+" is false", node.right.fact.fixed)
			}
			return setToTrueF(node, getContextRule(node)+"We know that "+nodeToStr(node.right)+" is false so "+nodeToStr(node)+" is true", node.right.fact.fixed)
		}
	} else {
		if node.fact.isKnown {
			if node.fact.value == trueF {
				return setToFalseF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is true so "+nodeToStr(node.right)+" is false", node.fact.fixed)
			}
			return setToTrueF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is false so "+nodeToStr(node)+" is true", node.fact.fixed)
		}
	}
	return nil
}

func andFunc(node *infTree, from *infTree, checked []string) error {
	leftFact := node.left.fact
	rightFact := node.right.fact

	if from == node.head {
		if leftFact.value == falseF || rightFact.value == falseF {
			return setToFalseF(node, getContextRule(node)+"We know that "+nodeToStr(getFalse(node.left, node.right))+" is false so "+nodeToStr(node)+" is false", getFalse(node.left, node.right).fact.fixed)
		}
		if leftFact.value == trueF && rightFact.value == trueF {
			return setToTrueF(node, getContextRule(node)+"We know that "+nodeToStr(node.left)+" and "+nodeToStr(node.right)+" are true so "+nodeToStr(node)+" is true", leftFact.fixed && rightFact.fixed)
		}
	} else {
		if node.fact.value == trueF {
			err := setToTrueF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is true so "+nodeToStr(node.right)+" and "+nodeToStr(node.left)+" are true", node.fact.fixed)
			if err == nil {
				err = setToTrueF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is true so "+nodeToStr(node.right)+" and "+nodeToStr(node.left)+" are true", node.fact.fixed)
			}
			return err
		}
		if node.fact.value == falseF {
			if rightFact.value == trueF {
				return setToFalseF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is false and "+nodeToStr(node.right)+" is true so "+nodeToStr(node.left)+" is false", node.fact.fixed && rightFact.fixed)
			} else if leftFact.value == trueF {
				return setToFalseF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is false and "+nodeToStr(node.left)+" is true so "+nodeToStr(node.right)+" is false", node.fact.fixed && leftFact.fixed)
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
			return setToTrueF(node, getContextRule(node)+"We know that "+nodeToStr(getTrue(node.left, node.right))+" is true so "+nodeToStr(node)+" is true", getTrue(node.left, node.right).fact.fixed)
		} else if leftFact.value == falseF && rightFact.value == falseF {
			return setToFalseF(node, getContextRule(node)+"We know that "+nodeToStr(node.left)+" and "+nodeToStr(node.right)+" are false so "+nodeToStr(node)+" is false", leftFact.fixed && rightFact.fixed)
		}
	} else {
		if node.fact.value == falseF {
			err := setToFalseF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is false so "+nodeToStr(node.right)+" and "+nodeToStr(node.left)+" are false", node.fact.fixed)
			if err == nil {
				err = setToFalseF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is false so "+nodeToStr(node.right)+" and "+nodeToStr(node.left)+" are false", node.fact.fixed)
			}
			return err
		} else if node.fact.value == trueF {
			if rightFact.value == falseF {
				return setToTrueF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is true and "+nodeToStr(node.right)+" is false so "+nodeToStr(node.left)+" is true", node.fact.fixed && rightFact.fixed)
			} else if leftFact.value == falseF {
				return setToTrueF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is true and "+nodeToStr(node.left)+" is false so "+nodeToStr(node.right)+" is true", node.fact.fixed && leftFact.fixed)
			} else if leftFact.value == trueF || rightFact.value == trueF {
				return nil
			}
			setToUnknownF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is true but we don't know anyting for childs so "+nodeToStr(node.left)+" and "+nodeToStr(node.right)+" are undetermined", checked)
			setToUnknownF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is true but we don't know anyting for childs so "+nodeToStr(node.left)+" and "+nodeToStr(node.right)+" are undetermined", checked)
		} else if node.fact.value == unknownF {
			setToUnknownF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is undetermined so "+nodeToStr(node.left)+" and "+nodeToStr(node.right)+" are undetermined", checked)
			setToUnknownF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is undetermined so "+nodeToStr(node.left)+" and "+nodeToStr(node.right)+" are undetermined", checked)
		}
	}
	return nil
}

func xorFunc(node *infTree, from *infTree, checked []string) error {
	leftFact := node.left.fact
	rightFact := node.right.fact

	if from == node.head {
		if leftFact.isKnown && (leftFact.value == rightFact.value) {
			return setToFalseF(node, getContextRule(node)+"We know that "+nodeToStr(node.left)+" are both of the same value so "+nodeToStr(node)+" is false", leftFact.fixed && rightFact.fixed)
		} else if leftFact.isKnown && rightFact.isKnown {
			return setToTrueF(node, getContextRule(node)+"We know that "+nodeToStr(getTrue(node.left, node.right))+" is true and "+nodeToStr(getFalse(node.left, node.right))+" is false so "+nodeToStr(node)+" is true", leftFact.fixed && rightFact.fixed)
		}
	} else {
		if node.fact.value == trueF {
			if leftFact.isKnown {
				if leftFact.value == falseF {
					return setToTrueF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is true and "+nodeToStr(node.left)+" is false so "+nodeToStr(node.right)+" is true", node.fact.fixed && leftFact.fixed)
				}
				return setToFalseF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is true and "+nodeToStr(node.left)+" is true so "+nodeToStr(node.right)+" is false", node.fact.fixed && leftFact.fixed)
			} else if rightFact.isKnown {
				if rightFact.value == falseF {
					return setToTrueF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is true and "+nodeToStr(node.right)+" is false so "+nodeToStr(node.left)+" is true", node.fact.fixed && rightFact.fixed)
				}
				return setToFalseF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is true and "+nodeToStr(node.left)+" is true so "+nodeToStr(node.right)+" is false", node.fact.fixed && rightFact.fixed)
			}
			setToUnknownF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is true but we don't know anyting for childs so "+nodeToStr(node.left)+" and "+nodeToStr(node.right)+" are undetermined", checked)
			setToUnknownF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is true but we don't know anyting for childs so "+nodeToStr(node.left)+" and "+nodeToStr(node.right)+" are undetermined", checked)
		} else if node.fact.value == falseF {
			if leftFact.isKnown {
				if leftFact.value == falseF {
					return setToFalseF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is false and "+nodeToStr(node.left)+" is false so "+nodeToStr(node.right)+" is true", node.fact.fixed && leftFact.fixed)
				}
				return setToTrueF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is false and "+nodeToStr(node.left)+" is true so "+nodeToStr(node.right)+" is true", node.fact.fixed && leftFact.fixed)
			} else if rightFact.isKnown {
				if rightFact.value == falseF {
					return setToFalseF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is false and "+nodeToStr(node.right)+" is false so "+nodeToStr(node.left)+" is true", node.fact.fixed && rightFact.fixed)
				}
				return setToTrueF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is false and "+nodeToStr(node.right)+" is true so "+nodeToStr(node.left)+" is true", node.fact.fixed && rightFact.fixed)
			}
			setToUnknownF(node.left, getContextRule(node)+"We know that "+nodeToStr(node)+" is false but we don't know anyting for childs so "+nodeToStr(node.left)+" and "+nodeToStr(node.right)+" are undetermined", checked)
			setToUnknownF(node.right, getContextRule(node)+"We know that "+nodeToStr(node)+" is false but we don't know anyting for childs so "+nodeToStr(node.left)+" and "+nodeToStr(node.right)+" are undetermined", checked)
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
		return setToTrueF(node.right, getContextRule(node)+"We know that "+nodeToStr(node.left)+" is true so "+nodeToStr(node.right)+" is true", node.left.fact.fixed)
	}
	return nil
}

func ioiFunc(node *infTree, from *infTree, checked []string) error {
	var to *infTree = getOtherSide(node, from)
	leftFact := node.left.fact
	rightFact := node.right.fact

	if leftFact.fixed {
		if leftFact.isKnown {
			if leftFact.value == trueF {
				return setToTrueF(node.right, "We know that "+nodeToStr(node.left)+" is true so "+nodeToStr(node.right)+" is true", true)
			}
			return setToFalseF(node.right, "We know that "+nodeToStr(node.left)+" is false so "+nodeToStr(node.right)+" is false", true)
		}
	} else if rightFact.fixed {
		if rightFact.isKnown {
			if rightFact.value == trueF {
				return setToTrueF(node.left, "We know that "+nodeToStr(node.right)+" is true so "+nodeToStr(node.left)+" is true", true)
			}
			return setToFalseF(node.left, "We know that "+nodeToStr(node.right)+" is false so "+nodeToStr(node.left)+" is false", true)
		}
	}
	if !to.fact.isKnown {
		if err := resolve(to, node, checked); err != nil {
			return err
		}
	}
	if to.fact.value == falseF {
		return setToFalseF(from, getContextRule(node)+"We know that "+nodeToStr(to)+" is false so "+nodeToStr(from)+" is false", to.fact.fixed)
	} else if to.fact.value == trueF {
		return setToTrueF(from, getContextRule(node)+"We know that "+nodeToStr(to)+" is true so "+nodeToStr(from)+" is true", to.fact.fixed)
	}
	return nil
}