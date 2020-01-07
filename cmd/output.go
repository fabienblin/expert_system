/* ************************************************************************** */
/*                                                          LE - /            */
/*                                                              /             */
/*   output.go                                        .::    .:/ .      .::   */
/*                                                 +:+:+   +:    +:  +:+:+    */
/*   By: jojomoon <jojomoon@student.le-101.fr>      +:+   +:    +:    +:+     */
/*                                                 #+#   #+    #+    #+#      */
/*   Created: 2019/10/30 17:51:41 by jmonneri     #+#   ##    ##    #+#       */
/*   Updated: 2020/01/07 15:32:04 by jojomoon    ###    #+. /#+    ###.fr     */
/*                                                         /                  */
/*                                                        /                   */
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
