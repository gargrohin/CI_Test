package util

import "testing"

func TestIpcheck(t *testing.T) {
	if ipcheck("000.0000.00.00") != 0 {
		t.Errorf("Expected 0")
	} else if ipcheck("192.168.1.1") != 1 {
		t.Errorf("Expected 1")
	} else if ipcheck("912.456..") != 0 {
		t.Errorf("Expected 0")
	}
}

func TestAlphanumeric(t *Testing.T) {
	if alphanumeric("asd123asd") != "asd123asd" {
		t.Errorf("shouldn't have changed")
	} else if alphanumeric("asd!@#123./><asd") != "asd123asd" {
		t.Errorf("Not changing")
	}
}
