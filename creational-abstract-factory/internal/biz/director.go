package biz

type Director interface {
	Genenal
	AdjustSalary(member Genenal, salary float64) error
}
