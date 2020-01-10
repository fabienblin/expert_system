/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/10/30 17:52:16 by jmonneri          #+#    #+#             */
/*   Updated: 2020/01/10 19:24:09 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"log"
	"os"
)

func initEnv() {
	env.factList = make(map[string]*fact)
	env.initialFacts = nil
	env.trees = nil
}

func main() {

	if len(os.Args) == 1 { // dynamic ruleset
		for {
			initEnv()
			parseDynamic()

			initAllFacts()
			buildTree()

			engine()
			printNode(env.tree, 8, nil)
		}
	} else if len(os.Args) == 2 { // file ruleset
		initEnv()
		parseFile(os.Args[1])

		initAllFacts()
		buildTree()

		engine()
		printNode(env.tree, 8, nil)
	} else { // error
		log.Fatal("Error. Retry later ...\n")
		os.Exit(1)
	}
}
