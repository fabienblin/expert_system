package main

import (
	"fmt"
)

func engine() {
	fmt.Printf("Engine\n")
	for _, query := range env.queries {
		backwardInfer(env.tree, env.tree, query)
		fmt.Printf("%s is %v\n", query, env.factList[query].isTrue)
	}
}

func backwardInfer(root *infTree, current *infTree, query string) {
	if current == nil {
		return
	} else if current == root {
		return
	} else if _, ok := env.factList[current.fact.op]; ok { // current is a fact
		if current.fact.op == query { // current is query
			return
		} else { // current is another fact
			return
		}
	} else if and == current.fact.op { // current is a +
		backwardInfer(root, current.right, query)
		backwardInfer(root, current.left, query)
		if current.left.fact.isTrue && current.right.fact.isTrue {
			current.fact.isTrue = true
			current.fact.isKnown = true
		}
		return
	} else if or == current.fact.op { // current is a |
		backwardInfer(root, current.right, query)
		backwardInfer(root, current.left, query)
		if current.left.fact.isTrue || current.right.fact.isTrue {
			current.fact.isTrue = true
			current.fact.isKnown = true
		}
		return
	} else if xor == current.fact.op { // current is a ^
		backwardInfer(root, current.right, query)
		backwardInfer(root, current.left, query)
		if current.left.fact.isTrue != current.right.fact.isTrue {
			current.fact.isTrue = true
			current.fact.isKnown = true
		}
		return
	} else if not == current.fact.op { // current is a !
		backwardInfer(root, current.right, query)
		return
	} else if ioi == current.fact.op { // current is a <=>
		backwardInfer(root, current.right, query)
		backwardInfer(root, current.left, query)
		return
	} else if imp == current.fact.op { // current is a =>
		backwardInfer(root, current.right, query)
		backwardInfer(root, current.left, query)
		return
	} else if "&" == current.fact.op { // current is a joint &
		backwardInfer(root, current.right, query)
		backwardInfer(root, current.left, query)
		return
	} else { // error ?
		return
	}
}
