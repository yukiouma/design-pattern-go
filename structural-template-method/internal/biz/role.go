package biz

type Role interface {
	Attack(role Role, state RoleState)
	Defend(damage int)
	ReceiveDamage(damage int)
	SetHp(value int) Role
	SetPower(value int) Role
	SetDefense(value int) Role
	SetMana(value int) Role
	Hp() int
	power() int
}
