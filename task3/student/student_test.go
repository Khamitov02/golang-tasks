package student

import (
	"errors"
	"testing"
)

func TestSetAge(t *testing.T) {
	tests := []struct {
		name    string
		age     int
		wantAge int
		wantErr error
	}{
		{"Set age to valid value", 25, 25, nil},
		{"Set age to 0", 0, 0, nil},
		{"Set age to negative value", -5, 0, ErrInvalidAge},
		{"Set age to unrealistic high value", 130, 0, ErrInvalidAge},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			student := &Student{"Alex", 22, 3.2}
			err := student.SetAge(tt.age)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("SetAge() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && student.Age != tt.wantAge {
				t.Errorf("SetAge() student.age = %v, wantAge %v", student.Age, tt.wantAge)
			}
		})
	}
}

func TestDoubleGPA(t *testing.T) {
	tests := []struct {
		name    string
		gpa     float64
		wantGPA float64
		wantErr error
	}{
		{"Double valid GPA", 1.6, 3.2, nil},
		{"Double zero GPA", 0, 0, nil},
		{"Double negative GPA", -2.0, 0, ErrInvalidGPA},
		{"Double GPA over 4.0", 2.5, 0, ErrInvalidGPA},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			student := &Student{"Alex", 22, tt.gpa}
			err := student.DoubleGPA()
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("DoubleGPA() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && student.Gpa != tt.wantGPA {
				t.Errorf("DoubleGPA() student.gpa = %v, wantGPA %v", student.Gpa, tt.wantGPA)
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			student := Student{tt.studentName, 22, 3.2}
			got := student.GetName()
			if got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
