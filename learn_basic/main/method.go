package main

import (
	"math"
)

// các hàm bên trong struct

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//func main() {
//	v := Vertex{3, 4}
//	fmt.Println(v.Abs())
//	v.Scale(2)
//	fmt.Println(v.Abs())
//	var x Vertex
//	x.Scale(2) // ok
//	p := &v
//	p.Scale(2) // ok
//
//	ScaleFunc(x, 10)  // fail
//	ScaleFunc(&v, 10) // ok
//
//	// ==> với method thì thằng sẽ được hiểu là c giá trị lẫn con trỏ, với function thì không
//	// p.Abs()  = (*p).Abs()
//
//	// 80% làm qua poiter nhé !
//}
