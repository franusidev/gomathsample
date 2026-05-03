package gomathsample

import "testing"

func Test_Add(t *testing.T) {
	result := Add[int](4,1)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
	result2 := Add[float64](4.5,6.2)
	if result2 != 10.7 {
		t.Error("incorrect result: expected 10.7, got", result2)
	}
}
