/*
*

	PING PONG Game - Ponto de ebulição da Agua de Kelvin em Celsius.

	Projeto com objetivo de estudo e aprendizado da linguagem GO

	O projeto faz parte do laboratorio de estudos do curso da DIO.me

	Professora Tenille Martins - Curso de GO Leng.

	Autor: Paulo Roberto - PRSant0s

	30/04/2024

	Versão 00.1

*
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	var canal chan string = make(chan string)

	go ping(canal)
	go imprimir(canal)
	go pong(canal)

	var entrada string
	fmt.Scanln(&entrada)
}

func ping(canal chan string) {
	for i := 0; ; i++ {
		canal <- "ping"
	}

}

func pong(canal chan string) {
	for i := 0; ; i++ {
		canal <- "pong"
	}

}
func imprimir(canal chan string) {
	for {
		msg := <-canal
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
