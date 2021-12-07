package day07

type CrabSubmarine struct {
	HorizontalPosition int
}

func (c *CrabSubmarine) DistanceTo(horizontalPosition int) int {
	distance := c.HorizontalPosition - horizontalPosition
	if distance < 0 {
		return -distance
	}
	return distance
}

func (c *CrabSubmarine) FuelTo(horizontalPos int) int {
	return c.DistanceTo(horizontalPos)
}

func (c *CrabSubmarine) FuelVariantTo(horizontalPos int) int {
	return triangularNumber(c.DistanceTo(horizontalPos))
}

// ex.:
// .____ 1
// ..___ 3
// ...__ 6
// ...._ 10
// (half 4 * 5 Rectangle)
func triangularNumber(n int) int {
	return (n * (n + 1)) / 2
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
