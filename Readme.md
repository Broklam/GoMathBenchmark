# System Benchmarking Tool

## Table of Contents

- [Introduction](#introduction)
- [Benchmark Parts](#benchmark-parts)
  - [Pi Calculation](#pi-calculation)
  - [Fibonacci Calculation](#fibonacci-calculation)
  - [Sorting](#sorting)
  - [Matrix Multiplication](#matrix-multiplication)
  - [Monte Carlo Simulation](#monte-carlo-simulation)
  - [Merge Sort](#merge-sort)
  - [Mandelbrot Set Generation](#mandelbrot-set-generation)
- [Single-core vs Multi-core Benchmarking](#single-core-vs-multi-core-benchmarking)
- [Memory Usage Reporting](#memory-usage-reporting)
- [Results Saving](#results-saving)
- [How to Run the Benchmark](#how-to-run-the-benchmark)
- [License](#license)

**PLANS**
[ ] Make an extensive GUI

[ ] Make test set adequate with proper RAM allocation

[ ] Maybe rewrite (and learn lol) c

[ ] Make a simple web server with leaderboard a.k.a. geekbench stuff

## Introduction

The System Benchmarking Tool is a Go-based program designed to assess the performance of your computer by running several computationally intensive tasks. The tool tests the system's capabilities in both single-core and multi-core environments, measuring execution time and memory usage for various algorithms.

## Benchmark Parts

The tool consists of the following benchmarks:

### Pi Calculation

The Pi calculation test uses a Monte Carlo method to estimate the value of Pi. It simulates random points inside a unit square and counts how many fall within a quarter circle inscribed within the square. The more points simulated, the closer the approximation to Pi.

### Fibonacci Calculation

This benchmark tests the recursive calculation of the Fibonacci sequence, where each number is the sum of the two preceding ones. The task is computationally intensive due to its exponential time complexity, particularly for larger numbers in the sequence.

### Sorting

The sorting benchmark measures the time taken to sort a large array of integers using Go's built-in `sort.Ints()` function, which implements an optimized QuickSort algorithm.

### Matrix Multiplication

This test benchmarks the multiplication of two large matrices. Matrix multiplication is computationally expensive, with a time complexity of O(n³), making it a good indicator of a system's computational throughput.

### Monte Carlo Simulation

This benchmark simulates random sampling to solve mathematical problems. The tool uses the Monte Carlo method to estimate areas under curves, the value of Pi, and other complex integrals. It is useful for testing floating-point arithmetic performance.

### Merge Sort

Merge sort is a classic divide-and-conquer algorithm. This benchmark tests the system's ability to handle sorting tasks that require merging two halves of an array repeatedly. It's known for its stable sorting with O(n log n) complexity.

### Mandelbrot Set Generation

The Mandelbrot set is a famous fractal used in mathematics to test graphical rendering and computational power. This benchmark generates a Mandelbrot set image by iterating a complex equation for each pixel and determining whether it belongs to the set.

## Single-core vs Multi-core Benchmarking

The tool offers two modes:

1. **Single-core Benchmarking**: Tasks are executed using only one CPU core. This mode provides insight into the single-threaded performance of the system.

2. **Multi-core Benchmarking**: Tasks are executed concurrently using all available CPU cores. This mode tests the system's ability to parallelize tasks and provides a measure of overall system throughput.

## Memory Usage Reporting

During the benchmarking process, the tool tracks memory usage. The Go runtime's memory statistics are used to report the peak memory consumed during each benchmark. Memory usage is displayed after each task to give a clear view of how resource-intensive each test is.

## Results Saving

After the benchmarking process completes, results are saved to a CSV file within a `results` directory. The filename includes a timestamp to prevent overwriting previous results. The CSV file contains the following columns:

- Task Name
- Time Taken (seconds)
- Memory Usage (MB)

This allows for easy comparison of performance over time or across different systems.

## How to Run the Benchmark

### Prerequisites

- [Go](https://golang.org/dl/) installed on your system.
- Basic understanding of the terminal/command line.

### Steps

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/your-username/system-benchmarking-tool.git
   cd system-benchmarking-tool
2. **Run the Benchmark**
   go run main.go 

   OR 

   Run executable according to your system from distros

   **Screenshots**

   <img width="682" alt="Screen Shot 2024-08-21 at 15 56 00" src="https://github.com/user-attachments/assets/87b2944d-72b2-4afa-9ad3-1a5b75906cf8">
   <img width="713" alt="Screen Shot 2024-08-21 at 15 56 36" src="https://github.com/user-attachments/assets/3d077993-5a5e-4dff-872e-ca040d3c285c">



