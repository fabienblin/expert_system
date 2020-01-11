/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/10/30 17:52:16 by jmonneri          #+#    #+#             */
/*   Updated: 2020/01/11 01:51:20 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"flag"
	"log"
	"os"
)

func initEnv() {
	env.factList = make(map[string]*fact)
	env.initialFacts = nil
	env.trees = nil
}

// flags f et v boolens

func main() {

	flagVerbose := flag.Bool("v", false, "verbose mode")
	flagForward := flag.Bool("f", false, "forward mode")

	flag.Parse()
	verbose = *flagVerbose
	args := flag.Args()
	if len(args) == 0 { // dynamic ruleset
		for {
			initEnv()
			parseDynamic()

			initAllFacts()
			buildTree()

			printNode(env.tree, 4, nil)
			engine(*flagForward)
		}
	} else if len(args) == 1 { // file ruleset

		initEnv()
		parseFile(args[0])

		initAllFacts()
		buildTree()

		printNode(env.tree, 4, nil)
		engine(*flagForward)
	} else { // error
		log.Fatal("\nUsage: ./bin/expert_system [OPTIONS] [FILE]\n[OPTIONS]: -v = verbose mode ; -f = forward chaining mode\n[FILE]: if not represented, start dynamic mode")
		os.Exit(1)
	}
}
