package cloudwatchlogsrepo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

// Option represents the function that set optional parameter.
type Option func(*BaseRepository) error

// WithEndpoint returns the function that set an endpoint parameter to BaseRepository.
func WithEndpoint(endpoint string) Option {
	return func(repo *BaseRepository) error {
		repo.client = cloudwatchlogs.New(sess, new(aws.Config).WithEndpoint(endpoint))
		return nil
	}
}
