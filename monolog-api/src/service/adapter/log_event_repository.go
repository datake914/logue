package adapter

import "github.com/datake914/logue/src/domain"

// LogEventRepository is the interface that access the log event resource.
type LogEventRepository interface {
	Search(lgname, lsname string, option domain.LogEventSearchOption) (domain.LogEvents, error)
	SearchWithFilter(lgname, lsname string, option domain.LogEventSearchOption) (domain.LogEvents, error)
}
