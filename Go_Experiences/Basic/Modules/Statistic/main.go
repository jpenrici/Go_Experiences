// main.go

package main

import (
	"Statistic/stats" // Module Name + folder path
	"fmt"
)

func main() {
	data := stats.SimpleStats{
		Values: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
	}

	var proc stats.DataProcessor = data

	result, err := proc.Process()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Result:\nMean: %.2f\nSD: %.2f\n", result.Mean, result.StdDev)
}
