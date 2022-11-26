package cor

type Role int

const (
	Reader Role = iota
	Appender
	Updater
	Deleter
)
