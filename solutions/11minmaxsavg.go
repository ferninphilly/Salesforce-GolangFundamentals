// Write a function `minmaxavg` which accepts a float64 slice and returns
// three values, the minimum value in the slice, the maximum value in the
// slice, and the average of all the values
package main

import "fmt"

func minmaxavg(f []float64) (float64, float64, float64) {
	min, max := f[0], f[0]
	sum := f[0]

	for i := 1; i < len(f); i++ {
		if f[i] > max {
			max = f[i]
		}
		if f[i] < min {
			min = f[i]
		}
		sum += f[i]
	}
	avg := sum / float64(len(f))
	return min, max, avg
}

func main() {
	nums := []float64{1.2, -4.5, 6.7, -13.9}
	fmt.Println(minmaxavg(nums))
}
