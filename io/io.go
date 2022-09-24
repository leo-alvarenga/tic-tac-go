package io

import (
	"fmt"

	"github.com/leo-alvarenga/tic-tac-go/ng"
)

func CLI(board *ng.Board) {
	ShowBoard(board.GetBoard())
}

func clearStdin() {

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
