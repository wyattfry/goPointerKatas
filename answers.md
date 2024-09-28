# Possible Solutions


```go
package pointers

// Function that swaps two integers using pointers
func Swap(a, b *int) {
	*a, *b = *b, *a
}

// Function that increments an integer by a given value using a pointer
func IncrementBy(p *int, val int) {
	*p += val
}

// Function that sets the value of an integer to zero using a pointer
func ResetToZero(p *int) {
	*p = 0
}

// Function that creates a new pointer to an integer with a given value
func NewIntPointer(val int) *int {
	return &val
}

// Function that multiplies the value of an integer by another value using a pointer
func Multiply(p *int, factor int) {
	*p *= factor
}

// Function that returns the dereferenced value of an integer pointer
func GetValue(p *int) int {
	return *p
}

// Function that sets a float to a new value using a pointer
func SetFloatValue(p *float64, val float64) {
	*p = val
}

// Function that checks if two integer pointers point to the same address
func AreSameAddress(a, b *int) bool {
	return a == b
}

// Function that changes a string pointer to point to a new string
func ChangeStringValue(p *string, newVal string) {
	*p = newVal
}

// Function that returns a pointer to a float64 initialized to zero
func NewZeroFloatPointer() *float64 {
	var zero float64 = 0.0
	return &zero
}
```