package main

import (
	"testing"
	"reflect"
)

func TestLastIndex(t *testing.T) {
	indexIsFive := lastIndex([]int{ 1, 2, 3, 3, 3, 3 }, 3)
	if indexIsFive != 5 {
		t.Errorf("lastIndex Expected 5, received %d.", indexIsFive)
	}
}

func TestHasDuplicates(t *testing.T) {
	includesDuplicates := []int{ 1, 2, 3, 3 }
	shouldBeTrue := hasDuplicates(includesDuplicates)
	if shouldBeTrue != true {
		t.Errorf("hasDuplicates expected true, received %t.", shouldBeTrue)
	}
	allUnique := []int{ 1, 2, 3, 4 }
	shouldBeFalse := hasDuplicates(allUnique)
	if shouldBeFalse != false {
		t.Errorf("hasDuplicates expected false, received %t.", shouldBeFalse)
	}
}

func TestUniqueNumbers(t *testing.T) {
	testList := []int{ 4, 4, 5, 5, 5, 1, 2, 3 }
	outputList := []int{ 1, 2, 3, 4, 5 }
	uniqueOutput := uniqueNumbers(testList)
	if !reflect.DeepEqual(uniqueOutput, outputList) {
		t.Errorf("uniqueNumbers expected %d, received %d.", outputList, uniqueOutput)
	}
}