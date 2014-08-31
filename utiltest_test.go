// Copyright 2014 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package ut

import (
	"strconv"
	"testing"
)

func ExampleAssertEqual() {
	// For a func TestXXX(t *testing.T)
	t := &testing.T{}
	AssertEqual(t, "10", strconv.Itoa(10))
}

func ExampleAssertEqualIndex() {
	// For a func TestXXX(t *testing.T)
	t := &testing.T{}

	data := []struct {
		in       int
		expected string
	}{
		{9, "9"},
		{11, "10"},
	}
	for i, item := range data {
		// Call a function to test.
		actual := strconv.Itoa(item.in)
		// Then do an assert as a one-liner.
		AssertEqualIndex(t, i, item.expected, actual)
	}
}

func TestAssertEqual(t *testing.T) {
	t.Parallel()
	j := true
	var i interface{} = &j
	AssertEqual(t, &j, i)
	if t.Failed() {
		t.Fatal("Expected success")
	}
}

func TestAssertEqualFail(t *testing.T) {
	t.Parallel()
	t2 := &testing.T{}
	wait := make(chan bool)
	go func() {
		defer func() {
			recover()
			wait <- true
		}()
		AssertEqual(t2, true, false)
		t.Fail()
	}()
	<-wait
}

func TestAssertEqualIndex(t *testing.T) {
	t.Parallel()
	j := true
	var i interface{} = &j
	AssertEqualIndex(t, 24, &j, i)
	if t.Failed() {
		t.Fatal("Expected success")
	}
}

func TestAssertEqualIndexFail(t *testing.T) {
	t.Parallel()
	t2 := &testing.T{}
	wait := make(chan bool)
	go func() {
		defer func() {
			recover()
			wait <- true
		}()
		AssertEqualIndex(t2, 24, true, false)
		t.Fail()
	}()
	<-wait
}

func TestAssertEqualf(t *testing.T) {
	t.Parallel()
	j := true
	var i interface{} = &j
	AssertEqualf(t, &j, i, "foo %s %d", "bar", 2)
	if t.Failed() {
		t.Fatal("Expected success")
	}
}

func TestAssertEqualfFail(t *testing.T) {
	t.Parallel()
	t2 := &testing.T{}
	wait := make(chan bool)
	go func() {
		defer func() {
			recover()
			wait <- true
		}()
		AssertEqualf(t2, true, false, "foo %s %d", "bar", 2)
		t.Fail()
	}()
	<-wait
}