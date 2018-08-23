package service

import (
	"reflect"
	"testing"

	"github.com/datake914/logue/src/repository/cloudwatchlogsrepo"
)

/**
 * Test of Log Stream Service Search.
 */
func TestSearchLogStream(t *testing.T) {
	// Setup & Teardown
	defer setupTeardownLogStream(t)()
	// TestCases
	testcases := []struct {
		name   string
		input  SearchLogStreamRequest
		output SearchLogStreamResponse
		err    error
	}{
		{
			"Normal#1",
			SearchLogStreamRequest{
				LogGroupName: "/aws/lambda/AwsServerlessExpressFunction",
			},
			SearchLogStreamResponse{
				LogStreams: []LogStreamModel{},
			},
			nil,
		},
	}
	// Create service.
	s := createLogStreamService()
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

// createLogStreamService creates log stream service.
func createLogStreamService() *LogStreamService {
	return NewLogStreamService(
		cloudwatchlogsrepo.NewLogStreamRepository(
			cloudwatchlogsrepo.WithEndpoint(""),
		),
	)
}

func setupTeardownLogStream(t *testing.T) func() {
	// Setup

	// Teardown
	return func() {
	}
}
