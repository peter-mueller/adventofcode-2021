package day08

import (
	"log"
	"sort"
)

type SegmentPatterns [10]Pattern

type Pattern []rune

func (p Pattern) String() string {
	return string(p)
}

func (p Pattern) Normalize() (o Pattern) {
	slice := p[:]
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice
}

func Solve(line NotebookLine) (result SegmentPatterns) {
	patterns := line.Patterns[:]

	//  aaaa
	// b    c
	// b    c
	//  dddd
	// e    f
	// e    f
	//  gggg

	// easy
	result[1] = must(find(patterns, hasLen(2)))
	result[4] = must(find(patterns, hasLen(4)))
	result[7] = must(find(patterns, hasLen(3)))
	result[8] = must(find(patterns, hasLen(7)))
	// harder
	result[3] = must(find(patterns, hasLen(5), hasNSimilarWith(result[7], 3)))
	result[9] = must(find(patterns, hasLen(6), hasNSimilarWith(result[3], 5)))
	result[6] = must(find(patterns, hasLen(6), hasNSimilarWith(result[1], 1)))
	result[5] = must(find(patterns, hasLen(5), hasNSimilarWith(result[6], 5)))
	result[0] = must(find(patterns, hasLen(6), isNot(result[9]), isNot(result[6])))
	result[2] = must(find(patterns, hasLen(5), isNot(result[3]), isNot(result[5])))

	return result
}

type predicate func(p Pattern) bool

func must(found Pattern, ok bool) Pattern {
	if !ok {
		log.Fatal("expected a result pattern")
	}
	return found
}

func find(patterns []Pattern, p ...predicate) (found Pattern, ok bool) {
nextpattern:
	for _, pattern := range patterns {
		for _, predicate := range p {
			if !predicate(pattern) {
				continue nextpattern
			}
		}
		return pattern, true
	}
	return nil, false
}

func hasLen(n int) predicate {
	return func(p Pattern) bool {
		return len(p) == n
	}
}

func isNot(a Pattern) predicate {
	return func(p Pattern) bool {
		return p.String() != a.String()
	}
}

func hasNSimilarWith(a Pattern, nSimilar int) predicate {
	return func(p Pattern) bool {
		count := 0
		for _, ra := range p {
			for _, rb := range a {
				if ra == rb {
					count++
				}
			}
		}
		return count == nSimilar
	}
}
