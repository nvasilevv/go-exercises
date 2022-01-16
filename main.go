package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func makeRange(min int, max int, doPrint bool) []int {
	a := make([]int, max-min+1)
	for i := range a {
		if doPrint {
			fmt.Println(min + i)
		}
		a[i] = min + i
	}
	return a
}

func raiseToPower(number float64, toPowOf float64) float64 {
	return math.Pow(number, toPowOf)
}

func generateRandomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func currentDate(format string) string {
	current := time.Now().Local()
	return current.Format(format)
}

func main() {
	b := makeRange(1, 1000, true)
	fmt.Println(generateRandomNumber(1, 10))
	fmt.Println(currentDate("2006-01-02"))
	fmt.Println(raiseToPower(2, 11))
	fmt.Println(b)
	fmt.Println("Hello World")
	time.Sleep(5 * time.Second)
}
