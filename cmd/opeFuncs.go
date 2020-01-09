/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   opeFuncs.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/11/11 14:34:50 by jmonneri          #+#    #+#             */
/*   Updated: 2020/01/09 20:28:46 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"errors"
)

func setToTrueF(node *infTree) error {
	if node.fact.value == falseF {
		return errors.New("Error: Contradiction in the facts")
	} else if !node.fact.isKnown {
		node.fact.isKnown = true
		node.fact.value = trueF
	}
	return nil
}

func setToFalseF(node *infTree) error {
	if node.fact.value == trueF {
		return errors.New("Error: Contradiction in the facts")
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
	if from == node.head {
		if node.right.fact.isKnown {
			if node.right.fact.value == trueF {
				return setToFalseF(node)
			}
			return setToTrueF(node)
		}
	} else {
		if node.fact.isKnown {
			if node.fact.value == trueF {
				return setToFalseF(node.right)
			}
			return setToTrueF(node.right)
		}
	}
	return nil
}

func andFunc(node *infTree, from *infTree, checked []string) error {
	leftFact := node.left.fact
	rightFact := node.right.fact

	if from == node.head {
		if leftFact.value == falseF || rightFact.value == falseF {
			return setToFalseF(node)
		}
		if leftFact.value == trueF && rightFact.value == trueF {
			return setToTrueF(node)
		}
	} else {
		if node.fact.value == trueF {
			err := setToTrueF(node.left)
			if err == nil {
				err = setToTrueF(node.right)
			}
			return err
		}
		if node.fact.value == falseF {
			if rightFact.value == trueF {
			} else if leftFact.value == trueF {
				return setToFalseF(node.right)
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
			return setToTrueF(node)
		} else if leftFact.value == falseF && rightFact.value == falseF {
			return setToFalseF(node)
		}
	} else {
		if node.fact.value == falseF {
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
		} else if node.fact.value == unknownF {
			setToUnknownF(node.left)
			setToUnknownF(node.right)
		}
	}
	return nil
}

func xorFunc(node *infTree, from *infTree, checked []string) error {
	leftFact := node.left.fact
	rightFact := node.right.fact

	if from == node.head {
		if leftFact.isKnown && (leftFact.value == rightFact.value) {
			return setToFalseF(node)
		} else if leftFact.isKnown && rightFact.isKnown {
			return setToTrueF(node)
		}
	} else {
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
		return setToTrueF(node.right)
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
		return setToFalseF(from)
	} else if to.fact.value == trueF {
		return setToTrueF(from)
	}
	return nil
}
