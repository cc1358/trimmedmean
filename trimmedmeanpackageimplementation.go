package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/cc1358/trimmedmean"
)

func main() {
	// Generate random number seed
	rand.Seed(time.Now().UnixNano())

	// Generate random integer data
	intData := make([]float64, 100)
	for i := range intData {
		intData[i] = float64(rand.Intn(1000)) // Generate integers between 0 and 999
	}

	// Generate random float data
	floatData := make([]float64, 100)
	for i := range floatData {
		floatData[i] = rand.Float64() * 1000 // Generate floats between 0 and 1000
	}

	// Create TrimmedMean instances
	intTrimmedMean := trimmedmean.New(intData)
	floatTrimmedMean := trimmedmean.New(floatData)

	// Calculate trimmed mean with trim = 0.05
	intSortedData := make([]float64, len(intData))
	copy(intSortedData, intData)
	sort.Float64s(intSortedData)
	intTrimmedMeanResult := intTrimmedMean.Calculate(0.05, 0.05)

	floatSortedData := make([]float64, len(floatData))
	copy(floatSortedData, floatData)
	sort.Float64s(floatSortedData)
	floatTrimmedMeanResult := floatTrimmedMean.Calculate(0.05, 0.05)

	// Print results
	fmt.Println("Trimmed Mean for Integer Data:", intTrimmedMeanResult)
	fmt.Println("Trimmed Mean for Float Data:", floatTrimmedMeanResult)
}
