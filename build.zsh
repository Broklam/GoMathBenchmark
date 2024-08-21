#!/bin/zsh


output_dir="builds"
mkdir -p "$output_dir"

echo "Compiling for Apple ARM64..."
GOOS=darwin GOARCH=arm64 go build -o "$output_dir/benchmark_apple_arm64" main.go

echo "Compiling for Apple Intel64..."
GOOS=darwin GOARCH=amd64 go build -o "$output_dir/benchmark_apple_intel64" main.go

echo "Compiling for Windows ARM..."
GOOS=windows GOARCH=arm64 go build -o "$output_dir/benchmark_windows_arm.exe" main.go

echo "Compiling for Windows X64..."
GOOS=windows GOARCH=amd64 go build -o "$output_dir/benchmark_windows_x64.exe" main.go


echo "Compiling for Linux X64..."
GOOS=linux GOARCH=amd64 go build -o "$output_dir/benchmark_linux_x64" main.go

echo "Compilation completed!"
