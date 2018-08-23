package service

import (
	"reflect"
	"testing"

	"github.com/datake914/logue/src/repository/cloudwatchlogsrepo"
)

/**
 * Test of Log Group Service Search.
 */
func TestSearchLogGroup(t *testing.T) {
	// Setup & Teardown
	defer setupTeardownLogGroup(t)()
	// TestCases
	testcases := []struct {
		name   string
		input  SearchLogGroupRequest
		output SearchLogGroupResponse
		err    error
	}{
		{
			"Normal#1",
			SearchLogGroupRequest{},
			SearchLogGroupResponse{
				LogGroups: []LogGroupModel{},
			},
			nil,
		},
	}
	// Create service.
	s := createLogGroupService()
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			// Execute service.
			actual, err := s.Search(testcase.input)
			if err != nil && testcase.err == nil {
				t.Fatalf("%+v", err)
			}
			// Assert error.
			if err != nil {
				assertError(t, err, testcase.err)
			}
			// Assert response.
			if !reflect.DeepEqual(&actual, &testcase.output) {
				t.Errorf("unexpected response returned\nactual: %v\nexpected: %v", actual, testcase.output)
			}
		})
	}
}

// createLogGroupService creates log group service.
func createLogGroupService() *LogGroupService {
	return NewLogGroupService(
		cloudwatchlogsrepo.NewLogGroupRepository(
			cloudwatchlogsrepo.WithEndpoint(""),
		),
	)
}

func setupTeardownLogGroup(t *testing.T) func() {
	// Setup

	// Teardown
	return func() {
	}
}
