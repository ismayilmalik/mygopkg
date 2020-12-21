package mygopkg

import "testing"

func TestAdd(t *testing.T)  {
	expected := 10
	num1 := 3
	num2 := 7

	if result := Add(num1, num2); result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

}