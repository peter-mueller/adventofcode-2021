package puzzle

type Puzzle interface {
	QuestionPrinter
	AnswerPrinter
}

type QuestionPrinter interface {
	PrintQuestion()
}

type AnswerPrinter interface {
	PrintAnswer()
}
