package singleton

import "testing"

func TestGetInstance(t *testing.T) {

	firstCounter := GetInstance()

	if firstCounter	== nil {
		t.Error("Expected pointer to singleton after calling GetInstance(), not nil")
	}

	expectedCounter := firstCounter

	count := firstCounter.AddOne()

	if count != 1 {
		t.Error("After calling for the first time to count, the count must be 1 but it is $d\n", count)
	}

	secondCounter := GetInstance()
	if secondCounter != expectedCounter {
		t.Error("Expected same counter in secondCounter but got different instance")
	}

	count = secondCounter.AddOne()
	if count != 2 {
		t.Error("After calling AddOne for the second counter, count must be 2 but was $d\n", count)
	}

}