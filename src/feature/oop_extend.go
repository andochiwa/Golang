package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

// Pupil 小学生继承Student
type Pupil struct {
	Student
}

// Graduate 大学生继承Student
type Graduate struct {
	Student
}

func (student *Student) test() {
	fmt.Println(student)
}

func main() {
	pupil := &Pupil{Student{"mary", 10, 3}}
	graduate := &Graduate{Student{"jary", 20, 5}}
	pupil.test()
	graduate.test()
}
