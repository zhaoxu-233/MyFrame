package main

import (
	"exercise_code/interface_struct"
	"exercise_code/learn_arry"
	"fmt"
)

type Fish struct {
	name string
}

type FakeFish Fish

func (f *Fish) swim() {
	fmt.Println("11")
}

//func defers() {
//	i := 0
//	i = 1
//	defers func() {
//		fmt.Printf("i : %p ", &i)
//	}()
//}
func main() {

	const (
		a = iota
		b
		c
		d = 9
		f
	)
	const (
		daya = 1 << iota
		dayb
		dayc
	)
	//fmt.Println(a, b, c, d, f)
	//fmt.Println(daya, dayb, dayc)
	//global.VarGlobal()
	//fmt.Println(defers.DeferReturn())
	//fmt.Println(defers.DeferReturnV1())
	//fmt.Println(defers.DeferReturnV2())
	//fmt.Println(control.IfNewVariable(120, 10))
	//control.LoopV1()
	//learn_arry.Arry()
	//learn_arry.Dem()
	//interface_struct.Create_user()
	//
	//var a8 int = 10
	//
	//var ptr *int
	//ptr = &a8
	//ptr1 := a8
	//fmt.Println(a8, &ptr, ptr, *ptr, &a8, ptr1, &ptr1)
	//fmt.Printf("%T", *ptr)
	//interface_struct.UseFish()
	//interface_struct.CreateFish()
	type intt int
	interface_struct.CreateFishP()
	fmt.Println(interface_struct.Sum[intt](1))
	interface_struct.MyPrint(33)
	var a1 = []int{1, 2, 3, 4}
	var b1 = []int{5, 6, 7, 8}
	fmt.Println(learn_arry.TestFoolopV1(a1, b1))
	fmt.Printf("%T", b1)

}
