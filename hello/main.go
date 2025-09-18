// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello,Go!")

// }
package main

import (
	"fmt"
)

const (
	_ = iota
	a = 1 << iota
	b
	c
)
const d = iota

 
func main() {
	fmt.Println("hello world")
	//rough work
	//var grades [4]int
	//grades := [...]int{97, 83, 90, 99}
	grades := [4]int{97, 83, 90}
	//Here down the statement, here sampleGrade is not pointing to grades, instead it copying values of grade.
	sampleGrade := grades
	sampleGrade[3] = 99
	//Here you can see pointerGrade is pointing to grades not copying values from grades
	pointerGrade := &grades
	pointerGrade[3] = 66
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", c, c)
	fmt.Printf("%v %T\n", d, d)
	fmt.Printf("%v %T\n", grades, grades)
	fmt.Printf("%v %T\n", pointerGrade, pointerGrade)
	fmt.Println(sampleGrade)
	//here we can notice that i Have kept "," for instead of "+" for concatination becuase len(grades) gives integer and other is a string , so it is mismatch type, give errors
	fmt.Println("The length of grades", len(grades))
	var matrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(matrix)

	//slice
	sampleSlice := []int{1, 2, 3}
	//reference to slice
	refSlice := sampleSlice
	refSlice[1] = 5
	fmt.Println(sampleSlice)
	fmt.Println(refSlice)
	fmt.Println("the length of sample slice", len(sampleSlice))
	fmt.Println("the capacity of sample slice", cap(sampleSlice))
	baseArray := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	// all below are pointing to baseArray, so if any changes done to slice base Array and other slices get effected or modified
	a1 := baseArray[:]
	a2 := baseArray[3:7]
	a3 := baseArray[6:]
	a4 := baseArray[1:5]
	fmt.Printf("%v %v %v\n", a1, len(a1), cap(a1))
	fmt.Printf("%v %v %v\n", a2, len(a2), cap(a2))
	fmt.Printf("%v %v %v\n", a3, len(a3), cap(a3))
	fmt.Printf("%v %v %v\n", a4, len(a4), cap(a4))
	//slice using make
	//a:=make([]int,lengthofslice,capacity);
	//a=append(a,1,3,5,53,9  );

	//map
	//product:=make(map[string]int)
	product := map[string]int{
		"keyboard": 2000,
		"mouse":    1000,
		"monitor":  7000,
	}
	//can't guarntee order of elements.
	product["laptop"] = 70000
	fmt.Println(product)
	delete(product, "keyboard")
	//even after deletion when we try to access , instead of giving error, it gives value as 0
	//in map if we try to access an value using key , which is not part, we don't get any error instead we get value as 0
	pop, ok := product["monitor"]
	//by using above method we can check whether value is present or not
	// _,ok:=product["key"], if ok is true, element is present else it is not present
	fmt.Println(pop, " ", ok)
	//in map when assigned other variable, reference is passed
	p := product
	p["monitor"] = 7500
	fmt.Println(p, "  ", len(product))
	fmt.Println(product["mouse"], " ", product["keyboard"])

}
