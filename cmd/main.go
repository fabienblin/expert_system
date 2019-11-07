package main

func main() {

	parse()

	// default fact Z={"Z", false, false}
	// two methods to access any fact.
	// 1: from env.factList (parser)
	// 2: from env.trees (engine)
	env.factList["Z"].isKnown = true
	env.trees[0].right.fact.isTrue = true

	// trees

	// tree

	engine()
}
