/* ************************************************************************** */
/*                                                          LE - /            */
/*                                                              /             */
/*   engine.go                                        .::    .:/ .      .::   */
/*                                                 +:+:+   +:    +:  +:+:+    */
/*   By: fablin <fablin@student.le-101.fr>          +:+   +:    +:    +:+     */
/*                                                 #+#   #+    #+    #+#      */
/*   Created: 2019/10/30 17:51:53 by jmonneri     #+#   ##    ##    #+#       */
/*   Updated: 2020/02/21 14:12:35 by fablin      ###    #+. /#+    ###.fr     */
/*                                                         /                  */
/*                                                        /                   */
/* ************************************************************************** */
package main
// Stocker dans le node (pas dans le fact) un "checked" que l'on vérifie au moment où l'on passe dessus au lieu de vérifier le checked déjà présent. De plus, en écrasant, il faut faire péter une erreur si on est dans un imp et continuer d'écraser si on est dans un ioi
import (
	"errors"
	"fmt"
	"log"
	"strings"
)

/*
 * Run the inference engine
 */
func engine(flagForward bool) {
	if flagForward {
		computeTrees()
		for _, query := range env.queries {
			if env.factList[query].value == defaultF {
				env.factList[query].value = falseF
			}
			fmt.Printf("# solution %s = %d\n", env.factList[query].op, env.factList[query].value)
		}
	} else {
		for _, query := range env.queries {
			if err := backwardChaining(env.factList[query], []string{}); err != nil {
				fmt.Printf("# solution %s = %s\n", env.factList[query].op, err)
			} else {
				fmt.Printf("# solution %s = %s\n", env.factList[query].op, output[env.factList[query].value])
			}
		}
	}
	if verbose {
		fmt.Printf(getNode(env.tree, 4, nil))
	}
}

