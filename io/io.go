package io

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/leo-alvarenga/tic-tac-go/cpu"
	"github.com/leo-alvarenga/tic-tac-go/ng"
)

func CLI(board *ng.Board) {
	args := os.Args

	if len(args) > 1 {
		switch args[1] {
		case "hard":
			gameLoopAI(board, new(cpu.HardCPU))

		case "2p":
			gameLoopTwoPlayers(board)

		case "two":
			gameLoopTwoPlayers(board)

		default:
			gameLoopAI(board, new(cpu.EasyCPU))
		}
	} else {
		gameLoopAI(board, new(cpu.EasyCPU))
	}
}

func gameLoopAI(board *ng.Board, ai cpu.CPU) {
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

			go ai.GetNextMove(board, cpuMove)
			marker = 'x'
		} else {
			pos = <-cpuMove
			marker = 'o'
		}

		s, err := board.MarkThisPosition(pos, marker)

		if err != nil {
			continue
		}

		ShowBoard(board.GetBoard())

		if s != "" {
			println(s)

			if playersTurn {
				println("Player wins!")
			} else {
				println("CPU wins!")
			}
		}

		playersTurn = !playersTurn
	}
}

func gameLoopTwoPlayers(board *ng.Board) {
	var pos int
	playerOnesTurn := true

	marker := 'x'

	ShowBoard(board.GetBoard())
	for !board.HasTheGameFinished() {
		fmt.Print("Make a move (0-8)")
		fmt.Scanf("%d", &pos)

		if playerOnesTurn {
			marker = 'x'
		} else {
			marker = 'o'
		}

		s, err := board.MarkThisPosition(pos, marker)
		if err != nil {
			continue
		}

		ShowBoard(board.GetBoard())
		if s != "" {
			println(s)

			if s != "draw" {
				if playerOnesTurn {
					println("Player one wins!")
				} else {
					println("Player two wins!")
				}
			}
		}

		playerOnesTurn = !playerOnesTurn

	}
}

func clearStdin() {
	var c *exec.Cmd

	cmd := true
	switch runtime.GOOS {
	case "linux":
		c = exec.Command("clear")
	case "windows":
		c = exec.Command("cmd", "/c", "cls")
	default:
		println()
		cmd = false
	}

	if cmd {
		c.Stdout = os.Stdout
		c.Run()
	}
}

func ShowBoard(b [9]rune) {
	clearStdin()

	println("\t\t\t\tFor reference:")

	for i := 0; i < 7; i += 3 {
		print("\t")

		for j := i; j < i+3; j++ {
			fmt.Printf(" %c ", b[j])

			if j != i+2 {
				print("|")
			}
		}

		fmt.Printf("\t|\t\t %d | %d | %d", i, i+1, i+2)

		if i < 6 {
			println("\n\t-----------\t|\t\t-----------")
		} else {
			println()
		}
	}
}
