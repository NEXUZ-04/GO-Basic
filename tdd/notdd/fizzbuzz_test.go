package main

import (
	"testing"
)

func TestFizzBuzz1(t *testing.T) {
	r := fizzbuzz(1)

	if r != "1" {
		t.Error("Expect 1 but receive ", r)
	}
}

func TestFizzBuzz2(t *testing.T) {
	r := fizzbuzz(2)

	if r != "2" {
		t.Error("Expect 2 but receive ", r)
	}
}

func TestFizzBuzz3(t *testing.T) {
	r := fizzbuzz(3)

	if r != "Fizz" {
		t.Error("Expect Fizz but receive ", r)
	}
}

func TestFizzBuzz4(t *testing.T) {
	r := fizzbuzz(4)

	if r != "4" {
		t.Error("Expect 4 but receive ", r)
	}
}

func TestFizzBuzz5(t *testing.T) {
	r := fizzbuzz(5)

	if r != "Buzz" {
		t.Error("Expect buzz but receive ", r)
	}
}

func TestFizzBuzz6(t *testing.T) {
	r := fizzbuzz(6)

	if r != "Fizz" {
		t.Error("Expect Fizz but receive ", r)
	}
}

func TestFizzBuzz7(t *testing.T) {
	r := fizzbuzz(7)

	if r != "7" {
		t.Error("Expect 7 but receive ", r)
	}
}

func TestFizzBuzz8(t *testing.T) {
	r := fizzbuzz(8)

	if r != "8" {
		t.Error("Expect 8 but receive ", r)
	}
}

func TestFizzBuzz9(t *testing.T) {
	r := fizzbuzz(9)

	if r != "Fizz" {
		t.Error("Expect Fizz but receive ", r)
	}
}

func TestFizzBuzz10(t *testing.T) {
	r := fizzbuzz(10)

	if r != "Buzz" {
		t.Error("Expect Buzz but receive ", r)
	}
}

func TestFizzBuzz11(t *testing.T) {
	r := fizzbuzz(11)

	if r != "11" {
		t.Error("Expect 11 but receive ", r)
	}
}

func TestFizzBuzz12(t *testing.T) {
	r := fizzbuzz(12)

	if r != "Fizz" {
		t.Error("Expect Fizz but receive ", r)
	}
}

func TestFizzBuzz13(t *testing.T) {
	r := fizzbuzz(13)

	if r != "13" {
		t.Error("Expect 13 but receive ", r)
	}
}

func TestFizzBuzz14(t *testing.T) {
	r := fizzbuzz(14)

	if r != "14" {
		t.Error("Expect 14 but receive ", r)
	}
}

func TestFizzBuzz15(t *testing.T) {
	r := fizzbuzz(15)

	if r != "FizzBuzz" {
		t.Error("Expect FizzBuzz but receive ", r)
	}
}
