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
	fmt.Println("Student's name returned")
	return s.name
}

func main() {
	//custom struct
	var student1 = &Student{"Erik", 20, 3.5}
	fmt.Println("Student's default values: ", student1)

	student1.setAge(5)
	student1.doubleGpa()

	fmt.Println("Student's values after change: ", student1)
	fmt.Printf("Student's name is: %s\n", student1.getName())

	//map examples
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

}
