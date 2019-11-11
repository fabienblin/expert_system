/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   output.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jmonneri <jmonneri@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/10/30 17:51:41 by jmonneri          #+#    #+#             */
/*   Updated: 2019/11/11 18:42:30 by jmonneri         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func outputError(err error) {
	fmt.Printf("Ca marche pas, normal les fonctions operateur sont pas dev")
}

func output() {
	fmt.Printf("Ca a l'air de marcher")
}
