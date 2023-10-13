package dragonball

type Dragonball struct {
	star int
}

func (d *Dragonball) Star() int {
	return d.star
}
