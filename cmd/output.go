/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   output.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/10/30 17:51:41 by jmonneri          #+#    #+#             */
/*   Updated: 2019/11/26 18:29:56 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func outputError(err error) {
	fmt.Println("Ca marche pas:", err)
}

func output() {
	printNode(env.tree, 4)
	fmt.Printf("Ca a l'air de marcher")
}
