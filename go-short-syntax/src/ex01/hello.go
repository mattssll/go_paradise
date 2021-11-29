package main

import (
	"fmt"
	"errors"
	"math"
)

// structs 
type person struct {
	name string
	age int
}

func main () {
	syntax()
	result := sum(5,4)
	fmt.Println("resutl is:", result)
	resultSqrt, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("square root is:",resultSqrt)
	}
	// working with the struct
	p := person{name: "Jake", age: 25}
	fmt.Println(p.name, p.age, p)
	// pointers
	// a pointer is a variable that stores a mem address of a variable: 
	// pointerOfVar1 = &Var1
	pointers()
	pInc := 7
	pointersInc(&pInc) // send the pointer with the &
	fmt.Println(pInc)
}

func syntax() {
	// vars
	x:= 5
	if x > 6 {
		fmt.Println("x is bigger than 6")
	} else if x > 2 {
		fmt.Println("x is bigger than 2")
	} else {
		fmt.Println("x is smaller than 2")
	}
	
	// ARRs, SLICES, and MAPS
	// arrays have defined number of items
	var myArr1 [5]int
	myArr1[2] = 7
	fmt.Println("printing myArr1: ", myArr1)
	// also possible to do myArr1 := [5]int{5,4,3,2,1}
	// slices instead of array (use not number inside the [])
	myArr2 := []int{5,4,3,2,1}
	myArr2 = append(myArr2,0) // append create new object
	fmt.Println("printing MyArr2", myArr2)
	// maps, like dicts in python - [keys type]value type
	vertices := make(map[string][]int) // map of slices (similar to arrs)
	vertices["triangle"] = []int{10, 5, 2}
	vertices["square"] = []int{4,4,4,4}
	vertices["circle"] = []int{25,3}
	delete(vertices, "circle")
	fmt.Println(vertices)

	// LOOPS
	// for
	fmt.Println("for loop")
	for i := 0; i<5; i++ {
		fmt.Println(i)
	}
	// as a while
	fmt.Println("while loop")
	i :=0
	for i < 5 {
		fmt.Println(i)
		i++
	}
	
	// LOOPS and ARRAYS, SLICES, and MAPS
	fmt.Println("for and arrays, slices, and maps")
	mySlice := []string{"abc", "cba", "abcd", "dcba"}
	for index, value := range mySlice {
		fmt.Println("index:", index, "value:", value)
	}
	myMap := make(map[string][]string)
	myMap["name"] = []string{"mary", "john"}
	myMap["color"] = []string{"yellow", "blue"}
	for key, value := range myMap {
		fmt.Println("key:", key, "value:", value)
	}
}

func sum(x int, y int) int {
	return x + y
}

// go has no exceptions, returning multiple vals from fx:
func sqrt(x float64) (float64, error) { // can return float64 or error
	if x < 0 {
		return 0, errors.New("can't accept negative values")
	}

	return math.Sqrt(x), nil // nil if no error
}

func pointers() {
	i := 7
	pointerToIOrMemAddress := &i
	fmt.Println(pointerToIOrMemAddress) // gets memory address that gives pointer to i
	valOfPointerOfI := *pointerToIOrMemAddress
	fmt.Println(valOfPointerOfI)
}

func pointersInc(x *int) { // accept pointer with the *
	*x++ // dereference the pointer, otherwise would increment mem address
}