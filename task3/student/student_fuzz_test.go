package student

import (
	"errors"
	"testing"
)

func FuzzSetAge(f *testing.F) {
	f.Add(10)
	f.Add(-1)  // Negative age
	f.Add(130) // Unrealistic high age
	f.Fuzz(func(t *testing.T, age int) {
		student := &Student{"Alex", 22, 3.2}
		err := student.SetAge(age)

		if age < 0 || age > 120 {
			if !errors.Is(err, ErrInvalidAge) {
				t.Errorf("Expected ErrInvalidAge for age %d, got %v", age, err)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for age %d: %v", age, err)
			}
			if student.Age != age {
				t.Errorf("Expected age = %v, got %v", age, student.Age)
			}
		}
	})
}

func FuzzDoubleGPA(f *testing.F) {
	f.Add(1.75)
	f.Add(-0.5) // Negative GPA
	f.Add(2.5)  // Doubling would exceed max GPA
	f.Fuzz(func(t *testing.T, gpa float64) {
		student := &Student{"Alex", 22, gpa}
		err := student.DoubleGPA()
		expectedGPA := gpa * 2

		if expectedGPA < 0.0 || expectedGPA > 4.0 {
			if !errors.Is(err, ErrInvalidGPA) {
				t.Errorf("Expected ErrInvalidGPA for GPA %v, got %v", gpa, err)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for GPA %v: %v", gpa, err)
			}
			if student.Gpa != expectedGPA {
				t.Errorf("Expected GPA = %v, got %v", expectedGPA, student.Gpa)
			}
		}
	})
}
