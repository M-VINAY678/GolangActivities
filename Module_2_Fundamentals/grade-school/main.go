package main

import (
	"fmt"
	"sort"
)

type Student struct {
	roster map[int][]string
}

func (r *Student) add(grade int, name string) {
	for _, value := range r.roster {

		for i := 0; i < len(value); i++ {
			if value[i] == name {
				fmt.Println("Invalid name")
				return
			}
		}
	}
	r.roster[grade] = append(r.roster[grade], name)
}
func (r *Student) getrosterByGrades(grade int) []string {
	sort.Strings(r.roster[grade])
	return r.roster[grade]
}
func (r *Student) getAllStudents() []string {
	var key []int
	for index := range r.roster {
		key = append(key, index)
	}
	sort.Ints(key)
	var results []string
	for _, value := range key {
		results = append(results, r.getrosterByGrades(value)...)
	}
	return results
}
func main() {
	s := Student{make(map[int][]string)}
	s.add(5, "Jim")
	s.add(1, "Anna")
	s.add(1, "Barb")
	s.add(1, "Charlie")
	s.add(2, "Alex")
	s.add(2, "Peter")
	s.add(2, "Zoe")
	s.add(2, "Zoe")
	fmt.Println(s.getrosterByGrades(2))
	fmt.Println(s.getAllStudents())
}
