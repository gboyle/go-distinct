/*

	Escape Analysis

	see:

		https://mayurwadekar2.medium.com/escape-analysis-in-golang-ee40a1c064c1

	notes:

		go build -gcflags="-m"

*/

package main

import "fmt"

func main() {
	h := heapAnalysis()
	s := stackAnalysis()

	fmt.Println("Called heapAnalysis", *h)
	fmt.Println("Called stackAnalysis", s)
}

//go:noinline
func stackAnalysis() int {
	data := 123
	return data
}

//go:noinline
func heapAnalysis() *int {
	data := 456
	return &data
}
