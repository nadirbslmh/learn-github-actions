package main

func main() {

}

func sum(numbers []int) int {
	result := 0

	for _, num := range numbers {
		result += num
	}

	return result
}

func add(a, b int) int {
	return a + b
}
