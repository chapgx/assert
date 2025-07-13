/*
Simple assertion package
*/
package assert

import (
	"errors"
	"fmt"
	"testing"
)

type Message interface {
	String() string
}

// Assert general assert on condition
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

// AssertOrReturn assert the condition or returns the error
func AssertOrReturn(condition bool, message any) error {
	if condition {
		return nil
	}

	s, ok := message.(string)
	if ok {
		return errors.New(s)
	}

	e, ok := message.(error)
	if ok {
		return e
	}

	m, ok := message.(Message)
	if ok {
		return errors.New(m.String())
	}

	return fmt.Errorf("%s %+v", "unknow message type", message)
}

// AssertT for testing
func AssertT(t *testing.T, condition bool, message any) {
	if condition {
		return
	}

	t.Error(message)
}

// Must takes any sets of arguments where the last one must be of type error. If error is not nil it panics
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
