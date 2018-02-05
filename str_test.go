package aca
import ("testing")

func Test(t *testing.T){
	var out string
	out = check_ip("172.24.23.17")
	if out != "valid"{
		t.Error("incorrect output ",out)
	}
	
}
