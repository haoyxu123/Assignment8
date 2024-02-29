package main

import (
	"math/rand"
	"testing"
)

func TestBootstrapSample(t *testing.T) {
	rand.Seed(0)

	data := []int{10, 20, 30, 40, 50}
	numSamples := 100

	bootstrapMeans := BootstrapSample(data, numSamples)

	if len(bootstrapMeans) != numSamples {
		t.Errorf("expected %d bootstrap means; got %d", numSamples, len(bootstrapMeans))
	}

	for _, mean := range bootstrapMeans {
		if mean < float64(data[0]) || mean > float64(data[len(data)-1]) {
			t.Errorf("bootstrap mean %f is out of range of the original data", mean)
		}
	}
}

func TestCalculateConfidenceInterval(t *testing.T) {
	samples := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	confidenceLevel := 0.95

	lower, upper := CalculateConfidenceInterval(samples, confidenceLevel)

	if lower != samples[0] || upper != samples[len(samples)-1] {
		t.Errorf("expected confidence interval [%f, %f]; got [%f, %f]", samples[0], samples[len(samples)-1], lower, upper)
	}

}

func BenchmarkBootstrapSample(b *testing.B) {
	data := []int{1, 2, 3, 4, 5}

	numSamples := 1000

	for n := 0; n < b.N; n++ {
		BootstrapSample(data, numSamples)
	}
}
