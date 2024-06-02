package main

import "fmt"

func main() {
	showSum()

}

func showSum() {
	53

	floats := []float64{1.0, 2.0, 3.0}
	ints := []int64{1, 2, 3}
	fmt.Println(sum(floats))
	fmt.Println(sum[int64](ints))
}

func sum[V int64 | float64](numbers []V) V {

	var sum V
	for _, num := range numbers {
		sum += num
	}
	return sum
}
