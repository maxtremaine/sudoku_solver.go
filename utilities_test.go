package main

import "testing"

func TestLastIndex(t *testing.T) {
	indexIsFive := lastIndex([]uint8{ 1, 2, 3, 3, 3, 3 }, 3)
	if indexIsFive != 5 {
		t.Errorf("lastIndex Expected 5, received %d.", indexIsFive)
	}
}

func TestHasDuplicates(t *testing.T) {
	includesDuplicates := []uint8{ 1, 2, 3, 3 }
	shouldBeTrue := hasDuplicates(includesDuplicates)
	if shouldBeTrue != true {
		t.Errorf("hasDuplicates expected true, received %t.", shouldBeTrue)
	}
	allUnique := []uint8{ 1, 2, 3, 4 }
	shouldBeFalse := hasDuplicates(allUnique)
	if shouldBeFalse != false {
		t.Errorf("hasDuplicates expected false, received %t.", shouldBeFalse)
	}
}