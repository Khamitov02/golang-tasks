// Package student provides functionality to manage student data.
package student

import (
	"errors"
	"fmt"
)

// Student represents a student with a name, age, and GPA.
type Student struct {
	Name string
	Age  int
	Gpa  float64
}

// ErrInvalidAge is returned when an invalid age is set.
var ErrInvalidAge = errors.New("invalid age")

// ErrInvalidGPA is returned when an invalid GPA is encountered.
var ErrInvalidGPA = errors.New("invalid GPA")

// SetAge sets the student's age.
// Returns an error if the age is negative or unrealistic.
func (s *Student) SetAge(age int) error {
	if age < 0 || age > 120 {
		return fmt.Errorf("SetAge: %w: %d", ErrInvalidAge, age)
	}
	s.Age = age
	fmt.Printf("Age changed to value: %v\n", s.Age)
	return nil
}

// DoubleGPA doubles the student's GPA.
// Returns an error if the resulting GPA is invalid (e.g., less than 0.0 or greater than 4.0).
func (s *Student) DoubleGPA() error {
	newGPA := s.Gpa * 2.0
	if newGPA < 0.0 || newGPA > 4.0 {
		return fmt.Errorf("DoubleGPA: %w: %v", ErrInvalidGPA, newGPA)
	}
	s.Gpa = newGPA
	fmt.Printf("GPA was doubled\n")
	return nil
}

// GetName returns the student's name.
func (s Student) GetName() string {
	return s.Name
}
