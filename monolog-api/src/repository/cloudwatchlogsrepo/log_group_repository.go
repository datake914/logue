package cloudwatchlogsrepo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/datake914/logue/src/domain"
)

// LogGroupRepository represents the struct that access the log group resource.
type LogGroupRepository struct {
	*BaseRepository
}

// NewLogGroupRepository returns a pointer to the LogGroupRepository struct.
func NewLogGroupRepository(opts ...Option) *LogGroupRepository {
	return &LogGroupRepository{
		BaseRepository: NewBaseRepository(opts...),
	}
}

// Search returns log group resources.
func (repo *LogGroupRepository) Search() (res domain.LogGroups, err error) {
	err = repo.client.DescribeLogGroupsPages(&cloudwatchlogs.DescribeLogGroupsInput{}, func(out *cloudwatchlogs.DescribeLogGroupsOutput, isLast bool) bool {
		for _, v := range out.LogGroups {
			res.LogGroups = append(res.LogGroups, domain.LogGroup{
				ID:        aws.StringValue(v.Arn),
				Name:      aws.StringValue(v.LogGroupName),
				CreatedAt: aws.Int64Value(v.CreationTime),
			})
		}
		return isLast
	})
	return res, err
}
