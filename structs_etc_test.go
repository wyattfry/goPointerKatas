package pointers

import "testing"

// 1. Structs
func TestUpdateAge(t *testing.T) {
	p := &Person{Name: "Alice", Age: 30}
	UpdateAge(p, 35)
	if p.Age != 35 {
		t.Errorf("UpdateAge failed, expected age 35, got %d", p.Age)
	}
}

// 2. Pointer Receivers in Methods
func TestCounter(t *testing.T) {
	c := &Counter{Value: 5}
	c.Increment()
	if c.Value != 6 {
		t.Errorf("Counter.Increment failed, expected 6, got %d", c.Value)
	}
	c.Reset()
	if c.Value != 0 {
		t.Errorf("Counter.Reset failed, expected 0, got %d", c.Value)
	}
}

// 3. Nil Pointers as Return Values
func TestFindPersonByName(t *testing.T) {
	person := FindPersonByName("Bob")
	if person == nil || person.Name != "Bob" {
		t.Errorf("FindPersonByName failed, expected person with name Bob, got %v", person)
	}

	person = FindPersonByName("")
	if person != nil {
		t.Errorf("FindPersonByName failed, expected nil, got %v", person)
	}
}

// 4. Passing Large Data Structures
func TestProcessLargeData(t *testing.T) {
	data := LargeData{}
	for i := 0; i < len(data.Data); i++ {
		data.Data[i] = i
	}
	ProcessLargeData(&data)
	// You can add specific assertions here if needed, but for now, we just check for proper processing.
}

// 5. Embedded Pointers in Structs
func TestUpgradeEngine(t *testing.T) {
	car := &Car{
		Make:  "Toyota",
		Model: "Corolla",
		Engine: &Engine{
			Horsepower: 130,
		},
	}
	UpgradeEngine(car, 200)
	if car.Engine.Horsepower != 200 {
		t.Errorf("UpgradeEngine failed, expected 200 horsepower, got %d", car.Engine.Horsepower)
	}
}

// 6. Interface and Pointers
func TestBookDescribe(t *testing.T) {
	b := &Book{Title: "The Go Programming Language"}
	desc := b.Describe()
	expected := "Book: The Go Programming Language"
	if desc != expected {
		t.Errorf("Book.Describe failed, expected '%s', got '%s'", expected, desc)
	}
}
