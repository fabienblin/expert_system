/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   engine.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/10/30 17:51:53 by jmonneri          #+#    #+#             */
/*   Updated: 2020/01/10 19:18:01 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

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
			fmt.Printf("%s is ambiguous\n", query)
		} else if !env.factList[query].isKnown {
			fmt.Printf("%s is unknown\n", query)
		} else {
			fmt.Printf("%s is %v\n", query, env.factList[query].isTrue)
		}
	}
}
