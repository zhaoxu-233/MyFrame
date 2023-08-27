package learn_arry

import "fmt"

func Arry() {
	a := [3]int{1}
	fmt.Printf("a value is %v,len is %d,cap is %d\n", a, len(a), cap(a))
}
