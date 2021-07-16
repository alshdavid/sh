package yargs

import (
	"fmt"
	"testing"
)

var FOO = "foo"
var BAR = "bar"
var FIZZ = "fizz"
var BUZZ = "buzz"
var VALUE_1 = "value_1"
var CASE_1 = fmt.Sprintf("--%s -%s --%s %s", FOO, BAR, FIZZ, VALUE_1)

func TestParse(t *testing.T) {
	params := Parse(VALUE_1)

	if !params.Has(FOO) {
		t.Fail()
	}
	if !params.Has(BAR) {
		t.Fail()
	}
	if !params.Has(FIZZ) {
		t.Fail()
	}
	if params.Has(BUZZ) {
		t.Fail()
	}
	if len(params.Values(FIZZ)) != 1 {
		t.Fail()
	}
}
