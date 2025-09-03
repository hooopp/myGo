package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type By func(p1, p2 *Person) bool

type personSorter struct {
	people []Person
	by func(p1, p2 *Person) bool
}

func (s *personSorter) Len() int {
	return len(s.people)
}

func (s *personSorter) Less (i, j int) bool {
	return s.by(&s.people[i], &s.people[j])
}

func (s *personSorter) Swap(i, j int) {
	s.people[i], s.people[j] = s.people[j], s.people[i]
}

func (by By) Sort (people []Person) {
	ps := &personSorter{
		people: people,
		by:     by,
	}
	sort.Sort(ps)
}

func main() {
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	By(func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}).Sort(people)

	fmt.Println(people)
}
