package cpu

import (
	"github.com/leo-alvarenga/tic-tac-go/ng"
)

func GetNextMove(b *ng.Board, position chan int) {
	position <- 0
}
