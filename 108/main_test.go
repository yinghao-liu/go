package main

import "testing"

// 	Field Greater Than Another Field Only valid for Numbers, time.Duration and time.Time types
type ValidateStruct struct {
	Start string
	End   string `validate:"gtfield=Start"`
	Numa  int
	Numb  int `validate:"gtfield=Numa"`
}

func TestValidate(t *testing.T) {
	var v ValidateStruct
	v.Start = "10:00"
	v.End = "10:01"
	v.Numa = 3
	v.Numb = 2
	Validate(v)
}

type Student struct {
	Name string `validate:"len=7"`
	Age  int    `validate:"min=18,max=30"`
}
type Worker struct {
	Name   string   `validate:"len=7"`
	Age    int      `validate:"min=25,max=30"`
	Skills []string `validate:"min=25,max=30"`
}

type Human struct {
	Profession string   `validate:"oneof=student worker"`
	Stud       *Student `validate:"required_if=Profession student"` // only for struct pointer, required tag works
	Work       *Worker  `validate:"required_if=Profession worker"`
}

func TestHuman(t *testing.T) {
	var h Human
	h.Profession = "student"
	var stu Student
	stu.Name = "francis"
	stu.Age = 1
	h.Stud = &stu
	Validate(h)
}
