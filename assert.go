/*
Simple assertion package
*/
package assert

import "testing"

// General assert on condition
func Assert(condition bool, message string) {
	if !condition {
		panic("assert failed: " + message)
	}
}

// Assert fot testing
func AssertT(t *testing.T, condition bool, message string) {
	if !condition {
		t.Error(message)
	}
}
