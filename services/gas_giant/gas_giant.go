package gas_giant

type GasGiant struct {
	Radius float64
}

func (g *GasGiant) CalculateGravity() float64 {
	return 0.5 / (g.Radius * g.Radius)
}
