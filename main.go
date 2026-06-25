package main

import (
	"fmt"
	"math"
)

func quadSolver (a, b, c float64)  {
	var x1, x2 float64
	det := (float64)((b*b) - (4 * a * c))

	
	if det > 0 {
	// real roots
		x1 = (-b + math.Sqrt(det))/ (2*a)
		x2 = (-b - math.Sqrt(det)) / (2*a)

		fmt.Println("The equation has distinct root")
		fmt.Printf("x1 = %.3f, x2 = %.3f\n", x1, x2)

	} else if det == 0{
	// equal roots
		x1 = (-b - math.Sqrt(det))/ (2*a)
		fmt.Println("The equation has equal roots")
		fmt.Printf("x1 = x2 = %.3f\n", x1)

	} else {
	// complex roots
	// rr = real root, cr = complex root

		rr := (-b / (2 * a))
		det = -1 * det
		cr := (math.Sqrt(det)) / (2 * a)

		fmt.Println("Complex roots")
		fmt.Printf("x1 = %.3f + %.3fi\n", rr, cr)
		fmt.Printf("x2 = %.3f - %.3fi\n", rr, cr)

	}
}

func main() {
	var a, b, c float64

	fmt.Println("Enter the value of a: ")
	fmt.Scan(&a)

	fmt.Println("Enter the value of b: ")
	fmt.Scan(&b)

	fmt.Println("Enter the value of c: ")
	fmt.Scan(&c)


	if a == 0 {
		fmt.Println( "Not a quadratic equation")
		return
	}


	quadSolver(a, b, c)

}