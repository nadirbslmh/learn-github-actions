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
