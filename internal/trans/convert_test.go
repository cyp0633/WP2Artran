package trans

import "testing"

func TestConvertTime(t *testing.T) {
	output:=convertTime("2022-12-28T08:30:29")
	t.Log(output)
	if output!="2022-12-28 08:30:29 +0000 +0000"{
		t.Error("convertTime failed")
	}
}