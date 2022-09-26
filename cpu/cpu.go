package cpu

import (
	"math/rand"
	"time"

	"github.com/leo-alvarenga/tic-tac-go/ng"
)

type CPU interface {
	GetNextMove(b *ng.Board, position chan int)
}

type EasyCPU struct{}
type HardCPU struct{}

func (c *EasyCPU) GetNextMove(b *ng.Board, position chan int) {
	pos := b.GetBestPosition('o')

	if pos < 0 {
		rand.Seed(time.Now().UnixNano())
		pos = rand.Intn(9)

		for !b.IsThisPositionEmpty(pos) {
			pos = rand.Intn(9)
		}
	}

	position <- pos
}

func (c *HardCPU) GetNextMove(b *ng.Board, position chan int) {
	rand.Seed(time.Now().UnixNano())
	pos := rand.Intn(9)

	for !b.IsThisPositionEmpty(pos) {
		pos = rand.Intn(9)
	}

	position <- pos
}
