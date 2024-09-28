package pointers

import "testing"

func TestSwap(t *testing.T) {
	a, b := 5, 10
	Swap(&a, &b)
	if a != 10 || b != 5 {
		t.Errorf("Swap failed, got a=%d, b=%d", a, b)
	}
}

func TestIncrementBy(t *testing.T) {
	val := 7
	IncrementBy(&val, 3)
	if val != 10 {
		t.Errorf("IncrementBy failed, expected 10, got %d", val)
	}
}

func TestResetToZero(t *testing.T) {
	val := 20
	ResetToZero(&val)
	if val != 0 {
		t.Errorf("ResetToZero failed, expected 0, got %d", val)
	}
}

func TestNewIntPointer(t *testing.T) {
	ptr := NewIntPointer(50)
	if ptr == nil || *ptr != 50 {
		t.Errorf("NewIntPointer failed, expected 50, got %v", ptr)
	}
}

func TestMultiply(t *testing.T) {
	val := 3
	Multiply(&val, 4)
	if val != 12 {
		t.Errorf("Multiply failed, expected 12, got %d", val)
	}
}

func TestGetValue(t *testing.T) {
	val := 42
	result := GetValue(&val)
	if result != 42 {
		t.Errorf("GetValue failed, expected 42, got %d", result)
	}
}

func TestSetFloatValue(t *testing.T) {
	var f float64 = 3.14
	SetFloatValue(&f, 1.23)
	if f != 1.23 {
		t.Errorf("SetFloatValue failed, expected 1.23, got %f", f)
	}
}

func TestAreSameAddress(t *testing.T) {
	x, y := 100, 100
	if !AreSameAddress(&x, &x) {
		t.Errorf("AreSameAddress failed, expected true, got false")
	}
	if AreSameAddress(&x, &y) {
		t.Errorf("AreSameAddress failed, expected false, got true")
	}
}

func TestChangeStringValue(t *testing.T) {
	str := "old"
	ChangeStringValue(&str, "new")
	if str != "new" {
		t.Errorf("ChangeStringValue failed, expected 'new', got '%s'", str)
	}
}

func TestNewZeroFloatPointer(t *testing.T) {
	ptr := NewZeroFloatPointer()
	if ptr == nil || *ptr != 0.0 {
		t.Errorf("NewZeroFloatPointer failed, expected 0.0, got %v", ptr)
	}
}
