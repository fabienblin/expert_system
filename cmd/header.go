package main

const (
	openBra     string = "("
	closeBra    string = ")"
	not         string = "!"
	and         string = "+"
	or          string = "|"
	xor         string = "^"
	imp         string = "=>"
	ioi         string = "<=>"
	com         string = "#"
	factSymbol  string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	factDeclar  string = "="
	queryDeclar string = "?"
	trueF       int    = 1
	falseF      int    = 0
	unknownF    int    = -1
)

type nodeInfo int

const (
	noInfo nodeInfo = iota + 1
	skipClimbUp
	rightAssociative
)

type precedence int

const (
	openBraPre precedence = iota + 1
	closeBraPre
	impPre
	ioiPre
	xorPre
	orPre
	andPre
	notPre
	factPre
)

type infTree struct {
	head       *infTree
	left       *infTree
	right      *infTree
	precedence precedence
	fact       *fact
}

type fact struct {
	op      string
	isTrue  bool
	isKnown bool
}

var env struct {
	rules        []string
	initialFacts []string
	queries      []string
	trees        []*infTree
	factList     map[string]*fact
	tree         *infTree
}
