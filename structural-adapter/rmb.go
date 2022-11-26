package structuraladapter

type Rmb float64

func NewRmb(v float64) Rmb {
	return Rmb(v)
}

func (r Rmb) Denomination() float64 {
	return float64(r)
}
