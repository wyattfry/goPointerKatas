package pointers

// Function that swaps two integers using pointers
func Swap(a, b *int) {
	// Incorrect implementation (does nothing)
}

// Function that increments an integer by a given value using a pointer
func IncrementBy(p *int, val int) {
	// Incorrect implementation (does nothing)
}

// Function that sets the value of an integer to zero using a pointer
func ResetToZero(p *int) {
	// Incorrect implementation (does nothing)
}

// Function that creates a new pointer to an integer with a given value
func NewIntPointer(val int) *int {
	// Broken implementation (returns nil)
	return nil
}

// Function that multiplies the value of an integer by another value using a pointer
func Multiply(p *int, factor int) {
	// Incorrect implementation (does nothing)
}

// Function that returns the dereferenced value of an integer pointer
func GetValue(p *int) int {
	// Broken implementation (returns 0 regardless of pointer)
	return 0
}

// Function that sets a float to a new value using a pointer
func SetFloatValue(p *float64, val float64) {
	// Incorrect implementation (does nothing)
}

// Function that checks if two integer pointers point to the same address
func AreSameAddress(a, b *int) bool {
	// Broken implementation (always returns false)
	return false
}

// Function that changes a string pointer to point to a new string
func ChangeStringValue(p *string, newVal string) {
	// Incorrect implementation (does nothing)
}

// Function that returns a pointer to a float64 initialized to zero
func NewZeroFloatPointer() *float64 {
	// Broken implementation (returns nil)
	return nil
}
