package testpkg

import "fmt"

var PkgLevelVariable string = "The value of this package level variable is the current string."

func Add(x int, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}

func MultipleReturnVals(first, second string) (string, string) {
	return second, first
}

func NamedReturnValues(sum int) (factor1, factor2 int) {
	factor1 = sum / 2
	factor2 = sum * 2
	return factor1, factor2
}

func ShortAssignmentOperator() int {
	x := 225
	return x
}

func LoopDemo() {
	str_array := []string{"This", "is", "an", "array", "of", "strings"}

	fmt.Println("Using indexed counters: ")
	for i := 0; i < len(str_array); i++ {
		fmt.Println(str_array[i])
	}

	fmt.Println("Using Foreach construct:")
	for _, element := range str_array {
		fmt.Println(element)
	}
}

type Location struct {
	Lat  float64
	Long float64
}

func MapDemo() {
	Locations := make(map[string]Location)
	Locations["BellLabs"] = Location{
		40.68433, -74.39967,
	}
	fmt.Println(Locations["BellLabs"])
}
