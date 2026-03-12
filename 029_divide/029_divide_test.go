package main

import "testing"

func TestExample1(t *testing.T) {
	if quotient := divide(10, 3); quotient != 3 {
		t.Errorf("Expect 10/3 = 3, got %d", quotient)
	}
}

func TestExample2(t *testing.T) {
	if quotient := divide(7, -3); quotient != -2 {
		t.Errorf("Expect 7/-3 = -2, got %d", quotient)
	}
}

func TestExample3(t *testing.T) {
	if quotient := divide(-2147483648, -1); quotient != 2147483647 {
		t.Errorf("Expect -2147483648/-1 = 2147483647, got %d", quotient)
	}
}

func TestExample4(t *testing.T) {
	if quotient := divide(-2147483648, 4); quotient != -536870911 {
		t.Errorf("Expect -2147483648/4 = -536870912, got %d", quotient)
	}
}

func TestExample5(t *testing.T) {
	if quotient := divide(1, 1); quotient != 1 {
		t.Errorf("Expect 1/1 = 1, got %d", quotient)
	}
}
