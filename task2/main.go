package main

import "fmt"

type Student struct {
	name string
	age  int
	gpa  float64
}

func (s *Student) setAge(age int) {
	s.age = age
	fmt.Printf("Age changed to value: %v\n", s.age)
}

func (s *Student) doubleGpa() {
	s.gpa = s.gpa * 2.0
	fmt.Printf("Gpa was doubled\n")
}

func (s Student) getName() string {
	return s.name
}

func main() {
	//Struct with methods example
	var student1 = &Student{"Erik", 20, 3.5}
	fmt.Println("Student's default values: ", student1)

	student1.setAge(5)
	student1.doubleGpa()
	fmt.Println("Student's values after change: ", student1)
	fmt.Printf("Student's name is: %s\n", student1.getName())

	//Map example
	box := make(map[int]string)
	box[1] = "gifts"
	box[4] = "rocks"
	box[77] = "balls"
	fmt.Printf("Map contains: %v\n", box)

	box[4] = "changed_rocks"
	fmt.Printf("Map after content updated: %v\n", box)

	if _, ok := box[99]; ok {
		fmt.Println("Box 99 does exist")
	} else {
		fmt.Println("Box 99 doesn't exist")
	}

	delete(box, 1)
	fmt.Printf("Map after content deleted: %v\n", box)

	// Array example
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array content:")
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	// Slice example
	slice := []int{10, 20, 30, 40, 50, 60}
	slice = append(slice, 70)
	fmt.Println("Slice elements:")
	for i, val := range slice {
		if i%2 == 0 {
			fmt.Printf("Value %d a even index %d\n", val, i)
		} else {
			fmt.Printf("Value %d at odd index %d\n", val, i)
		}
	}

	// Additional loop and condition
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Printf("Index %d is even\n", i)
		} else {
			fmt.Printf("Index %d is odd\n", i)
		}
	}
}
