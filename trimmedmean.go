package trimmedmean

import "sort"

type TrimmedMean struct {
	data []float64
}

func New(data []float64) *TrimmedMean {
	return &TrimmedMean{data: data}
}

func (tm *TrimmedMean) Calculate(trimLower, trimUpper float64) float64 {
	if len(tm.data) == 0 {
		return 0 // or return an error
	}

	sortedData := make([]float64, len(tm.data))
	copy(sortedData, tm.data)
	sort.Float64s(sortedData)

	trimCount := int(float64(len(sortedData)) * trimLower)
	if trimCount > 0 {
		sortedData = sortedData[trimCount:]
		sortedData = sortedData[:len(sortedData)-trimCount]
	}

	trimCount = int(float64(len(sortedData)) * trimUpper)
	if trimCount > 0 {
		sortedData = sortedData[:len(sortedData)-trimCount]
		sortedData = sortedData[trimCount:]
	}

	sum := 0.0
	for _, value := range sortedData {
		sum += value
	}
	return sum / float64(len(sortedData))
}
