package service

import (
	"github.com/datake914/logue/src/common/util"
	"github.com/datake914/logue/src/domain"
	"github.com/datake914/logue/src/service/adapter"
)

// LogEventService represents the struct that execute log event services.
type LogEventService struct {
	*BaseService
	LogEventRepo adapter.LogEventRepository
}

// NewLogEventService returns a pointer to LogEventService struct.
func NewLogEventService(logEventRepo adapter.LogEventRepository, opts ...Option) *LogEventService {
	return &LogEventService{
		BaseService:  NewBaseService(opts...),
		LogEventRepo: logEventRepo,
	}
}

// SearchLogEventRequest represents the request parameters that search log event resources.
type SearchLogEventRequest struct {
	LogGroupName  string  `json:"log_group_name" query:"log_group_name" validate:"required"`
	LogStreamName string  `json:"log_stream_name" query:"log_stream_name" validate:"required"`
	Limit         *int64  `json:"limit" query:"limit"`
	NextToken     *string `json:"next_token" query:"next_token"`
	StartFromHead *bool   `json:"start_from_head" query:"start_from_head"`
	From          *int64  `json:"from" query:"from"`
	To            *int64  `json:"to" query:"to"`
	FilterPattern *string `json:"filter_pattern" query:"filter_pattern"`
}

// SearchLogEventResponse represents the response parameters that search log event resources.
type SearchLogEventResponse struct {
	LogEvents         []LogEventModel `json:"log_events"`
	NextBackwardToken string          `json:"next_backward_token"`
	NextForwardToken  string          `json:"next_forward_token"`
}

// LogEventModel represents the common log event parameter
type LogEventModel struct {
	Message   string `json:"message"`
	CreatedAt int64  `json:"created_at"`
}

// Search returns log event resources.
func (s *LogEventService) Search(req SearchLogEventRequest) (res SearchLogEventResponse, err error) {
	// Convert service request to repository request.
	option := domain.LogEventSearchOption{}
	util.Copy(&req, &option)
	// Execute.
	var out domain.LogEvents
	if req.FilterPattern != nil {
		out, err = s.LogEventRepo.Search(req.LogGroupName, req.LogStreamName, option)
	} else {
		out, err = s.LogEventRepo.SearchWithFilter(req.LogGroupName, req.LogStreamName, option)
	}
	if err != nil {
		return res, err
	}
	// Convert repository response to service response.
	util.Copy(&out, &res)
	return res, nil
}
