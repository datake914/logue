package domain

// LogStream represents the log stream resource.
type LogStream struct {
	ID             string
	Name           string
	EventCreatedAt int64
	EventUpdatedAt int64
}

// LogStreams represents the log stream resources.
type LogStreams struct {
	LogStreams []LogStream
}
