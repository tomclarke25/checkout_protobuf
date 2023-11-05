# Benchmarking Serialization and Deserialization in Go

This repository contains benchmark tests for serialization and deserialization using JSON and Protobuf in Go. The tests are designed to compare the performance of these two popular data formats.

## Test Results

The benchmark tests were run on a machine with the following specifications:

- Processor: Apple M1 Pro
- RAM: 16 GB
- OS: macOS Sonoma 14.0

The results of the benchmark tests are as follows:

## Comparison Test

The comparison test benchmarks the serialization and deserialization of a simple Person struct using both JSON and Protobuf.

- JSON Serialization: 5605927 iterations, 205.1 ns/op, 64 B/op, 1 allocs/op
- JSON Deserialization: 1358926 iterations, 876.3 ns/op, 368 B/op, 8 allocs/op
- Protobuf Serialization: 11921901 iterations, 100.3 ns/op, 48 B/op, 1 allocs/op
- Protobuf Deserialization: 7280392 iterations, 164.5 ns/op, 144 B/op, 3 allocs/op

## Complex Test

The complex test benchmarks the serialization and deserialization of a Person struct with a slice of Person structs as friends using both JSON and Protobuf.

- JSON Serialization: 151064 iterations, 7929 ns/op, 3201 B/op, 1 allocs/op
- JSON Deserialization: 29268 iterations, 40671 ns/op, 9064 B/op, 218 allocs/op
- Protobuf Serialization: 356241 iterations, 3364 ns/op, 2048 B/op, 1 allocs/op
- Protobuf Deserialization: 153213 iterations, 7655 ns/op, 8688 B/op, 160 allocs/op

## Conclusion

From the benchmark results, it can be observed that Protobuf is faster and more memory efficient than JSON for both serialization and deserialization. This is true for both simple and complex data structures.

However, it's important to note that while Protobuf provides better performance, JSON is more human-readable and widely used. Therefore, the choice between JSON and Protobuf should be made based on the specific requirements of your project.
