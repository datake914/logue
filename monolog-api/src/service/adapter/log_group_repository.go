package adapter

import "github.com/datake914/logue/src/domain"

// LogGroupRepository is the interface that access the log group resource.
type LogGroupRepository interface {
	Search() (domain.LogGroups, error)
}
