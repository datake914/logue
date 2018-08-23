package cloudwatchlogsrepo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/datake914/logue/src/domain"
)

// LogStreamRepository represents the struct that access the log stream resource.
type LogStreamRepository struct {
	*BaseRepository
}

// NewLogStreamRepository returns a pointer to the LogStreamRepository struct.
func NewLogStreamRepository(opts ...Option) *LogStreamRepository {
	return &LogStreamRepository{
		BaseRepository: NewBaseRepository(opts...),
	}
}

// Search returns log group resources.
func (repo *LogStreamRepository) Search(name string) (res domain.LogStreams, err error) {
	in := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: aws.String(name),
	}
	err = repo.client.DescribeLogStreamsPages(in, func(out *cloudwatchlogs.DescribeLogStreamsOutput, isLast bool) bool {
		for _, v := range out.LogStreams {
			res.LogStreams = append(res.LogStreams, domain.LogStream{
				ID:             aws.StringValue(v.Arn),
				Name:           aws.StringValue(v.LogStreamName),
				EventCreatedAt: aws.Int64Value(v.FirstEventTimestamp),
				EventUpdatedAt: aws.Int64Value(v.LastEventTimestamp),
			})
		}
		return isLast
	})
	return res, err
}
