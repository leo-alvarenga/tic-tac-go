package io

import (
	"fmt"

	"github.com/leo-alvarenga/tic-tac-go/cpu"
	"github.com/leo-alvarenga/tic-tac-go/ng"
)

func CLI(board *ng.Board) {
	gameLoop(board)
}

func gameLoop(board *ng.Board) {
	var pos int
	playersTurn := true

	cpuMove := make(chan int)
	marker := 'x'

	ShowBoard(board.GetBoard())
	for !board.HasTheGameFinished() {
		if playersTurn {
			fmt.Print("Make a move (0-8)")
			fmt.Scanf("%d", &pos)

			if pos > 9 {
				pos = pos % 9
			} else if pos < 0 {
				pos = 0
			}

			go cpu.GetNextMove(board, cpuMove)
			marker = 'x'
		} else {
			pos = <-cpuMove
			marker = 'o'
		}

		s, _ := board.MarkThisPosition(pos, marker)

		if s != "" {
			println(s)
		}

		ShowBoard(board.GetBoard())
		playersTurn = !playersTurn
	}
}

func clearStdin() {
	println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
}

func ShowBoard(b [9]rune) {
	clearStdin()

	for i := 0; i < 7; i += 3 {
		print("\t")

		for j := i; j < i+3; j++ {
			fmt.Printf(" %c ", b[j])

			if j != i+2 {
				print("|")
			}
		}

		if i < 6 {
			println("\n\t-----------")
		} else {
			println()
		}
	}
}
