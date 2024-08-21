package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"runtime/debug"
	"sort"
	"sync"
	"time"
)

// Merge sort algorithm
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// Mandelbrot set algorithm
func mandelbrot(width, height int, maxIterations int) [][]int {
	result := make([][]int, height)
	for y := 0; y < height; y++ {
		result[y] = make([]int, width)
		for x := 0; x < width; x++ {
			c := complex(float64(x)/float64(width)*3.5-2.5, float64(y)/float64(height)*2.0-1.0)
			z := complex(0, 0)
			iterations := 0
			for real(z)*real(z)+imag(z)*imag(z) < 4 && iterations < maxIterations {
				z = z*z + c
				iterations++
			}
			result[y][x] = iterations
		}
	}
	return result
}

func calculatePi(iterations int) float64 {
	const numPoints = 1000000
	var pi float64
	for i := 0; i < iterations; i++ {
		for j := 0; j < numPoints; j++ {
			pi += (math.Pi / numPoints)
		}
		if i%10 == 0 {
			printProgress("Pi Calculation", i, iterations)
		}
	}
	return pi
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func sorting(array []int, total int) {
	sort.Ints(array)
	for i := 0; i < total; i++ {
		if i%100 == 0 {
			printProgress("Sorting", i, total)
		}
	}
}

func matrixMultiplication(size int, total int) {
	matrixA := make([][]float64, size)
	matrixB := make([][]float64, size)
	result := make([][]float64, size)

	for i := range matrixA {
		matrixA[i] = make([]float64, size)
		matrixB[i] = make([]float64, size)
		result[i] = make([]float64, size)
		for j := range matrixA[i] {
			matrixA[i][j] = rand.Float64()
			matrixB[i][j] = rand.Float64()
		}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sum := 0.0
			for k := 0; k < size; k++ {
				sum += matrixA[i][k] * matrixB[k][j]
			}
			result[i][j] = sum
		}
		if i%10 == 0 {
			printProgress("Matrix Multiplication", i, size)
		}
	}
}

func monteCarloSimulation(points int) float64 {
	var insideCircle int
	for i := 0; i < points; i++ {
		x, y := rand.Float64(), rand.Float64()
		if x*x+y*y <= 1 {
			insideCircle++
		}
		if i%100000 == 0 {
			printProgress("Monte Carlo Simulation", i, points)
		}
	}
	return 4.0 * float64(insideCircle) / float64(points)
}

func printProgress(task string, current, total int) {
	percentage := float64(current) / float64(total) * 100
	fmt.Printf("\r%s: %.2f%% complete", task, percentage)
	if current == total-1 {
		fmt.Println()
	}
}

func printMemoryUsage() float64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	memUsage := float64(m.Alloc) / 1024 / 1024
	fmt.Printf("\nMemory usage: %.2fMB\n", memUsage)
	return memUsage
}

