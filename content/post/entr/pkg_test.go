package main

import "testing"

func Test_Inc(t *testing.T) {
	actual := Inc(1)
	if actual != 2 {
		t.Errorf("Expecting: <2>, got <%d>", actual)
	}
}
