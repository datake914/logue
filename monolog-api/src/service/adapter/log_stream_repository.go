package adapter

import "github.com/datake914/logue/src/domain"

// LogStreamRepository is the interface that access the log stream resource.
type LogStreamRepository interface {
	Search(name string) (domain.LogStreams, error)
}
