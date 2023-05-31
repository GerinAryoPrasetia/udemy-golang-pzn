package main

import "fmt"

func main() {
	total := sumSalary("Gerin", 100, 1000, 4, 5)
	fmt.Println(total)

	// pass slice to variadic function
	salary := []int{100, 1000, 4, 33}
	total2 := sumSalary("Zuck", salary...)
	fmt.Println(total2)
}

// varargs (variadic arguments) should be the last parameter.
// It can be zero or more. and it can't be at the beginning or middle of the parameter list.
// use variadic arguments when you don't know how many arguments will be passed to the function.
// and dont want to create a slice manually, when you call the function.
// we can also pass a slice to a variadic function.
// when pass a slice to a variadic function, we need to add ... after the slice.
func sumSalary(name string, salaries ...int) string {
	total := 0
	for _, num := range salaries {
		total += num
	}

	return fmt.Sprintf("%s has total salary of Rp %d", name, total)
}
