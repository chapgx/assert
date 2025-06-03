/*
Simple assertion package
*/
package assert

import (
	"fmt"
	"testing"
)

type Message interface {
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

	m, ok := message.(Message)
	if ok {
		panic("assert failed: " + m.String())
	}

	fmt.Println("assert failed unable to decode message paramter panic with raw value")
	panic(message)

}

// Assert for testing
func AssertT(t *testing.T, condition bool, message any) {
	if condition {
		return
	}

	t.Error(message)
}

// Takes any sets of arguments where the last one must be of type error. If error is not nil it panics
func Must(args ...any) []any {
	if len(args) <= 0 {
		panic("No arguments passed to the function")
	}

	e := args[len(args)-1]

	if args[len(args)-1] == nil {
		return args[:len(args)-1]
	}

	evalue, ok := e.(error)
	if !ok {
		panic("last argument is not of error type")
	}

	if evalue == nil {
		return args[:len(args)-1]
	}

	panic(evalue)
}
