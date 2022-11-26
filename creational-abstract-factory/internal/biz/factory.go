package biz

type Factory interface {
	CreateGeneral(name string) Genenal
	CreateDirector(member Genenal) (Director, error)
}
