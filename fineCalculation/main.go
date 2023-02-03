package main

import (
	"fmt"
)

func main() {
	var d1, m1, y1, d2, m2, y2 int
	fmt.Scanln(&d1, &m1, &y1)
	fmt.Scanln(&d2, &m2, &y2)

	var fine int
	if y2 > y1 {
		fine = 12000
	} else if m2 > m1 {
		fine = 500 * (m2 - m1)
	} else if d2 > d1 {
		fine = 15 * (d2 - d1)
	}

	fmt.Println(fine)
}
