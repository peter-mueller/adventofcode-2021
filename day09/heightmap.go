package day09

type Heightmap struct {
	data   map[Location]Height
	bounds Rectangle
}

func (hm *Heightmap) Bounds() Rectangle {
	return hm.bounds
}

func (hm *Heightmap) At(l Location) (h Height, ok bool) {
	if !hm.bounds.Contains(l) {
		return 0, false
	}
	return hm.data[l], true
}

type Location struct{ X, Y int }

func (l Location) West() Location {
	l.X -= 1
	return l
}
func (l Location) East() Location {
	l.X += 1
	return l
}
func (l Location) North() Location {
	l.Y -= 1
	return l
}
func (l Location) South() Location {
	l.Y += 1
	return l
}

type Rectangle struct{ Min, Max Location }

func (r Rectangle) Contains(l Location) bool {
	return inbounds(l.X, r.Min.X, r.Max.X) &&
		inbounds(l.Y, r.Min.Y, r.Max.Y)
}

func inbounds(val, min, max int) bool {
	return min <= val && val <= max
}

type Height int
