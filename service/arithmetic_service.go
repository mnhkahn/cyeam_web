package service

import (
	"fmt"
	"math/rand"
)

func GenArithmetic(t string, r *rand.Rand) string {
	var res = ""
	switch t {
	case "2-2-2-add-1":
		a := r.Int63n(79)
		b := r.Int63n(89 - a)
		c := r.Int63n(99 - a - b)
		res = fmt.Sprintf("%d + %d + %d =", a, b, c)
	case "2-2-2-sub-1":
		a := r.Int63n(79) + 20 // 20～99
		b := r.Int63n(a - 10)
		c := r.Int63n(a - b - 10)
		res = fmt.Sprintf("%d - %d - %d =", a, b, c)
	case "1-1-mul-1":
		a := r.Int63n(8) + 1 // 1~9
		b := r.Int63n(8) + 1 // 1~9
		res = fmt.Sprintf("%d × %d = ", a, b)
	case "1-1-mul-sub-2-1":
		a := r.Int63n(8) + 1 // 1~9
		b := r.Int63n(8) + 1 // 1~9
		c := r.Int63n(a * b)
		res = fmt.Sprintf("%d × %d - %d =", a, b, c)
	}
	return res
}
