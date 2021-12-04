package day03

type PowerCalc struct {
	report       []DiagnosticNumber
	numberLength int
}

func NewPowerCalc(input <-chan DiagnosticNumber) *PowerCalc {
	c := &PowerCalc{
		report: make([]DiagnosticNumber, 0),
	}
	for n := range input {
		c.report = append(c.report, n)
	}
	if len(c.report) > 0 {
		c.numberLength = len(c.report[0])
	}
	return c
}

func (c *PowerCalc) PowerConsumption() uint {
	var (
		gamma   = ""
		epsilon = ""
	)

	for pos := 0; pos < c.numberLength; pos++ {
		zeros, ones := count(c.report, pos)
		switch {
		case zeros > ones:
			gamma += "1"
			epsilon += "0"
		case ones > zeros:
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaInt := DiagnosticNumber(gamma).Uint()
	epsilonInt := DiagnosticNumber(epsilon).Uint()
	return gammaInt * epsilonInt
}
