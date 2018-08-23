package service

import (
	"errors"
	"os"
	"reflect"
	"testing"
)

// TestMain is the entry point.
func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	cleanup()
	os.Exit(ret)
}

// setup is executed once before test executing.
func setup() {
}

// cleanup is executed once after test executing.
func cleanup() {
}

func assertError(t *testing.T, actual, expected error) {
	if expected != nil {
		// Skip unexpected exception.
		if reflect.TypeOf(errors.New("unexpected")) == reflect.TypeOf(expected) {
			return
		}
		if reflect.TypeOf(actual) != reflect.TypeOf(expected) {
			t.Errorf("unexpected error returned.\nactual:%v\nexpected: %v", reflect.TypeOf(actual), reflect.TypeOf(expected))
		}
	} else {
		t.Fatalf("%+v", actual)
	}
}
