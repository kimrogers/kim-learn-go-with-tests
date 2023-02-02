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

func SumAll(numbersToSum ...[]int) []int {

	// // initial
	// numOfSlicesProvided := len(numbersToSum)
	// sliceResults := make([]int, numOfSlicesProvided)
	// for index, value := range numbersToSum {
	// 	sliceResults[index] = Sum(value)
	// }
	// return sliceResults

	// refactored to reduce focus on capacity
	var sliceResults []int
	for _, value := range numbersToSum {
		sliceResults = append(sliceResults, Sum(value))
	}
	return sliceResults

}

func SumAllTails(numbersToSum ...[]int) []int {

	var sliceResults []int
	for _, value := range numbersToSum {
		if len(value) == 0 {
			sliceResults = append(sliceResults, 0)
		} else {
			tail := value[1:]
			sliceResults = append(sliceResults, Sum(tail))
		}
	}
	return sliceResults
}
