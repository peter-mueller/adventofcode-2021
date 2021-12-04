package day03

import (
	"bufio"
	_ "embed"
	"log"
	"strconv"
	"strings"
)

type DiagnosticReport interface {
	Read() <-chan DiagnosticNumber
}

type DiagnosticNumber string

func (n DiagnosticNumber) Uint() uint {
	value, err := strconv.ParseUint(string(n), 2, 32)
	if err != nil {
		log.Fatalf("failed to parse diagnosticnumber as int: %s", n)
	}
	return uint(value)
}

func NewDiagnosticReport() DiagnosticReport {
	return &mockDiagnosticRepost{}
}

type mockDiagnosticRepost struct{}

//go:embed input.txt
var input string

func (mockDiagnosticRepost) Read() <-chan DiagnosticNumber {
	c := make(chan DiagnosticNumber, 1000)
	go func() {
		defer close(c)
		scanner := bufio.NewScanner(strings.NewReader(input))
		for scanner.Scan() {
			text := scanner.Text()
			c <- DiagnosticNumber(text)
		}
	}()
	return c
}
