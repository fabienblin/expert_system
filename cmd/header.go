/* ************************************************************************** */
/*                                                          LE - /            */
/*                                                              /             */
/*   header.go                                        .::    .:/ .      .::   */
/*                                                 +:+:+   +:    +:  +:+:+    */
/*   By: jojomoon <jojomoon@student.le-101.fr>      +:+   +:    +:    +:+     */
/*                                                 #+#   #+    #+    #+#      */
/*   Created: 2019/10/30 17:52:04 by jmonneri     #+#   ##    ##    #+#       */
/*   Updated: 2020/02/06 00:14:20 by jojomoon    ###    #+. /#+    ###.fr     */
/*                                                         /                  */
/*                                                        /                   */
/* ************************************************************************** */
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
	trueF       int    = 2
	falseF      int    = 1
	unknownF    int    = 0
	defaultF    int    = -1
	errorF			int		 = -2
)

var verbose bool = false

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
	isKnown bool
	value   int
	fixed		bool
}

var env struct {
	rules        []string
	initialFacts []string
	queries      []string
	trees        []*infTree
	factList     map[string]*fact
	tree         *infTree
}

var opeFunc map[string]func(*infTree, *infTree, []string) error

func init() {
	opeFunc = map[string]func(*infTree, *infTree, []string) error{
		and: andFunc,
		not: notFunc,
		xor: xorFunc,
		or:  orFunc,
		imp: impFunc,
		ioi: ioiFunc,
	}
}

var opeFuncFor map[string]func(*infTree) (bool, error)

func init() {
	opeFuncFor = map[string]func(*infTree) (bool, error){
		and: andFuncFor,
		not: notFuncFor,
		xor: xorFuncFor,
		or:  orFuncFor,
		imp: impFuncFor,
		ioi: ioiFuncFor,
	}
}

var output = [3]string {
	"Undetermined",
	"False",
	"True",
}