/* ************************************************************************** */
/*                                                          LE - /            */
/*                                                              /             */
/*   main.go                                          .::    .:/ .      .::   */
/*                                                 +:+:+   +:    +:  +:+:+    */
/*   By: fablin <fablin@student.le-101.fr>          +:+   +:    +:    +:+     */
/*                                                 #+#   #+    #+    #+#      */
/*   Created: 2019/10/30 17:52:16 by jmonneri     #+#   ##    ##    #+#       */
/*   Updated: 2020/02/21 14:57:03 by fablin      ###    #+. /#+    ###.fr     */
/*                                                         /                  */
/*                                                        /                   */
/* ************************************************************************** */
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func initEnv() {
	env.factList = make(map[string]*fact)
	env.initialFacts = nil
	env.trees = nil
	env.tree = nil
}

// flags f et v boolens

func main() {

	flagVerbose := flag.Bool("v", false, "verbose mode")
	flagForward := flag.Bool("f", false, "forward mode")

	flag.Parse()
	verbose = *flagVerbose
	args := flag.Args()
	if len(args) == 0 { // dynamic ruleset
		fmt.Printf("Using dynamic mode. \nPlease write the rules followed by initial facts then your query.\nType 'exit' to stop.\nType 'run' to run inference engine.\n")
		for {
			fmt.Printf("You may redefine known facts to retry the query.\n")
			initEnv()
			parseDynamic()

			initAllFacts()
			buildTree()

			if verbose {
				printNode(env.tree, 4, nil)
			}
			engine(*flagForward)
		}
	} else if len(args) == 1 { // file ruleset

		initEnv()
		parseFile(args[0])

		initAllFacts()
		buildTree()

		if verbose {
			fmt.Println(getNode(env.tree, 4, nil))
		}
		engine(*flagForward)
	} else { // error
		log.Fatal("\nUsage: ./bin/expert_system [OPTIONS] [FILE]\n[OPTIONS]: -v = verbose mode ; -f = forward chaining mode\n[FILE]: if not represented, start dynamic mode")
		os.Exit(1)
	}
}
