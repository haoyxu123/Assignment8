package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"sort"
	"strconv"
	"testing"
	"time"
)

func BootstrapSample(data []int, numSamples int) []float64 {
	bootstrapMeans := make([]float64, numSamples)
	for i := 0; i < numSamples; i++ {
		sum := 0
		for j := 0; j < len(data); j++ {
			index := rand.Intn(len(data))
			sum += data[index]
		}
		bootstrapMeans[i] = float64(sum) / float64(len(data))
	}
	return bootstrapMeans
}

// CalculateConfidenceInterval calculates the confidence interval from bootstrap samples
func CalculateConfidenceInterval(samples []float64, confidenceLevel float64) (float64, float64) {
	sort.Float64s(samples)
	lowerIndex := int((1 - confidenceLevel) / 2 * float64(len(samples)))
	upperIndex := int(float64(len(samples)) * (confidenceLevel + (1-confidenceLevel)/2))
	return samples[lowerIndex], samples[upperIndex]
}

func main() {
	cpuProfile, err := os.Create("cpu_profile.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(cpuProfile)
	defer pprof.StopCPUProfile()
	startTime := time.Now()
	rand.Seed(time.Now().UnixNano())

	file, err := os.Open("C:/Users/haoyx/Desktop/titanic/train.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var data []int

	_, err = reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if record[2] == "1" {
			survived, err := strconv.Atoi(record[1])
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, survived)
		}
		endTime := time.Now()
		duration := endTime.Sub(startTime)
		fmt.Printf("Execution time: %v\n", duration)
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
		fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
		fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
		fmt.Printf("\tNumGC = %v\n", m.NumGC)
	}

	// Perform bootstrapping
	numBootstrapSamples := 1000
	bootstrapResults := BootstrapSample(data, numBootstrapSamples)

	// Calculate 95% confidence interval
	lower, upper := CalculateConfidenceInterval(bootstrapResults, 0.95)
	fmt.Printf("95%% Confidence Interval for the survival rate of 1st class passengers: [%f, %f]\n", lower, upper)
	memProfile, err := os.Create("mem_profile.mprof")
	if err != nil {
		log.Fatal(err)
	}
	defer memProfile.Close()
	runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(memProfile); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
func TestBootstrapSample(t *testing.T) {
	data := []int{1, 0, 1, 1, 0} // Simplified dataset
	numSamples := 10
	results := BootstrapSample(data, numSamples)
	if len(results) != numSamples {
		t.Errorf("Expected %d samples, got %d", numSamples, len(results))
	}
}

func BenchmarkBootstrapSample(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = rand.Intn(2) // Simplified random data
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BootstrapSample(data, 1000)
	}
}
