package main
import "fmt"
func main() {

	fmt.Print("Input first number: ")
	var a int
	fmt.Scanln(&a)

	fmt.Print("Input second number: ")
	var b int 
	fmt.Scanln(&b)

	if a > b {
		fmt.Printf("%d is greater than %d", a, b)
	} else if b > a {
		fmt.Printf("%d is greater than %d", b, a)
	} else {
		fmt.Println("Both numbers are equal.")
	}

}