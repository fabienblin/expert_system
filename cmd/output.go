/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   output.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/10/30 17:51:41 by jmonneri          #+#    #+#             */
/*   Updated: 2020/01/09 17:30:42 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func outputError(err error) {
	fmt.Println("Ca marche pas:", err)
}

func output() {
	fmt.Printf("Ca a l'air de marcher")
}
