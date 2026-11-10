package main
import "fmt"
func main() {
	fmt.Print("Input a number: ")
	var x int
	fmt.Scanln(&x)

	if x > 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are not an adult.")
	}

}