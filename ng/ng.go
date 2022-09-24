package ng

import "fmt"

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
}

func (b *Board) New() *Board {
	b = new(Board)

	b.finished = false
	b.draw = false
	for i := range b.content {
		b.content[i] = ' '
	}

	return b
}

func (b *Board) IsThisPositionEmpty(pos int) bool {

	if b == nil || pos > len(b.content) || pos < 0 {
		return false
	}

	return b.content[pos] == ' '
}

func (b *Board) verifyRule(r rule, res chan bool) {
	if r.next <= 0 {
		markers := []rune{}
		for j := r.start_pos; j <= r.start_pos+(2*r.increment); j += r.increment {
			if b.IsThisPositionEmpty(j) {
				res <- false
			}

			markers = append(markers, b.content[j])
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

				markers = append(markers, b.content[j])
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

	b.content[pos] = marker
	res := b.VerifyWinningConditions()

	if res != "" {
		b.finished = true
	}

	return res, nil
}

func (b *Board) HasTheGameFinished() bool {
	return b.finished
}

func (b *Board) GetBoard() [9]rune {
	return b.content
}
