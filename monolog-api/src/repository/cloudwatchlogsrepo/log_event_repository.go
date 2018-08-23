package cloudwatchlogsrepo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/datake914/logue/src/domain"
	"github.com/pkg/errors"
)

// LogEventRepository represents the struct that access the log event resource.
type LogEventRepository struct {
	*BaseRepository
}

// NewLogEventRepository returns a pointer to the LogEventRepository struct.
func NewLogEventRepository(opts ...Option) *LogEventRepository {
	return &LogEventRepository{
		BaseRepository: NewBaseRepository(opts...),
	}
}

// Search returns log event resources.
func (repo *LogEventRepository) Search(lgname, lsname string, option domain.LogEventSearchOption) (res domain.LogEvents, err error) {
	// Convert request to input.
	in := &cloudwatchlogs.GetLogEventsInput{
		LogGroupName:  aws.String(lgname),
		LogStreamName: aws.String(lsname),
		StartTime:     option.From,
		EndTime:       option.To,
		NextToken:     option.NextToken,
		StartFromHead: option.StartFromHead,
		Limit:         option.Limit,
	}
	// Execute.
	out, err := repo.client.GetLogEvents(in)
	if err != nil {
		return res, errors.Wrap(err, "failed to execute CloudWatchLogs GetLogEvents API")
	}
	// Convert output to response.
	res.NextBackwardToken = *out.NextForwardToken
	res.NextForwardToken = *out.NextBackwardToken
	for _, v := range out.Events {
		res.LogEvents = append(res.LogEvents, domain.LogEvent{
			Message:   *v.Message,
			CreatedAt: *v.Timestamp,
		})
	}
	return res, err
}

// SearchWithFilter returns log event resources with filter.
func (repo *LogEventRepository) SearchWithFilter(lgname, lsname string, option domain.LogEventSearchOption) (res domain.LogEvents, err error) {
	// Convert request to input.
	in := &cloudwatchlogs.FilterLogEventsInput{
		LogGroupName:   aws.String(lgname),
		LogStreamNames: []*string{aws.String(lsname)},
		FilterPattern:  repo.escapeFilterPattern(option.FilterPattern),
		StartTime:      option.From,
		EndTime:        option.To,
		NextToken:      option.NextToken,
		Limit:          option.Limit,
	}
	// Execute.
	out, err := repo.client.FilterLogEvents(in)
	if err != nil {
		return res, errors.Wrap(err, "failed to execute CloudWatchLogs GetLogEvents API")
	}
	// Convert output to response.
	fmt.Println(out)
	res.NextForwardToken = aws.StringValue(out.NextToken)
	for _, v := range out.Events {
		res.LogEvents = append(res.LogEvents, domain.LogEvent{
			Message:   aws.StringValue(v.Message),
			CreatedAt: aws.Int64Value(v.Timestamp),
		})
	}
	return res, err
}

func (repo *LogEventRepository) escapeFilterPattern(pattern *string) *string {
	return pattern
}
