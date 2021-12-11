package day09

type FilterKernel func(l Location, hm *Heightmap) bool

func (k FilterKernel) Apply(hm *Heightmap) (filtered []Location) {

	b := hm.Bounds()

	for x := b.Min.X; x <= b.Max.X; x++ {
		for y := b.Min.Y; y <= b.Max.Y; y++ {

			l := Location{X: x, Y: y}
			ok := k(l, hm)
			if ok {
				filtered = append(filtered, l)
			}

		}
	}
	return filtered
}

type SelectKernel func(l Location, hm *Heightmap) []Location

func (kernel SelectKernel) SelectFrom(start Location, hm *Heightmap) (selected []Location) {
	var (
		visited = map[Location]bool{start: true}
		tovisit = []Location{start}
	)
	next := func() Location {
		next := tovisit[0]
		visited[next] = true
		tovisit = tovisit[1:]
		return next
	}

	for len(tovisit) != 0 {
		selected := kernel(next(), hm)
		for _, newL := range selected {
			if visited[newL] {
				continue
			}
			tovisit = append(tovisit, newL)
		}
	}

	for l := range visited {
		selected = append(selected, l)
	}
	return selected
}
