package main

//import "fmt"

//func main() {
//	Negative := MakeNegative(1)
//	fmt.Println(Negative)
//}

func MakeNegative(x int) int {

	if x >= 0 {
		return -x
	}
	return x
}
