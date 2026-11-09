package main
import "fmt"

func main() {

	//declares with defined lengths
	var array1 = [3]int{7, 14, 21}
	fmt.Println(array1)


	//declares with defined lengths
	var array2 = []int{5, 10, 15, 20, 25}
	fmt.Println(array2)
    fmt.Println(array2[4])


	//changing elements of an array
	var array3 =[2]string{"Hello", "World"}
	array3[1] = "Golang"
	fmt.Println(array3)


	//initialize only specific elements & find length of an array
	var array4 = [3]int{0:3, 2:12}
	fmt.Println(array4)
	fmt.Println(len(array4))

}