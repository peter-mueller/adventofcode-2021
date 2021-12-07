package day07

type CrabSubmarine struct {
	HorizontalPosition int
}

type ByPos []CrabSubmarine

func (a ByPos) Len() int {
	return len(a)
}
func (s ByPos) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByPos) Less(i, j int) bool {
	return s[i].HorizontalPosition < s[j].HorizontalPosition
}
