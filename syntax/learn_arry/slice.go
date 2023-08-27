package learn_arry

import "fmt"

func Dem() {
	a := []int{}
	for i := 0; i < 19; i++ {

		a = append(a, 1)

	}
	fmt.Print(a)
}

func main1() {
	var s []int //创建一个空切片
	for i := 1; i <= 18; i++ {
		s = append(s, i) //在切片末尾追加i
	}
	fmt.Println(s)
}

func TestFoolopV1(a, b []int) []int {
	c := append(a, b...)
	return c
}
