//defines a function which take parameters and returns their sum. 
package main
import "fmt"

func myFunction(x int, y int) int {
	return x + y
}

 func main() {
	
	var x int
	var y int

	fmt.Print("Input first number: ")
	fmt.Scanln(&x)

	fmt.Print("Input second number: ")
	fmt.Scanln(&y)

	result := myFunction(x, y)
	fmt.Print("The sum is: ", result)

}