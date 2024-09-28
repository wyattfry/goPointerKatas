package pointers

// 1. Structs: Update the age of a person using a pointer
type Person struct {
	Name string
	Age  int
}

func UpdateAge(p *Person, newAge int) {
	// Incorrect implementation (does nothing)
}

// 2. Pointer Receivers in Methods: Counter struct with methods to increment and reset the value
type Counter struct {
	Value int
}

func (c *Counter) Increment() {
	// Incorrect implementation (does nothing)
}

func (c *Counter) Reset() {
	// Incorrect implementation (does nothing)
}

// 3. Nil Pointers as Return Values: Find a person by name or return nil if not found
func FindPersonByName(name string) *Person {
	// Broken implementation (always returns nil)
	return nil
}

// 4. Passing Large Data Structures: Process a large data structure efficiently with pointers
type LargeData struct {
	Data [1000]int
}

func ProcessLargeData(data *LargeData) {
	// Incorrect implementation (does nothing)
}

// 5. Embedded Pointers in Structs: Car struct with a pointer to an Engine, upgrade the engine's horsepower
type Engine struct {
	Horsepower int
}

type Car struct {
	Make   string
	Model  string
	Engine *Engine
}

func UpgradeEngine(c *Car, hp int) {
	// Incorrect implementation (does nothing)
}

// 6. Interface and Pointers: Implement an interface for a Book that can be described
type Describable interface {
	Describe() string
}

type Book struct {
	Title string
}

func (b *Book) Describe() string {
	// Broken implementation (returns an empty string)
	return ""
}