func backwardChaining(query *fact, checked []string) error {
	// On check que le fact n'ait pas déjà été demandé (anti-boucle).
	if stringInSlice(query.op, checked) {
		return nil
	}
<<<<<<< HEAD

	// fmt.Printf("backward current = %v\timplies = %v\tdepend=%v\n", current.fact.op, implies, depend)
	if depend {
		if implies { // (2) go to head
			if current.fact.op != ioi && current.fact.op != imp { // loop to head and propagate isTrue to query fact
				backwardInfer(current.head, query, implies, depend)
				// (3) define right side
				if current.fact.isKnown == false && current.head.fact.isKnown == true {
					if _, ok := env.factList[current.fact.op]; ok { // define queried fact
						current.fact.isTrue = current.head.fact.isTrue
						current.fact.isKnown = current.head.fact.isKnown
					} else if and == current.fact.op { // current is a +
						current.fact.isTrue = current.head.fact.isTrue
						current.fact.isKnown = current.head.fact.isKnown
					} else if or == current.fact.op { // current is a |
						current.fact.isTrue = current.head.fact.isTrue
						current.fact.isKnown = current.head.fact.isKnown
					} else if xor == current.fact.op { // current is a ^
						current.fact.isTrue = current.head.fact.isTrue
						current.fact.isKnown = current.head.fact.isKnown
					} else if not == current.fact.op { // current is a !
						current.fact.isTrue = !current.head.fact.isTrue
						current.fact.isKnown = current.head.fact.isKnown
					}
				} else if _, ok := env.factList[current.fact.op]; ok && current.fact.isKnown == true && current.head.fact.isKnown == true && current.head.fact.isTrue == false {
					log.Printf("Paradox in fact definition of %v\n", current.fact.op)
					// os.Exit(1)
				}
			} else { // (2.1 and 2.2) if => or <=>
				if current.fact.op == imp {
					forwardInfer(current, false)
				} else if current.fact.op == ioi {
					forwardInfer(current.right, false)
					forwardInfer(current.left, false)
				}
				current.fact.isKnown = true
				current.fact.isTrue = current.right.fact.isTrue || current.left.fact.isTrue
			}
<<<<<<< HEAD
		}
	} else {
		if implies {
			if _, ok := env.factList[current.fact.op]; ok { // current is a fact
				if current.fact.op == query { // current is query
					backwardInfer(current, query, implies, true) // (1.1)
				}
			} else {
				backwardInfer(current.right, query, implies, depend)
				backwardInfer(current.left, query, implies, depend)
=======
		} else {

=======
	checked = append(checked, query.op)
	// On trouve les règles définissant la query
	if verbose {
		fmt.Printf("Searching for rules defining %s\n", query.op)
	}
	for _, rule := range env.trees {
		if err := digInRule(query, rule, checked); err != nil {
			return err
>>>>>>> a3141c1340f428b7417537b69dc99d328e70032e
		}
	}
	if query.value == defaultF {
		if verbose {
			fmt.Printf("Set %s to false by default: not enought information\n", query.op)
		}
		query.value = falseF
		query.isKnown = true
		query.fixed = false
	}
	if verbose {
		fmt.Printf("Stop searching %s\n", query.op)
	}
	return nil
}

func digInRule(fact *fact, node *infTree, checked []string) error {
	if strings.Contains(factSymbol, node.fact.op) {
		if node.fact == fact {
			if verbose {
				fmt.Printf("Rule found for %s searched:\n%s", fact.op, getContextRule2(node))
>>>>>>> master
			}
			if err := resolve(node, node, checked); err != nil {
				node.fact.value = errorF
				node.fact.isKnown = false
				return err
			}
		}
<<<<<<< HEAD
	}
=======
		return nil
	}
	if node.fact.op != imp && node.left != nil {
		if err := digInRule(fact, node.left, checked); err != nil {
			return err
		}
	}
	return digInRule(fact, node.right, checked)
>>>>>>> master
}

<<<<<<< HEAD
/*
 * INCOMPLETE
 * Forward inference engine
 * Args : current starts on root node, implies is true on right side of =>
 * When implies is true, information is infered from head node
 */
func forwardInfer(current *infTree, implies bool) {
	if current == nil {
		return
	}
	// fmt.Printf("forward current = %v\timplies = %v\n", current.fact.op, implies)
	if implies {
		if _, ok := env.factList[current.fact.op]; ok { // current is a fact
			if current.fact.isKnown == false && current.head.fact.isKnown == true {
				current.fact.isTrue = current.head.fact.isTrue
				current.fact.isKnown = current.head.fact.isKnown
			} else if _, ok := env.factList[current.fact.op]; ok && current.fact.isKnown == true && current.head.fact.isKnown == true && current.head.fact.isTrue == false {
				log.Printf("Paradox in fact definition of %v\n", current.fact.op)
				// os.Exit(1)
			}
		} else if and == current.fact.op { // current is a +
			current.fact.isTrue = current.head.fact.isTrue
			current.fact.isKnown = current.head.fact.isKnown
			forwardInfer(current.right, implies)
			forwardInfer(current.left, implies)
		} else if or == current.fact.op { // current is a |
			current.fact.isTrue = current.head.fact.isTrue
			current.fact.isKnown = false
			forwardInfer(current.right, implies)
			forwardInfer(current.left, implies)
		} else if xor == current.fact.op { // current is a ^
			current.fact.isTrue = current.head.fact.isTrue
			current.fact.isKnown = false
			forwardInfer(current.right, implies)
			forwardInfer(current.left, implies)
		} else if not == current.fact.op { // current is a !
			current.fact.isTrue = !current.head.fact.isTrue
			current.fact.isKnown = current.head.fact.isKnown
			forwardInfer(current.right, implies)
		} else if ioi == current.fact.op { // current is a <=>
			forwardInfer(current.right, implies)
			forwardInfer(current.left, implies)
		} else if imp == current.fact.op { // current is a =>
			forwardInfer(current.left, implies)
		}
	} else {
		if and == current.fact.op { // current is a +
			forwardInfer(current.right, implies)
			forwardInfer(current.left, implies)
			current.fact.isTrue = current.left.fact.isTrue && current.right.fact.isTrue
			current.fact.isKnown = current.right.fact.isKnown && current.left.fact.isKnown
		} else if or == current.fact.op { // current is a |
			forwardInfer(current.right, implies)
			forwardInfer(current.left, implies)
			current.fact.isTrue = current.left.fact.isTrue || current.right.fact.isTrue
			current.fact.isKnown = current.right.fact.isKnown && current.left.fact.isKnown
		} else if xor == current.fact.op { // current is a ^
			forwardInfer(current.right, implies)
			forwardInfer(current.left, implies)
			current.fact.isTrue = current.left.fact.isTrue != current.right.fact.isTrue
			current.fact.isKnown = current.right.fact.isKnown && current.left.fact.isKnown
		} else if not == current.fact.op { // current is a !
			forwardInfer(current.right, implies)
			current.fact.isTrue = !current.right.fact.isTrue
			current.fact.isKnown = current.right.fact.isKnown
		} else if ioi == current.fact.op { // current is a <=>
			forwardInfer(current.left, implies)
			forwardInfer(current.right, implies)
		} else if imp == current.fact.op { // current is a =>
			forwardInfer(current.left, implies)
			current.fact.isTrue = current.left.fact.isTrue
			current.fact.isKnown = current.left.fact.isKnown
		} else if "&" == current.fact.op { // current is a joint &
			forwardInfer(current.right, implies)
			forwardInfer(current.left, implies)
=======
func resolve(node *infTree, from *infTree, checked []string) error {
	var err error = nil
	if node == nil{
		return nil
	}
	if node.fact.value == errorF {
		return errors.New("Error : There was already a contradiction for the fact " + node.fact.op)
	}
	if from != node.head && !(node.fact.op == imp || node.fact.op == ioi) {
		err = resolve(node.head, node, checked)
		if node == from || err != nil {
			return err
		}
	}
	if strings.Contains(factSymbol, node.fact.op) {
		if !node.fact.isKnown {
			return backwardChaining(node.fact, checked)
>>>>>>> a3141c1340f428b7417537b69dc99d328e70032e
		}
		return nil
	}
	if from == node.head {
		err := resolve(node.left, node, checked)
		if err == nil {
			err = resolve(node.right, node, checked)
		}
		if err != nil {
			return err
		}
	}
	// On lance la fonction de l'operateur
	return opeFunc[node.fact.op](node, from, checked)
}

func consequensesRelaunch() error {
	for _, fact := range env.factList {
		if fact.isKnown && !fact.fixed {
			if err := backwardChaining(fact, []string{}); err != nil {
				return err
			}
		}
	}
	return nil
}

func seekForOtherSide(node *infTree, checked []string) (bool, error) {
	var otherSide *infTree = getOtherSide(node.head, node)
	if err := resolve(otherSide, node.head, checked); err != nil {
		return false, err
	}
 	if otherSide.fact.isKnown {
		return true, opeFunc[node.head.fact.op](node.head, node, checked)
	}
	return false, nil
}