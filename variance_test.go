package variance

import "testing"

func TestVariance(t *testing.T) {
	var acc Accumulator

	acc.Add(0)
	acc.Add(1)
	acc.Add(2)
	if exp, got := 1., acc.Mean(); got != exp {
		t.Error("expect mean", exp, "got", got)
	}
	if exp, got := 1., acc.Variance(); got != exp {
		t.Error("expect variance", exp, "got", got)
	}
	acc.Del(0)
	if exp, got := 1.5, acc.Mean(); got != exp {
		t.Error("expect mean", exp, "got", got)
	}
	if exp, got := 0.5, acc.Variance(); got != exp {
		t.Error("expect variance", exp, "got", got)
	}
}
