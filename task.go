package main

type Gun struct {
	On    bool
	Ammo  int
	Power int
}

func (g *Gun) Shoot() bool {
	if !g.On {
		return false
	}

	if g.Ammo > 0 {
		g.Ammo--
		return true
	} else {
		return false
	}
}

func (g *Gun) RideBike() bool {
	if !g.On {
		return false
	}

	if g.Power > 0 {
		g.Power--
		return true
	} else {
		return true
	}
}

type A struct {
	a int
}

type B struct {
	A // По умолчанию указывается тип A в переменную A
	b int
}


