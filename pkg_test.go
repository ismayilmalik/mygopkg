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

func TestMultiply(t *testing.T)  {
	expected := 25
	num1 := 5
	num2 := 5

	if result := Multiply(num1, num2); result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

}