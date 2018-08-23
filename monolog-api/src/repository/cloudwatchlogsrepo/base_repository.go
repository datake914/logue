package cloudwatchlogsrepo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

var sess = session.Must(session.NewSession())

// BaseRepository represents the struct that access the CloudWatchLogs.
type BaseRepository struct {
	client *cloudwatchlogs.CloudWatchLogs
}

// NewBaseRepository returns a pointer to the BaseRepository struct.
func NewBaseRepository(opts ...Option) *BaseRepository {
	repo := &BaseRepository{
		client: cloudwatchlogs.New(sess, new(aws.Config)),
	}
	for _, opt := range opts {
		opt(repo)
	}
	return repo
}
