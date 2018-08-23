package service

import (
	"reflect"
	"testing"

	"github.com/datake914/logue/src/repository/cloudwatchlogsrepo"
)

/**
 * Test of Log Group Service Search.
 */
func TestSearchLogEvent(t *testing.T) {
	// Setup & Teardown
	defer setupTeardownLogEvent(t)()
	// TestCases
	testcases := []struct {
		name   string
		input  SearchLogEventRequest
		output SearchLogEventResponse
		err    error
	}{
		{
			"Normal#1",
			SearchLogEventRequest{
				LogGroupName:  "/aws/lambda/AwsServerlessExpressFunction",
				LogStreamName: "2016/10/18/[$LATEST]e4c529f8778b4813a3430c55d346e134",
			},
			SearchLogEventResponse{
				LogEvents: []LogEventModel{},
			},
			nil,
		},
	}
	// Create service.
	s := createLogEventService()
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

// createLogEventService creates log event service.
func createLogEventService() *LogEventService {
	return NewLogEventService(
		cloudwatchlogsrepo.NewLogEventRepository(
			cloudwatchlogsrepo.WithEndpoint(""),
		),
	)
}

func setupTeardownLogEvent(t *testing.T) func() {
	// Setup

	// Teardown
	return func() {
	}
}
