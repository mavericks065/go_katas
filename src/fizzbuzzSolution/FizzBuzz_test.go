package fizzbuzzSolution_test

import (
	. "fizzbuzzSolution"
	"testing"
)

func Test_Fizz(t *testing.T) {
	result := GetResult(3)

	if result != "fizz" {
		t.Errorf("Test_Fizz " + "3")
	} else {
		t.Log("success")
	}

	result1 := GetResult(9)

	if result1 != "fizz" {
		t.Errorf("Test_Fizz " + "9")
	} else {
		t.Log("success")
	}

	result3 := GetResult(33)

	if result3 != "fizz" {
		t.Errorf("Test_Fizz " + "33")
	} else {
		t.Log("success")
	}
}

func Test_Buzz(t *testing.T) {
	result := GetResult(5)

	if result != "buzz" {
		t.Errorf("Test_Fizz " + "5")
	} else {
		t.Log("success")
	}

	result1 := GetResult(50)

	if result1 != "buzz" {
		t.Errorf("Test_Fizz " + "50")
	} else {
		t.Log("success")
	}

	result2 := GetResult(95)

	if result2 != "buzz" {
		t.Errorf("Test_Fizz " + "95")
	} else {
		t.Log("success")
	}
}

func Test_FizzBuzz(t *testing.T) {

	result1 := GetResult(15)

	if result1 != "fizzbuzz" {
		t.Errorf("Test_Fizz " + "15")
	} else {
		t.Log("success")
	}

	result2 := GetResult(30)

	if result2 != "fizzbuzz" {
		t.Errorf("Test_Fizz " + "30")
	} else {
		t.Log("success")
	}
}

func Test_Other(t *testing.T) {

	result := GetResult(4)

	if result != "4" {
		t.Errorf("Test_Other value " + "4")
	} else {
		t.Log("success")
	}

	result2 := GetResult(8)

	if result2 != "8" {
		t.Errorf("Test_Other value " + "8")
	} else {
		t.Log("success")
	}
}
