package biz

type warrior struct {
	hp int
	p  int
}

func NewWarrior() Role {
	return &warrior{}
}

func (r *warrior) Attack(role Role, state RoleState) {
	switch state {
	case Attack:
		role.ReceiveDamage(r.p)
		role.Attack(r, coolDown)
	case Defend:
		role.Defend(r.p)
	default:
		role.ReceiveDamage(r.p)
	}
}

func (r *warrior) Defend(damage int) {
	r.ReceiveDamage(damage)
}

func (r *warrior) ReceiveDamage(damage int) {
	r.hp -= damage
}

func (r *warrior) SetHp(value int) Role {
	r.hp = value
	return r
}

func (r *warrior) SetPower(value int) Role {
	r.p = value
	return r
}
func (r *warrior) SetDefense(value int) Role {
	return r
}

func (r *warrior) SetMana(value int) Role {
	return r
}

func (r *warrior) Hp() int {
	return r.hp
}

func (r *warrior) power() int {
	return r.p
}
