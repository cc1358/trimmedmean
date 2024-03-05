package trimmedmean

import (
	"testing"
)

func TestTrimmedMean(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tm := New(data)

	// Test symmetric trimming
	result := tm.Calculate(0.05, 0.05) // Trim 5% from both ends
	expected := 5.5                    // Expected trimmed mean: (2+3+4+5+6+7+8+9)/8 = 5.5
	if result != expected {
		t.Errorf("Symmetric trimming: Expected %f but got %f", expected, result)
	}

	// Test asymmetric trimming
	result = tm.Calculate(0.1, 0.05) // Trim 10% from the lower end and 5% from the upper end
	expected = 5.625                 // Expected trimmed mean: (3+4+5+6+7+8+9)/7 = 5.625
	if result != expected {
		t.Errorf("Asymmetric trimming: Expected %f but got %f", expected, result)
	}

	// Test when data is empty
	emptyData := []float64{}
	tmEmpty := New(emptyData)
	result = tmEmpty.Calculate(0.1, 0.1)
	expected = 0.0 // Expected trimmed mean of empty data is 0
	if result != expected {
		t.Errorf("Empty data: Expected %f but got %f", expected, result)
	}
}
