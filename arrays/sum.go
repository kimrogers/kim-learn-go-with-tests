package main

func Sum(numbers []int) int {
	total := 0
	for _, value := range numbers {
		total += value
		// fmt.Printf("index: %v value: %v\n", index, value)
	}

	return total
}

func main() {
	Sum([]int{4, 5, 6, 7})
}

func SumAll(array1 []int, array2 []int) []int {

	// return []int{0}
	// return [0,1]
	// return array1[0]

}
