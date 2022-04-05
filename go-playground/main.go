package main

import (
	"fmt"
	"go-playground/testpkg"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Random number generator:", rand.Intn(100))

	x := 10
	y := 20

	fmt.Println("Adding", x, "and", y, "gives", testpkg.Add(x, y))
	fmt.Println("Subtracting", y, "from", x, "gives", testpkg.Subtract(x, y))

	str1, str2 := testpkg.MultipleReturnVals("First", "Second")

	fmt.Println("The values of strings from multiple return valued function is:", str1, str2)

	fmt.Println(testpkg.NamedReturnValues(200))
	fmt.Println(testpkg.PkgLevelVariable)
	fmt.Println("ShortAssignmentOperator is", testpkg.ShortAssignmentOperator())

	var intptr *int = new(int)
	*intptr = rand.Intn(3600)
	fmt.Println("Random integer from the pointer is:", *intptr)

	testpkg.LoopDemo()

	testpkg.MapDemo()
}
