package main

import "testing"

func TestSetAge(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want int
	}{
		{"Set age to 25", 25, 25},
		{"Set age to 0", 0, 0},
		{"Set age to negative value", -5, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			student := &Student{"Alex", 22, 3.2}
			student.setAge(tt.age)
			if student.age != tt.want {
				t.Errorf("setAge() = %v, want %v", student.age, tt.want)
			}
		})
	}
}

func TestDoubleGpa(t *testing.T) {
	tests := []struct {
		name string
		gpa  float64
		want float64
	}{
		{"Double positive GPA", 3.2, 6.4},
		{"Double zero GPA", 0, 0},
		{"Double negative GPA", -2.0, -4.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			student := &Student{"Alex", 22, tt.gpa}
			student.doubleGpa()
			if student.gpa != tt.want {
				t.Errorf("doubleGpa() = %v, want %v", student.gpa, tt.want)
			}
		})
	}
}

func TestGetName(t *testing.T) {
	tests := []struct {
		name        string
		studentName string
		want        string
	}{
		{"Get valid name", "Alex", "Alex"},
		{"Get empty name", "", ""},
		{"Wrong name expectation", "Alex", "John"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			student := Student{tt.studentName, 22, 3.2}
			got := student.getName()
			if got != tt.want {
				t.Errorf("getName() = %v, want %v", got, tt.want)
			}
		})
	}
}
