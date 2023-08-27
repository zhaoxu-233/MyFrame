package interface_struct

import "fmt"

type list[T any] interface {
}

func Sum[T Number](vals ...T) T {
	var res T
	for _, i := range vals {
		res = res + i
	}
	return res
}

//类型约束简洁写法
type Number interface {
	~int | string | float64
}

func MyPrint[C any](c C) C {
	fmt.Println(c)
	return c
}
