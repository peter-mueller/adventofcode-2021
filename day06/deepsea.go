package day06

type DeepSea struct {
	fishs map[BirthInDuration]LanternfishCount
}

func NewDeepSea() DeepSea {
	return DeepSea{
		fishs: make(map[BirthInDuration]LanternfishCount),
	}
}

func (d *DeepSea) AddFish(l Lanternfish) {
	d.fishs[BirthInDuration(l.BirthIn)] += 1
}

func (d *DeepSea) TotalCount() int {
	total := 0
	for _, count := range d.fishs {
		total += int(count)
	}
	return total
}

type BirthInDuration Duration
type LanternfishCount int

func (s *DeepSea) SimulateDuration(n Duration) {
	for i := Duration(0); i < n; i++ {
		s.SimulateDay()
	}
}

func (s *DeepSea) SimulateDay() {
	nextfishs := make(map[BirthInDuration]LanternfishCount)

	// parent
	nextfishs[6] = s.fishs[0]
	// newborn
	nextfishs[8] = s.fishs[0]
	delete(s.fishs, 0)

	for birthInDuration, count := range s.fishs {
		nextfishs[birthInDuration-1] += count
	}

	s.fishs = nextfishs
}
