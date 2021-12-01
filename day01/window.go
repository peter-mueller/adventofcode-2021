package day01

type Window struct {
	values []Measurement
	pos    int
}

func NewWindow(len int) *Window {
	values := make([]Measurement, len)
	for i := range values {
		values[i] = NoMeasurement
	}
	return &Window{
		values: values,
	}
}

func (w *Window) Sum() Measurement {
	sum := Measurement(0)
	for _, m := range w.values {
		if m == NoMeasurement {
			return NoMeasurement
		}
		sum += m
	}
	return sum
}

func (w *Window) Put(m Measurement) {
	w.values[w.pos] = m
	w.pos = (w.pos + 1) % len(w.values)
}
