package cor

type Handler interface {
	SetNext(Handler)
	Handle(string) bool
}
