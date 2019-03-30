package main

import (
	"fmt"
)

func main() {
	a := (93 == 93)
	b := (93 <= 93)
	c := (93 >= 93)
	d := (93 != 93)
	e := (93 < 93)
	f := (93 > 93)

	fmt.Println(a, b, c, d, e, f)
}
