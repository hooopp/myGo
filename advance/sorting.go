package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type ByName []Person

func (a ByName) Len() int {
	return len(a)
}

func (a ByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	// numbers := []int{5, 2, 9, 1, 5, 6}
	// sort.Ints(numbers)
	// fmt.Println("Sorted numbers:", numbers)

	// stringSlice := []string{"John", "Anthony", "Zoe"}
	// sort.Strings(stringSlice)
	// fmt.Println("Sorted strings:", stringSlice)

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	fmt.Println("Before sorting:", people)
	sort.Sort(ByAge(people))
	fmt.Println("After sorting:", people)

	people2 := []Person{
		{"Zlice", 30},
		{"Sob", 25},
		{"Aharlie", 35},
	}

	fmt.Println("Before sorting by name:", people2)
	sort.Sort(ByName(people2))
	fmt.Println("After sorting by name:", people2)

}
