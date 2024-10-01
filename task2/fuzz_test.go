package main

import "testing"

func FuzzSetAge(f *testing.F) {
	f.Add(10)
	f.Add(-1) // Add a failing case for negative ages
	f.Fuzz(func(t *testing.T, age int) {
		student := &Student{"Alex", 22, 3.2}
		student.setAge(age)

		if student.age != age {
			t.Errorf("Expected age = %v, got %v", age, student.age)
		}

		// Wrong behavior check
		if age < 0 && student.age == age {
			t.Errorf("Negative age should not be allowed, got %v", student.age)
		}
	})
}

func FuzzDoubleGpa(f *testing.F) {
	f.Add(3.5)
	f.Add(-3.0) // Add a case for negative GPA
	f.Fuzz(func(t *testing.T, gpa float64) {
		student := &Student{"Alex", 22, gpa}
		student.doubleGpa()
		expectedGpa := gpa * 2

		if student.gpa != expectedGpa {
			t.Errorf("Expected GPA = %v, got %v", expectedGpa, student.gpa)
		}

		// Wrong behavior check
		if gpa < 0 && student.gpa == expectedGpa {
			t.Errorf("Negative GPA should not be doubled, got %v", student.gpa)
		}
	})
}
