package biz

type magician struct {
	base Role
	mana int
}

func NewMagician() Role {
	warrior := NewWarrior()
	return &magician{
		base: warrior,
	}
}

func (r *magician) Attack(role Role, state RoleState) {
	originalPower := r.base.power()
	r.SetPower(originalPower + r.mana)
	r.base.Attack(role, state)
	r.SetPower(originalPower)
}

func (r *magician) Defend(damage int) {
	r.base.Defend(damage)
}

func (r *magician) ReceiveDamage(damage int) {
	r.base.ReceiveDamage(damage)
}

func (r *magician) SetHp(value int) Role {
	r.base.SetHp(value)
	return r
}

func (r *magician) SetPower(value int) Role {
	r.base.SetPower(value)
	return r
}
func (r *magician) SetDefense(value int) Role {
	return r
}

func (r *magician) SetMana(value int) Role {
	r.mana = value
	return r
}

func (r *magician) Hp() int {
	return r.base.Hp()
}

func (r *magician) power() int {
	return r.base.power()
}
