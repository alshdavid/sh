package yargs

import (
	"testing"
)

func TestNewFlagCollection(t *testing.T) {
	flags := NewFlagCollection("-fb --foobar --foo test --bar value1  --bar value2")

	if flags.Has("f") == false {
		t.Error()
	}
	if flags.Has("b") == false {
		t.Error()
	}
	if flags.Has("foobar") == false {
		t.Error()
	}
	if flags.Has("foo") == false {
		t.Error()
	}
	if len(flags.Get("f")) != 1 || flags.Get("f")[0] != "true" {
		t.Error()
	}
	if len(flags.Get("b")) != 1 || flags.Get("b")[0] != "true" {
		t.Error()
	}
	if len(flags.Get("foobar")) != 1 || flags.Get("foobar")[0] != "true" {
		t.Error()
	}
	if len(flags.Get("foo")) != 1 || flags.Get("foo")[0] != "test" {
		t.Error()
	}
	if len(flags.Get("bar")) != 2 || flags.Get("bar")[0] != "value1" {
		t.Error()
	}
	if len(flags.Get("bar")) != 2 || flags.Get("bar")[1] != "value2" {
		t.Error()
	}
}
