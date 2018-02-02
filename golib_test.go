package main

import "testing"

func TestValidip(t *testing.T) {

	ret := Validip("str")
	if ret != 1 {
		t.Errorf("Correct check not done")
	}
}
