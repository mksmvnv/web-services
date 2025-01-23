package main

import "fmt"

func ExecutePipeline() {
}

func SingleHash() {
}

func MultiHash() {
}

func CombineResults() {
}

func main() {
	inputData := []int{1, 2}
	in := make(chan int, 1)

	go func() {
		for _, num := range inputData {
			in <- num
		}
		close(in)
	}()

	for val := range in {
		fmt.Println(val)
	}

}
