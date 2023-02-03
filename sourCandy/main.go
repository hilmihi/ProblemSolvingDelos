package main

import (
	"fmt"
)

func main() {
	var student, candies, first_student int
	fmt.Scanln(&student, &candies, &first_student)

	left_candy := candies % student
	last_student := (first_student - 1) + left_candy
	if last_student > student {
		last_student = last_student - student
	}

	fmt.Println(last_student)
}
