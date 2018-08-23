package domain

// LogEvent represents the log event resource.
type LogEvent struct {
	Message   string
	CreatedAt int64
}

// LogEvents represents the log event resources.
type LogEvents struct {
	LogEvents         []LogEvent
	NextBackwardToken string
	NextForwardToken  string
}

// LogEventSearchOption represents the options for searching log event resources.
type LogEventSearchOption struct {
	Limit         *int64
	NextToken     *string
	StartFromHead *bool
	From          *int64
	To            *int64
	FilterPattern *string
}
