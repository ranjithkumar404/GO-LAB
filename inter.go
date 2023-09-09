
// a go program to illustrate Interface concept
package main
import (
	"fmt"
	"math"
)

type shape interface{
	Area() float64
}

type circle struct{
	radius float64
}

type square struct{
	side float64
}

func (c circle) Area() float64{
	return math.Pi*c.radius*c.radius
}

func (s square) Area() float64{
	return s.side*s.side
}

func main(){
	var s shape
	var a float64
	var b float64
	fmt.Println("Enter the  radius for the circle")
	fmt.Scanln(&a)
	fmt.Println("Enter the  side for the square")
	fmt.Scanln(&b)

	s=circle{radius:a}
	fmt.Println("The area of circle is ",s.Area())
	s=square{side:b}
	fmt.Println("the area of the square is",s.Area())
}