package pointers

import "testing"

func TestDoubleArray(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	DoubleArray(&arr)
	expected := [5]int{2, 4, 6, 8, 10}
	if arr != expected {
		t.Errorf("DoubleArray failed, expected %v, got %v", expected, arr)
	}
}

func TestAppendToSlice(t *testing.T) {
	slice := []int{1, 2, 3}
	AppendToSlice(&slice, 4)
	expected := []int{1, 2, 3, 4}
	if len(slice) != len(expected) || slice[3] != 4 {
		t.Errorf("AppendToSlice failed, expected %v, got %v", expected, slice)
	}
}

func TestResetSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	ResetSlice(slice)
	expected := []int{0, 0, 0, 0}
	for i := range slice {
		if slice[i] != expected[i] {
			t.Errorf("ResetSlice failed, expected %v, got %v", expected, slice)
			break
		}
	}
}

func TestSwapMapValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	SwapMapValues(m, "a", "b")
	if m["a"] != 2 || m["b"] != 1 {
		t.Errorf("SwapMapValues failed, got %v", m)
	}
}

func TestRemoveKeyFromMap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	RemoveKeyFromMap(m, "a")
	if _, exists := m["a"]; exists {
		t.Errorf("RemoveKeyFromMap failed, key 'a' still exists")
	}
}

func TestSumSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	result := SumSlice(slice)
	expected := 15
	if result != expected {
		t.Errorf("SumSlice failed, expected %d, got %d", expected, result)
	}
}

func TestReverseArray(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	ReverseArray(&arr)
	expected := [5]int{5, 4, 3, 2, 1}
	if arr != expected {
		t.Errorf("ReverseArray failed, expected %v, got %v", expected, arr)
	}
}

func TestAreSlicesEqual(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	if !AreSlicesEqual(slice1, slice2) {
		t.Errorf("AreSlicesEqual failed, slices should be equal")
	}
}

func TestMaxInSlice(t *testing.T) {
	slice := []int{1, 3, 2, 5, 4}
	result := MaxInSlice(slice)
	expected := 5
	if result == nil || *result != expected {
		t.Errorf("MaxInSlice failed, expected %d, got %v", expected, result)
	}
}

func TestIncrementMapValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	IncrementMapValues(m, 2)
	expected := map[string]int{"a": 3, "b": 4, "c": 5}
	for k, v := range m {
		if v != expected[k] {
			t.Errorf("IncrementMapValues failed for key %s, expected %d, got %d", k, expected[k], v)
		}
	}
}
