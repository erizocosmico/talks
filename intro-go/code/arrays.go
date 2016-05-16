// And array has a fixed size
var numberArray = [4]int{1, 2, 3, 4}

var numbers = []int{1, 2, 3, 4}
var noNumbers = nil

// We can add more numbers
numbers = append(numbers, 5)

// nil is a valid slice
noNumbers = append(noNumbers, 1)

// Make the slice with length of 0 but capacity for 50 elements
var sliceWithCapacity = make([]int, 0, 50)
