/*
Simple assertion package
*/
package assert

import (
	"fmt"
	"testing"
)

type msg interface {
	String() string
}

// General assert on condition
func Assert(condition bool, message any) {
	if condition {
		return
	}

	s, ok := message.(string)
	if ok {
		panic("assert failed: " + s)
	}

	e, ok := message.(error)
	if ok {
		panic("assert failed: " + e.Error())
	}

	m, ok := message.(msg)
	if ok {
		panic("assert failed: " + m.String())
	}

	fmt.Println("assert failed unable to decode message paramter panic with raw value")
	panic(message)

}

// Assert fot testing
func AssertT(t *testing.T, condition bool, message any) {
	if condition {
		return
	}

	t.Error(message)
}
