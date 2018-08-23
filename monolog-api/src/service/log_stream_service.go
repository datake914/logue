package service

import (
	"github.com/datake914/logue/src/common/util"
	"github.com/datake914/logue/src/service/adapter"
)

// LogStreamService represents the struct that execute log stream services.
type LogStreamService struct {
	*BaseService
	LogStreamRepo adapter.LogStreamRepository
}

// NewLogStreamService returns a pointer to LogStreamService struct.
func NewLogStreamService(logStreamRepo adapter.LogStreamRepository, opts ...Option) *LogStreamService {
	return &LogStreamService{
		BaseService:   NewBaseService(opts...),
		LogStreamRepo: logStreamRepo,
	}
}

// SearchLogStreamRequest represents the request parameters that search log stream resources.
type SearchLogStreamRequest struct {
	LogGroupName string `json:"log_group_name" query:"log_group_name" validate:"required"`
}

// SearchLogStreamResponse represents the response parameters that search log stream resources.
type SearchLogStreamResponse struct {
	LogStreams []LogStreamModel `json:"log_groups"`
}

// SearchLogStreamModel represents the common log stream parameter
type LogStreamModel struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
}

// Search returns log stream resources.
func (s *LogStreamService) Search(req SearchLogStreamRequest) (res SearchLogStreamResponse, err error) {
	// Execute.
	out, err := s.LogStreamRepo.Search(req.LogGroupName)
	if err != nil {
		return res, err
	}
	// Convert repository response to service response.
	util.Copy(&out, &res)
	return res, nil
}
