// Construido como parte da disciplina: Sistemas Distribuidos - PUCRS - Escola Politecnica
//  Professor: Fernando Dotti  (https://fldotti.github.io/)

/*
LANCAR N PROCESSOS EM SHELL's DIFERENTES, UMA PARA CADA PROCESSO, O SEU PROPRIO ENDERECO EE O PRIMEIRO DA LISTA
go run chat.go 127.0.0.1:5001  127.0.0.1:6001    ...   // o processo na porta 5001
go run chat.go 127.0.0.1:6001  127.0.0.1:5001    ...   // o processo na porta 6001
go run chat.go ...  127.0.0.1:6001  127.0.0.1:5001     // o processo na porta ...
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "SD/BEB"
)

var player1Score int
var player2Score int

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify at least one address:port!")
		fmt.Println("go run chatBEB.go 127.0.0.1:5001  127.0.0.1:6001   127.0.0.1:7001")
		fmt.Println("go run chatBEB.go 127.0.0.1:6001  127.0.0.1:5001   127.0.0.1:7001")
		fmt.Println("go run chatBEB.go 127.0.0.1:7001  127.0.0.1:6001   127.0.0.1:5001")
		return
	}
	var registro []string

	addresses := os.Args[1:]
	fmt.Println(addresses)

	beb := BestEffortBroadcast_Module{
		Req: make(chan BestEffortBroadcast_Req_Message),
		Ind: make(chan BestEffortBroadcast_Ind_Message)}

	//beb.Init(addresses[0])
	beb.InitD(addresses[0], false)

	fmt.Println(`
	######                          ######                                  #####                                       
	#     #  ####   ####  #    #    #     #   ##   #####  ###### #####     #     #  ####  #  ####   ####   ####  #####  
	#     # #    # #    # #   #     #     #  #  #  #    # #      #    #    #       #    # # #      #      #    # #    # 
	######  #    # #      ####      ######  #    # #    # #####  #    #     #####  #      #  ####   ####  #    # #    # 
	#   #   #    # #      #  #      #       ###### #####  #      #####           # #      #      #      # #    # #####  
	#    #  #    # #    # #   #     #       #    # #      #      #   #     #     # #    # # #    # #    # #    # #   #  
	#     #  ####   ####  #    #    #       #    # #      ###### #    #     #####   ####  #  ####   ####   ####  #    # 
	`)

	// enviador de broadcasts
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		var msg string

		for {
			printCommands()

			if scanner.Scan() {
				msg = scanner.Text()
				msg += "§" + addresses[0]
			}

			req := BestEffortBroadcast_Req_Message{
				Addresses: addresses[0:],
				Message:   string(msg)}
			beb.Req <- req // ENVIA PARA TODOS PROCESSOS ENDERECADOS NO INICIO
		}
	}()

	// receptor de broadcasts
	go func() {
		currentRoundValue := ""
		playerText := ""
		for {
			in := <-beb.Ind // RECEBE MENSAGEM DE QUALQUER PROCESSO
			message := strings.Split(in.Message, "§")
			in.From = message[1]
			registro = append(registro, in.Message)
			in.Message = message[0]
			playerText = message[0]

			if currentRoundValue != "" {
				result(currentRoundValue, playerText)
				fmt.Println("=========================")
				fmt.Println("P1 score: ", player1Score)
				fmt.Println("P2 score: ", player2Score)
				fmt.Println("=========================")
				printCommands()
				currentRoundValue = ""
			} else {
				currentRoundValue = playerText
			}
		}
	}()
	blq := make(chan int)
	<-blq
}

func result(p1 string, p2 string) {

	combinedResults := p1 + p2
	switch combinedResults {
	case "rs", "pr", "sp":
		fmt.Println("P1 Won 🏆")
		player1Score++
	case "rp", "ps", "sr":
		fmt.Println("P1 Lost 💩")
		player2Score++
	case "rr", "pp", "ss":
		fmt.Println("Match Draws 🤙🏻")
	}
}

func printCommands() {
	fmt.Println(`Press "r" for rock`)
	fmt.Println(`Press "p" for paper`)
	fmt.Println(`Press "s" for scissor`)
	fmt.Println()
}
