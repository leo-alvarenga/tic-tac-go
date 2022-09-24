package main

import (
	"github.com/leo-alvarenga/tic-tac-go/io"
	"github.com/leo-alvarenga/tic-tac-go/ng"
)

func main() {
	var board *ng.Board
	board = board.New()

	io.CLI(board)
}
