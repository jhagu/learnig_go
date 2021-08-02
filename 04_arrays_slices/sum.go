package arrays_slices

func Sum(numbers []int) int {
	var sum int
	/* for i := 0; i < 5; i++ {
		sum += numbers[i]
	} */

	// range -> for index, element := range elements
	// the underscore (_) omit the index (we dont need it in this example)
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {

	/* make allows you to create a slice with a starting capacity,
	in this case, the length of the numbersToSum parameter
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}  */

	// another approach, no need to woory about the capacity
	// a cleaner way IMO

	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:] // everything except the head
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
