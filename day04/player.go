package day04

type result struct {
	board Board
	score int
	moves int
}

type player struct {
	board    Board
	moves    int
	finished bool
}

func (p *player) play(input <-chan int) chan result {
	c := make(chan result)
	go func() {
		defer close(c)
		for n := range input {
			if p.finished {
				continue
			}
			p.moves += 1
			marked := p.board.Mark(n)
			if !marked {
				continue
			}
			if p.board.IsFinished() {
				p.finished = true
				c <- result{
					board: p.board,
					score: p.board.Score(n),
					moves: p.moves,
				}
			}
		}
	}()
	return c
}
