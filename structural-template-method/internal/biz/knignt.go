package biz

type knight struct {
	base    Role
	defense int
}

func NewKnight() Role {
	warrior := NewWarrior()
	return &knight{
		base: warrior,
	}
}

func (r *knight) Attack(role Role, state RoleState) {
	r.base.Attack(role, state)
}

func (r *knight) Defend(damage int) {
	damage = damage - r.defense
	if damage < 0 {
		damage = -damage
	}
	r.ReceiveDamage(damage)
}

func (r *knight) ReceiveDamage(damage int) {
	r.base.ReceiveDamage(damage)
}

func (r *knight) SetHp(value int) Role {
	r.base.SetHp(value)
	return r
}

func (r *knight) SetPower(value int) Role {
	r.base.SetPower(value)
	return r
}
func (r *knight) SetDefense(value int) Role {
	r.defense = value
	return r
}

func (r *knight) SetMana(value int) Role {
	return r
}

func (r *knight) Hp() int {
	return r.base.Hp()
}

func (r *knight) power() int {
	return r.base.power()
}
