package testutils

import (
	"errors"
	"reflect"
	"testing"
)

func AssertErrorEquals(test *testing.T, expected error, actual error) {
	test.Helper()

	if !errors.Is(actual, expected) {
		test.Error("expected error:", expected, "got:", actual)
	}
}

func AssertNoError(test *testing.T, actual error) {
	test.Helper()

	if !isNil(actual) {
		test.Error("Expected no error, got:", actual)
	}
}

func AssertError(test *testing.T, actual error) {
	test.Helper()

	if isNil(actual) {
		test.Error("Expected (any) error, got nil")
	}
}

func AssertNotNil(test *testing.T, actual any, name string) {
	test.Helper()

	if isNil(actual) {
		test.Error("Expected", name, "got nil")
	}
}

func AssertNil(test *testing.T, actual any, name string) {
	test.Helper()

	if !isNil(actual) {
		test.Error("Expected nil", name, "got:", actual)
	}
}

func AssertTrue(test *testing.T, actual bool) {
	test.Helper()

	if !actual {
		test.Error("Expected true, got false")
	}
}

func AssertFalse(test *testing.T, actual bool) {
	test.Helper()

	if actual {
		test.Error("Expected false, got true")
	}
}

// AssertEquals compares if two VALUES are the same for value types. Also works for pointers to check if the pointer location is the same.
// WARNING: Will fail when trying to compare functions
func AssertEquals(test *testing.T, expected any, actual any) {
	test.Helper()

	if !reflect.DeepEqual(expected, actual) {
		test.Errorf("Expected %+v, but got %+v", expected, actual)
	}
}

func AssertPanics(test *testing.T, expectedToPanic func()) {
	test.Helper()

	defer func() {
		_error := recover()
		if _error == nil {
			test.Error("Expected panic, didn't happen")
		}
	}()

	expectedToPanic()
}

func AssertIn[T any](t *testing.T, expectedIn []T, actual T) {
	t.Helper()

	gotExpected := false
	for _, expected := range expectedIn {
		if reflect.DeepEqual(expected, actual) {
			gotExpected = true
			break
		}
	}

	if !gotExpected {
		t.Errorf("Expected in %+v, but got %+v", expectedIn, actual)
	}
}

func AssertEmpty(test *testing.T, actual any) {
	test.Helper()

	if !isEmpty(actual) {
		test.Error("Expected empty, got", actual)
	}
}

func AssertNotEmpty(test *testing.T, actual any) {
	test.Helper()

	if isEmpty(actual) {
		test.Error("Expected not empty, got empty")
	}
}

func isEmpty(value any) bool {
	_value := reflect.ValueOf(value)

	switch _value.Kind() {
	case reflect.String, reflect.Slice, reflect.Map, reflect.Array, reflect.Chan:
		return _value.Len() == 0
	default:
		return isNil(value)
	}
}

func isNil(value any) bool {
	if value == nil {
		return true
	}

	// Use reflection to handle cases where value is an interface or pointer
	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return val.IsNil()
	}

	return false
}
