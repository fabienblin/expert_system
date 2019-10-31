/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/10/30 17:52:16 by jmonneri          #+#    #+#             */
/*   Updated: 2019/10/30 17:53:05 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) == 1 { // dynamic ruleset
		parseDynamic()
	} else if len(os.Args) == 2 { // file ruleset
		parseFile(os.Args[1])
	} else { // error
		fmt.Println("Error. Retry later ...")
		os.Exit(1)
	}
	for _, tree := range env.trees {
		fmt.Printf("\nROOT : \n----------------------------\n")
		printNode(&tree, 4)
	}
	engine()
}
