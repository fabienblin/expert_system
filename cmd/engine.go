package main

import (
	"fmt"
)

/*
 * Run the inference engine
 */
func engine() {
	for _, query := range env.queries {
		backwardInfer(env.tree, query, false)
		if env.factList[query].isTrue && !env.factList[query].isKnown {
			fmt.Printf("%s is undefined\n", query)
		} else {
			fmt.Printf("%s is %v\n", query, env.factList[query].isTrue)
		}
	}
}

/*
 * Backward inference engine
 * Args : current starts on root node, query is the infered fact, implies is true on right side of =>
 */
func backwardInfer(current *infTree, query string, implies bool) {
	if current == nil {
		return
	}
	// fmt.Printf("current = %v\n", current.fact.op)
	if _, ok := env.factList[current.fact.op]; ok { // current is a fact
		if implies && current.fact.isKnown == false {
			current.fact.isTrue = current.head.fact.isTrue
			current.fact.isKnown = current.head.fact.isKnown
		}
		if current.fact.op == query { // current is query
			return
		}

		return
	} else if and == current.fact.op { // current is a +
		if implies {
			current.fact.isTrue = current.head.fact.isTrue
			current.fact.isKnown = current.head.fact.isKnown
		}
		backwardInfer(current.right, query, implies)
		backwardInfer(current.left, query, implies)
		current.fact.isTrue = current.left.fact.isTrue && current.right.fact.isTrue
		current.fact.isKnown = current.left.fact.isKnown && current.right.fact.isKnown
		return
	} else if or == current.fact.op { // current is a |
		if implies {
			current.fact.isTrue = current.head.fact.isTrue
			current.fact.isKnown = false
		}
		backwardInfer(current.right, query, implies)
		backwardInfer(current.left, query, implies)
		current.fact.isTrue = current.left.fact.isTrue || current.right.fact.isTrue
		current.fact.isKnown = current.left.fact.isKnown || current.right.fact.isKnown
		return
	} else if xor == current.fact.op { // current is a ^
		if implies {
			current.fact.isTrue = current.head.fact.isTrue
			current.fact.isKnown = false
		}
		backwardInfer(current.right, query, implies)
		backwardInfer(current.left, query, implies)
		current.fact.isTrue = current.left.fact.isTrue != current.right.fact.isTrue
		current.fact.isKnown = current.left.fact.isKnown != current.right.fact.isKnown
		return
	} else if not == current.fact.op { // current is a !
		if implies {
			current.fact.isTrue = !current.head.fact.isTrue
			current.fact.isKnown = !current.head.fact.isKnown
		}
		backwardInfer(current.right, query, implies)
		current.fact.isTrue = !current.right.fact.isTrue
		current.fact.isKnown = !current.right.fact.isKnown
	} else if ioi == current.fact.op { // current is a <=>
		backwardInfer(current.right, query, true)
		current.fact.isTrue = current.right.fact.isTrue
		current.fact.isKnown = current.right.fact.isKnown
		backwardInfer(current.left, query, true)
		return
	} else if imp == current.fact.op { // current is a =>
		backwardInfer(current.left, query, implies)
		current.fact.isTrue = current.left.fact.isTrue
		current.fact.isKnown = current.left.fact.isKnown
		backwardInfer(current.right, query, true)
		return
	} else if "&" == current.fact.op { // current is a joint &
		backwardInfer(current.right, query, implies)
		backwardInfer(current.left, query, implies)
		return
	} else { // error ?
		return
	}
}