func saveResultsToCSV(results [][]string) {
	err := os.MkdirAll("results", os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	currentTime := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("results/result_%s.csv", currentTime)

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range results {
		err := writer.Write(record)
		if err != nil {
			fmt.Println("Error writing to CSV file:", err)
		}
	}

	fmt.Println("Results saved to", filename)
}

func benchmark() {
	fmt.Println("Benchmarking system characteristics...")
	start := time.Now()

	results := [][]string{
		{"Task", "Time Taken (s)", "Memory Usage (MB)"},
	}

	benchmarkSingleCore(&results)
	benchmarkMultiCore(&results)

	end := time.Now()
	timeTaken := end.Sub(start).Seconds()

	results = append(results, []string{
		"Total Benchmark Time", fmt.Sprintf("%.2f", timeTaken), fmt.Sprintf("%.2f", printMemoryUsage()),
	})

	saveResultsToCSV(results)
}

func benchmarkSingleCore(results *[][]string) {
	fmt.Println("Benchmarking single core...")

	start := time.Now()
	for i := 0; i < 5000; i++ {
		calculatePi(20)
	}
	timeTaken := time.Since(start).Seconds()
	memUsage := printMemoryUsage()

	*results = append(*results, []string{"Single-core Pi Calculation", fmt.Sprintf("%.2f", timeTaken), fmt.Sprintf("%.2f", memUsage)})

	start = time.Now()
	for i := 0; i < 5000; i++ {
		fibonacci(35)
		if i%100 == 0 {
			printProgress("Fibonacci Calculation", i, 5000)
		}
	}
	timeTaken = time.Since(start).Seconds()
	memUsage = printMemoryUsage()

	*results = append(*results, []string{"Single-core Fibonacci Calculation", fmt.Sprintf("%.2f", timeTaken), fmt.Sprintf("%.2f", memUsage)})

	start = time.Now()
	sorting(make([]int, 100000000), 100000000)
	timeTaken = time.Since(start).Seconds()
	memUsage = printMemoryUsage()

	*results = append(*results, []string{"Single-core Sorting", fmt.Sprintf("%.2f", timeTaken), fmt.Sprintf("%.2f", memUsage)})

	start = time.Now()
	matrixMultiplication(2000, 2000)
	timeTaken = time.Since(start).Seconds()
	memUsage = printMemoryUsage()

	*results = append(*results, []string{"Single-core Matrix Multiplication", fmt.Sprintf("%.2f", timeTaken), fmt.Sprintf("%.2f", memUsage)})

	start = time.Now()
	monteCarloSimulation(50000000)
	timeTaken = time.Since(start).Seconds()
	memUsage = printMemoryUsage()

	*results = append(*results, []string{"Single-core Monte Carlo Simulation", fmt.Sprintf("%.2f", timeTaken), fmt.Sprintf("%.2f", memUsage)})

	start = time.Now()
	mergeSort(make([]int, 1000000))
	timeTaken = time.Since(start).Seconds()
	memUsage = printMemoryUsage()

	*results = append(*results, []string{"Single-core Merge Sort", fmt.Sprintf("%.2f", timeTaken), fmt.Sprintf("%.2f", memUsage)})

	start = time.Now()
	mandelbrot(800, 600, 1000)
	timeTaken = time.Since(start).Seconds()
	memUsage = printMemoryUsage()

	*results = append(*results, []string{"Single-core Mandelbrot Set", fmt.Sprintf("%.2f", timeTaken), fmt.Sprintf("%.2f", memUsage)})
}

func benchmarkMultiCore(results *[][]string) {
	fmt.Println("Benchmarking multi-core...")

	var wg sync.WaitGroup
	numCores := runtime.NumCPU()

	start := time.Now()
	wg.Add(numCores)
	for i := 0; i < numCores; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000/numCores; j++ {
				calculatePi(20)
				fibonacci(35)
			}
		}()
	}
	wg.Wait()
	timeTaken := time.Since(start).Seconds()
	memUsage := printMemoryUsage()

	*results = append(*results, []string{"Multi-core Pi and Fibonacci Calculation", fmt.Sprintf("%.2f", timeTaken), fmt.Sprintf("%.2f", memUsage)})

	start = time.Now()
	wg.Add(numCores)
	for i := 0; i < numCores; i++ {
		go func() {
			defer wg.Done()
			sorting(make([]int, 100000000), 100000000)
			matrixMultiplication(2000, 2000)
			monteCarloSimulation(50000000)
		}()
	}
	wg.Wait()
	timeTaken = time.Since(start).Seconds()
	memUsage = printMemoryUsage()

	*results = append(*results, []string{"Multi-core Sorting, Matrix Multiplication, and Monte Carlo Simulation", fmt.Sprintf("%.2f", timeTaken), fmt.Sprintf("%.2f", memUsage)})

	start = time.Now()
	wg.Add(numCores)
	for i := 0; i < numCores; i++ {
		go func() {
			defer wg.Done()
			mergeSort(make([]int, 1000000))
			mandelbrot(800, 600, 1000)
		}()
	}
	wg.Wait()
	timeTaken = time.Since(start).Seconds()
	memUsage = printMemoryUsage()

	*results = append(*results, []string{"Multi-core Merge Sort and Mandelbrot Set", fmt.Sprintf("%.2f", timeTaken), fmt.Sprintf("%.2f", memUsage)})
}

func main() {
	fmt.Println("Starting Benchmarking...")
	debug.SetGCPercent(-1) // Disable garbage collection to better measure memory usage
	benchmark()
	fmt.Println("Benchmarking completed.")
}
