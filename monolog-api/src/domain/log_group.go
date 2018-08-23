package domain

// LogGroup represents the log group resource.
type LogGroup struct {
	ID        string
	Name      string
	CreatedAt int64
}

// LogGroups represents the log group resources.
type LogGroups struct {
	LogGroups []LogGroup
}
