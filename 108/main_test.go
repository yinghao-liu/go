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
