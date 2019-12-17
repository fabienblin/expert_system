package main

import (
	"fmt"
)

/*
 * Run the inference engine
 */
func engine() {
	//forwardInfer(env.tree, false)
	for _, query := range env.queries {
		fmt.Print("")
		backwardInfer(env.tree, query, false, false)
		if env.factList[query].isTrue && !env.factList[query].isKnown {
			fmt.Printf("%s is undefined\n", query)
		} else {
			fmt.Printf("%s is %v\n", query, env.factList[query].isTrue)
		}
	}
}

/*
 * Backward inference engine recursive
 * Args : current starts on root node (env.tree),
 * query is the queried fact
 * implies is true on right side of =>, true on both sides of <=>
 * depend is true when searching for fact dependencies
 *
 * 1 - find the queried fact on right side of => or both sides of <=> (implies=true)
 *	1.1 - fact found : backward (depend=true)
 * 2 - go back to first => or <=> (implies=false)
 *	2.1 - if => : forward on left side
 *	2.2 - if <=> : forward on both sides
 * 3 - define queried fact
 */
func backwardInfer(current *infTree, query string, implies bool, depend bool) {
	if current == nil {
		return
	}

	// fmt.Printf("backward current = %v\timplies = %v\tdepend=%v\n", current.fact.op, implies, depend)
	if depend {
		if implies { // (2) go to head
			if current.fact.op != ioi && current.fact.op != imp { // loop to head and propagate isTrue to query fact
				// fmt.Print("going up\n")
				backwardInfer(current.head, query, implies, depend)
				// fmt.Print("going down\n")
				// (3) define right side
				if current.fact.isKnown == false {
					if _, ok := env.factList[current.fact.op]; ok {
						current.fact.isKnown = current.head.fact.isKnown
						current.fact.isTrue = current.head.fact.isTrue
					} else if and == current.fact.op { // current is a +
						current.fact.isTrue = current.head.fact.isTrue
						current.fact.isKnown = current.head.fact.isKnown
					} else if or == current.fact.op { // current is a |
						current.fact.isTrue = current.head.fact.isTrue
						current.fact.isKnown = false
					} else if xor == current.fact.op { // current is a ^
						current.fact.isTrue = current.head.fact.isTrue
						current.fact.isKnown = false
					} else if not == current.fact.op { // current is a !
						current.fact.isTrue = !current.head.fact.isTrue
						current.right.fact.isKnown = current.head.fact.isKnown
					}

				}
			} else { // if => or <=>
				forwardInfer(current, false)
			}
		} else {

		}
		return
	} else {
		if implies {
			if _, ok := env.factList[current.fact.op]; ok { // current is a fact
				if current.fact.op == query { // current is query
					// fmt.Print("found query ", query, "\n")
					backwardInfer(current, query, implies, true) // (1.1)
				}
			} else {
				backwardInfer(current.right, query, implies, depend)
				backwardInfer(current.left, query, implies, depend)
			}
		} else { // (1) start here
			// fmt.Print(current.fact.op, query, implies, depend, "\n")
			if current.fact.op == "&" {
				backwardInfer(current.right, query, implies, depend)
				backwardInfer(current.left, query, implies, depend)
			} else if current.fact.op == ioi {
				backwardInfer(current.right, query, true, depend)
				backwardInfer(current.left, query, true, depend)
			} else if current.fact.op == imp {
				backwardInfer(current.right, query, true, depend)
			}
		}
		return
	}
	return
}

/*
 * INCOMPLETE
 * Forward inference engine
 * Args : current starts on root node, implies is true on right side of =>
 */
func forwardInfer(current *infTree, implies bool) {
	if current == nil {
		return
	}
	// fmt.Printf("forward current = %v\timplies = %v\n", current.fact.op, implies)
	if implies {
		if _, ok := env.factList[current.fact.op]; ok { // current is a fact
			if current.fact.isKnown == false {
				current.fact.isTrue = current.head.fact.isTrue
				current.fact.isKnown = current.head.fact.isKnown
			}
		} else if and == current.fact.op { // current is a +
			forwardInfer(current.right, implies)
			forwardInfer(current.left, implies)
			current.fact.isTrue = current.head.fact.isTrue
			current.fact.isKnown = current.head.fact.isKnown
		} else if or == current.fact.op { // current is a |
			forwardInfer(current.right, implies)
			forwardInfer(current.left, implies)
			current.fact.isTrue = current.head.fact.isTrue
			current.fact.isKnown = false
		} else if xor == current.fact.op { // current is a ^
			forwardInfer(current.right, implies)
			forwardInfer(current.left, implies)
			current.fact.isTrue = current.head.fact.isTrue
			current.fact.isKnown = false
		} else if not == current.fact.op { // current is a !
			forwardInfer(current.right, implies)
			current.fact.isTrue = !current.head.fact.isTrue
			current.right.fact.isKnown = current.head.fact.isKnown
		} else if ioi == current.fact.op { // current is a <=>
			forwardInfer(current.right, implies)
			forwardInfer(current.left, implies)
		} else if imp == current.fact.op { // current is a =>

			forwardInfer(current.left, false)
		}
	} else {
		if and == current.fact.op { // current is a +
			forwardInfer(current.right, implies)
			forwardInfer(current.left, implies)
			current.fact.isTrue = current.left.fact.isTrue && current.right.fact.isTrue
			current.fact.isKnown = true
		} else if or == current.fact.op { // current is a |
			forwardInfer(current.right, implies)
			forwardInfer(current.left, implies)
			current.fact.isTrue = current.left.fact.isTrue || current.right.fact.isTrue
			current.fact.isKnown = true
		} else if xor == current.fact.op { // current is a ^
			forwardInfer(current.right, implies)
			forwardInfer(current.left, implies)
			current.fact.isTrue = current.left.fact.isTrue != current.right.fact.isTrue
			current.fact.isKnown = true
		} else if not == current.fact.op { // current is a !
			forwardInfer(current.right, implies)
			current.fact.isTrue = !current.right.fact.isTrue
			current.fact.isKnown = true
		} else if ioi == current.fact.op { // current is a <=>
			forwardInfer(current, true)
			current.fact.isTrue = current.right.fact.isTrue || current.left.fact.isTrue
			current.fact.isKnown = true
		} else if imp == current.fact.op { // current is a =>
			forwardInfer(current, true)
			current.fact.isTrue = current.left.fact.isTrue
			current.fact.isKnown = true
		} else if "&" == current.fact.op { // current is a joint &
			forwardInfer(current.right, implies)
			forwardInfer(current.left, implies)
		}
	}

	return
}
