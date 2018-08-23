package service

import (
	"github.com/datake914/logue/src/common/util"
	"github.com/datake914/logue/src/service/adapter"
)

// LogGroupService represents the struct that execute log group services.
type LogGroupService struct {
	*BaseService
	LogGroupRepo adapter.LogGroupRepository
}

// NewLogGroupService returns a pointer to LogGroupService struct.
func NewLogGroupService(logGroupRepo adapter.LogGroupRepository, opts ...Option) *LogGroupService {
	return &LogGroupService{
		BaseService:  NewBaseService(opts...),
		LogGroupRepo: logGroupRepo,
	}
}

// SearchLogGroupRequest represents the request parameters that search log group resources.
type SearchLogGroupRequest struct {
}

// SearchLogGroupResponse represents the response parameters that search log group resources.
type SearchLogGroupResponse struct {
	LogGroups []LogGroupModel `json:"log_groups"`
}

// SearchLogGroupModel represents the common log group parameter
type LogGroupModel struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
}

// Search returns log group resources.
func (s *LogGroupService) Search(req SearchLogGroupRequest) (res SearchLogGroupResponse, err error) {
	// Execute.
	out, err := s.LogGroupRepo.Search()
	if err != nil {
		return res, err
	}
	// Convert repository response to service response.
	util.Copy(&out, &res)
	return res, nil
}
