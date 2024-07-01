package terrestrial

type Terrestrial struct {
	Radius float64
	Mass   float64
}

func (t *Terrestrial) CalculateGravity() float64 {
	return t.Mass / (t.Radius * t.Radius)
}
