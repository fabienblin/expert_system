/* ************************************************************************** */
/*                                                          LE - /            */
/*                                                              /             */
/*   opeFuncs.go                                      .::    .:/ .      .::   */
/*                                                 +:+:+   +:    +:  +:+:+    */
/*   By: jojomoon <jojomoon@student.le-101.fr>      +:+   +:    +:    +:+     */
/*                                                 #+#   #+    #+    #+#      */
/*   Created: 2019/11/11 14:34:50 by jmonneri     #+#   ##    ##    #+#       */
/*   Updated: 2020/01/07 15:51:41 by jojomoon    ###    #+. /#+    ###.fr     */
/*                                                         /                  */
/*                                                        /                   */
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

func getOtherSide(node *infTree, firstSide *infTree) *infTree {
	if firstSide == node.left {
		return node.right
	}
	return node.left
}

func notFunc(node *infTree, from *infTree, checked []string) error {
	i++
	fmt.Printf("%*sNotFunc\n", i, " ")

	if from == node.head {
		if node.right.fact.isKnown {
			if node.right.fact.value == trueF {
				i--
				return setToFalseF(node)
			}
			i--
			return setToTrueF(node)
		}
	} else {
		if node.fact.isKnown {
			if node.fact.value == trueF {
				i--
				return setToFalseF(node.right)
			}
			i--
			return setToTrueF(node.right)
		}
	}
	i--
	return nil
}

func andFunc(node *infTree, from *infTree, checked []string) error {
	i++
	fmt.Printf("%*sAndFunc\n", i, " ")
	leftFact := node.left.fact
	rightFact := node.right.fact

	if from == node.head {
		if leftFact.value == falseF || rightFact.value == falseF {
			i--
			return setToFalseF(node)
		}
		if leftFact.value == trueF && rightFact.value == trueF {
			i--
			return setToTrueF(node)
		}
	} else {
		if node.fact.value == trueF {
			err := setToTrueF(node.left)
			if err == nil {
				err = setToTrueF(node.right)
			}
			i--
			return err
		}
		if node.fact.value == falseF {
			if rightFact.value == trueF {
				i--
				return setToFalseF(node.left)
			} else if leftFact.value == trueF {
				i--
				return setToFalseF(node.right)
			}
		}
	}
	i--
	return nil
}

func orFunc(node *infTree, from *infTree, checked []string) error {
	i++
	fmt.Printf("%*sOrFunc\n", i, " ")
	leftFact := node.left.fact
	rightFact := node.right.fact

	if from == node.head {
		if leftFact.value == trueF || rightFact.value == trueF {
			i--
			return setToTrueF(node)
		} else if leftFact.value == falseF && rightFact.value == falseF {
			i--
			return setToFalseF(node)
		}
	} else {
		if node.fact.value == falseF {
			err := setToFalseF(node.left)
			if err == nil {
				err = setToFalseF(node.right)
			}
			i--
			return err
		} else if node.fact.value == trueF {
			if rightFact.value == falseF {
				i--
				return setToTrueF(node.left)
			} else if leftFact.value == falseF {
				i--
				return setToTrueF(node.right)
			}
			setToUnknownF(node.left)
			setToUnknownF(node.right)
		} else if node.fact.value == unknownF {
			setToUnknownF(node.left)
			setToUnknownF(node.right)
		}
	}
	i--
	return nil
}

func xorFunc(node *infTree, from *infTree, checked []string) error {
	i++
	fmt.Printf("%*sXorFunc\n", i, " ")
	leftFact := node.left.fact
	rightFact := node.right.fact

	if from == node.head {
		if leftFact.isKnown && (leftFact.value == rightFact.value) {
			i--
			return setToFalseF(node)
		} else if leftFact.isKnown && rightFact.isKnown {
			i--
			return setToTrueF(node)
		}
	} else {
		if node.fact.value == trueF {
			if leftFact.isKnown {
				if leftFact.value == falseF {
					i--
					return setToTrueF(node.right)
				}
				return setToFalseF(node.right)
			} else if rightFact.isKnown {
				if rightFact.value == falseF {
					i--
					return setToTrueF(node.left)
				}
				i--
				return setToFalseF(node.left)
			}
			setToUnknownF(node.left)
			setToUnknownF(node.right)
		} else if node.fact.value == falseF {
			if leftFact.isKnown {
				if leftFact.value == falseF {
					i--
					return setToFalseF(node.right)
				}
				i--
				return setToTrueF(node.right)
			} else if rightFact.isKnown {
				if rightFact.value == falseF {
					i--
					return setToFalseF(node.left)
				}
				i--
				return setToTrueF(node.left)
			}
			setToUnknownF(node.left)
			setToUnknownF(node.right)
		}
	}
	i--
	return nil
}

func impFunc(node *infTree, from *infTree, checked []string) error {
	i++
	fmt.Printf("%*sImpFunc\n", i, " ")

	if !node.left.fact.isKnown {
		if err := resolve(node.left, node, checked); err != nil {
			i--
			return err
		}
	}
	if node.left.fact.value == trueF {
		i--
		return setToTrueF(node.right)
	}
	i--
	return nil
}

func ioiFunc(node *infTree, from *infTree, checked []string) error {
	i++
	fmt.Printf("%*sIoiFunc\n", i, " ")
	var to *infTree = getOtherSide(node, from)

	if !to.fact.isKnown {
		if err := resolve(to, node, checked); err != nil {
			i--
			return err
		}
	}
	if to.fact.value == falseF {
		i--
		return setToFalseF(from)
	} else if to.fact.value == trueF {
		i--
		return setToTrueF(from)
	}
	i--
	return nil
}
