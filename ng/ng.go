package ng

import (
	"fmt"
)

var (
	vertical       rule = rule{"vertical", 0, 1, 3}
	horizontal     rule = rule{"horizontal", 0, 3, 1}
	diagonal_left  rule = rule{"diagonal", 0, -1, 4}
	diagonal_right rule = rule{"diagonal", 2, -1, 2}
)

type rule struct {
	name      string
	start_pos int
	next      int
	increment int
}

type Board struct {
	finished bool
	draw     bool
	content  [9]rune
	empties  [9]bool
}

func (b *Board) New() *Board {
	b = new(Board)

	b.finished = false
	b.draw = false
	for i := range b.content {
		b.SetPositionContent(i, ' ', true)
		b.empties[i] = true
	}

	return b
}

func (b *Board) GetPositionContent(pos int) rune {
	if b == nil || pos > len(b.content) || pos < 0 {
		return 'n'
	}

	return b.content[pos]
}

func (b *Board) SetPositionContent(pos int, marker rune, overwrite bool) error {
	if !overwrite && !b.IsThisPositionEmpty(pos) {
		return fmt.Errorf("'%d' is not empty", pos)
	}

	b.content[pos] = marker

	if !overwrite {
		b.empties[pos] = false
	}

	return nil
}

func (b *Board) IsThisPositionEmpty(pos int) bool {
	return b.GetPositionContent(pos) == ' '
}

func (b *Board) getSequence(r rule, marker rune, res chan []int) {
	pos := []int{}

	if r.name == diagonal_left.name {
		for j := r.start_pos; j <= r.start_pos+(2*r.increment); j += r.increment {
			if b.GetPositionContent(j) == marker {
				pos = append(pos, j)
			}
		}
	} else {
		for i := r.start_pos; i <= r.start_pos+(2*r.next); i += r.next {
			for j := i; j <= i+(2*r.increment); j += r.increment {
				if b.GetPositionContent(j) == marker {
					pos = append(pos, j)
				}
			}

			if len(pos) < 2 {
				pos = []int{}
			}
		}
	}

	if len(pos) < 2 {
		pos = []int{}
	}

	res <- pos
}

func (b *Board) verifyRule(r rule, res chan bool) {
	if r.name == diagonal_left.name {
		markers := []rune{}
		for j := r.start_pos; j <= r.start_pos+(2*r.increment); j += r.increment {
			if b.IsThisPositionEmpty(j) {
				res <- false
			}

			markers = append(markers, b.GetPositionContent(j))
		}

		if markers[0] == markers[1] && markers[1] == markers[2] {
			res <- true
		}
	} else {
	OUTER:
		for i := r.start_pos; i <= r.start_pos+(2*r.next); i += r.next {
			markers := []rune{}
			for j := i; j <= i+(2*r.increment); j += r.increment {
				if b.IsThisPositionEmpty(j) {
					continue OUTER
				}

				markers = append(markers, b.GetPositionContent(j))
			}

			if markers[0] == markers[1] && markers[1] == markers[2] {
				res <- true
			}
		}
	}

	res <- false
}

func (b *Board) VerifyWinningConditions() string {

	rules := [4]rule{diagonal_left, diagonal_right, vertical, horizontal}
	res := [4]chan bool{make(chan bool), make(chan bool), make(chan bool), make(chan bool)}

	out := ""
	for i, r := range rules {
		go b.verifyRule(r, res[i])
	}

	for i, r := range res {
		c := <-r
		if c && out == "" {
			out = rules[i].name
		}
	}

	return out
}

func (b *Board) MarkThisPosition(pos int, marker rune) (string, error) {
	if !b.IsThisPositionEmpty(pos) {
		return "", fmt.Errorf("position %d is not empty", pos)
	}

	err := b.SetPositionContent(pos, marker, false)
	if err != nil {
		return "", err
	}

	res := b.VerifyWinningConditions()

	if res != "" {
		b.finished = true
	} else if len(b.GetEmptyPositions()) == 0 {
		b.draw = true
		b.finished = true
		res = "draw"
	}

	return res, nil
}

func (b *Board) HasTheGameFinished() bool {
	return b.finished
}

func (b *Board) GetAdjs(pos int) []int {
	adjs := [][]int{
		{1, 3, 4},
		{0, 2, 3, 4, 5},
		{1, 4, 5},
		{0, 1, 4, 6, 7},
		{0, 1, 2, 3, 5, 6, 7, 8},
		{1, 2, 4, 7, 8},
		{3, 4, 7},
		{3, 4, 5, 6, 8},
		{4, 5, 7},
	}

	return adjs[pos]
}

func (b *Board) GetEmptyAdjPositions(pos int) []int {
	empty := []int{}
	adjs := b.GetAdjs(pos)

	for _, a := range adjs {
		if b.IsThisPositionEmpty(a) {
			empty = append(empty, a)
		}
	}

	return empty
}

func (b *Board) GetBestPosition(marker rune) int {
	seqs := []chan []int{make(chan []int), make(chan []int), make(chan []int), make(chan []int)}
	rules := []rule{horizontal, vertical, diagonal_left, diagonal_right}

	for i, r := range rules {
		go b.getSequence(r, marker, seqs[i])
	}

	useThis := -1
	var current []int
	for i, s := range seqs {
		current = <-s

		useThis = i
		if len(current) > 0 {
			for _, pos := range current {
				if b.GetPositionContent(pos) != marker {
					useThis = -1
				}
			}
		}
	}

	if useThis < 0 {
		return -1
	}

	if len(current) > 0 {
		first, second := b.GetAdjs(current[0]), b.GetAdjs(current[1])

		for _, pos := range second {
			if b.IsThisPositionEmpty(pos) {
				return pos
			}
		}

		for _, pos := range first {
			if b.IsThisPositionEmpty(pos) {
				return pos
			}
		}
	}

	return -1
}

func (b *Board) GetEmptyPositions() []int {
	pos := []int{}

	for i, empty := range b.empties {
		if empty {
			pos = append(pos, i)
		}
	}

	return pos
}

func (b *Board) GetBoard() [9]rune {
	return b.content
}
