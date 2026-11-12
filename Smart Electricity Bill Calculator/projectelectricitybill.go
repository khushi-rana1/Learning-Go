/*Go program that accepts electricity consumption (in units) for 7 days and 
calculates the total consumption, average daily consumption, & estimated electricity bill.*/

package main

import "fmt"

func main() {

    var consumption [7]float64
	fmt.Println("Enter electricity consumption for 7 days:")
	for i := 0; i < 7; i++ {

        fmt.Printf("Day %d: ", i+1)
        fmt.Scan(&consumption[i])

    }

    
    total := 0.0
    highest := consumption[0]
	for i := 0; i < 7; i++ {

        total = total + consumption[i]

        if consumption[i] > highest {
            highest = consumption[i]
        }
    }

    var rate int
	if total <= 100 {
        rate = 3
    } else if total <= 200 {
        rate = 5
    } else {
        rate = 7
    }

	bill := total * float64(rate)

    fmt.Printf("\nTotal Consumption: %.2f units\n", total)
    fmt.Printf("Highest Daily Consumption: %.2f units\n", highest)
	fmt.Printf("\nApplicable Rate: Rs. %.2f/unit\n", float64(rate))
    fmt.Printf("Estimated Electricity Bill: Rs. %.2f\n", bill)

}
