package day03

import (
	"log"
)

type LifeSupportRatingCalc struct {
	report       []DiagnosticNumber
	numberLength int
}

func NewLifeSupportRatingCalc(input <-chan DiagnosticNumber) *LifeSupportRatingCalc {
	c := &LifeSupportRatingCalc{
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

func (c *LifeSupportRatingCalc) LifeSupportRating() uint {
	numbersForOxygen := c.report
	for pos := 0; pos < c.numberLength; pos++ {
		zeros, ones := count(numbersForOxygen, pos)
		switch {
		case zeros > ones:
			numbersForOxygen = filter(numbersForOxygen, pos, '0')
		case ones > zeros:
			numbersForOxygen = filter(numbersForOxygen, pos, '1')
		case ones == zeros:
			numbersForOxygen = filter(numbersForOxygen, pos, '1')
		}
		if len(numbersForOxygen) <= 1 {
			break
		}
	}
	if len(numbersForOxygen) != 1 {
		log.Fatalf("error oxygen: %+v", numbersForOxygen)
	}
	oxygen := numbersForOxygen[0]

	numbersForCO2 := c.report
	for pos := 0; pos < c.numberLength; pos++ {
		zeros, ones := count(numbersForCO2, pos)
		switch {
		case zeros > ones:
			numbersForCO2 = filter(numbersForCO2, pos, '1')
		case ones > zeros:
			numbersForCO2 = filter(numbersForCO2, pos, '0')
		case ones == zeros:
			numbersForCO2 = filter(numbersForCO2, pos, '0')
		}
		if len(numbersForCO2) <= 1 {
			break
		}
	}
	if len(numbersForCO2) != 1 {
		log.Fatalf("error co2: %+v", numbersForCO2)
	}
	co2 := numbersForCO2[0]

	return oxygen.Uint() * co2.Uint()
}

func filter(input []DiagnosticNumber, pos int, r rune) (output []DiagnosticNumber) {
	for _, n := range input {
		if n[pos] == byte(r) {
			output = append(output, n)
		}
	}
	return output
}

func count(input []DiagnosticNumber, pos int) (zeros, ones int) {
	for _, n := range input {
		r := n[pos]
		switch r {
		case '0':
			zeros += 1
		case '1':
			ones += 1
		}
	}
	return zeros, ones
}
