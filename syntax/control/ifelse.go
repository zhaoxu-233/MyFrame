package control

import "fmt"

func IfNewVariable(start, end int) string {
	if distance := start - end; distance > 100 {
		return "好远"
	} else if distance > 20 {
		return "可以接受"
	} else {
		return "太远了"
	}

}

func LoopV1() {
	arr := [3]string{"H1", "H2", "H3"}
	for i, v := range arr {
		fmt.Println(&i)
		fmt.Println(arr[i])
		fmt.Println(v)
	}
}
