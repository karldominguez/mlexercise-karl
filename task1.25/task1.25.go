package main
 
import (
	"fmt"
	"math"
)
 
//Sqrt computes the square root of a number using Newton's method
func Sqrt(x float64) float64 {
	z := 1.0
	prev := 0.0
 
	for i := 1; i < 100; i++ {
		prev = z
		z = z - (((z * z) - x) / (2 * z))
 
		fmt.Println("\nPrevious val", prev)
		fmt.Println("Current val", z)
		fmt.Println("\nNumber of Loops", i)
 
		difference := z - prev
		if difference < 0 {
			difference = difference * -1
		}
 
		fmt.Println("Difference", difference)
 
		if (difference) < 0.000000000000001 {
			break
		}
	}
	return prev
}
 
func main() {
	fmt.Println("My Value", Sqrt(2))
	fmt.Println("Go Value", math.Sqrt(2))
}