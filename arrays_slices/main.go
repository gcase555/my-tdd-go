package arrays_slices

func SumArray(numbers [5]int) int {
	sum := 0

	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumSlice(numbers []int) int {
	var sum int

	for _, num := range numbers {
		sum += num
	}
	return sum
}

//// Sum All that uses "make" to initialize a slice of 0's that we can iterate over and set values on
//// if you dont do this you get an error since "sums" has no initial values in the slice to set
//func SumAll(numbersToSum ...[]int) []int{
//	lengthOfNumbers := len(numbersToSum)
//	sums := make([]int, lengthOfNumbers)
//
//	for i, numbers := range numbersToSum {
//		sums[i] = SumSlice(numbers)
//
//	}
//	return sums
//}

// Sum All that uses append on an empty variable instead of "make". This works but append can be slower then the other approach.
func SumAll(numbersToSum ...[]int) []int{
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, SumSlice(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		 if len(numbers) == 0 {
		 	sums = append(sums, 0)
		 } else {
		 	tail := numbers[1:]
		 	sums = append(sums, SumSlice(tail))
		 }
	}
	return sums
}
