package biz

type RoleState int

const (
	Attack RoleState = iota
	Defend
	coolDown
)
