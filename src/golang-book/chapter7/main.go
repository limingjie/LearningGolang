package main

import "fmt"

func main() {
	arr := []float64{123.4, 234.5, 345.6, 456.7, 567.8}
	fmt.Println(average(arr))
}

func average(arr []float64) float64 {
	var sum float64
	for _, v := range arr {
		sum += v
	}

	return sum / float64(len(arr))
}
