package defers

func DeferReturn() int {
	a := 0
	defer func() {
		a = 1 //改的不是返回值的a，是栈上的a
	}()
	return a
}

func DeferReturnV1() (a int) {
	a = 0
	defer func() {
		a = 1 //改的是返回值的本体
	}()
	return a
}

type MyStruct struct {
	name string
}

func DeferReturnV2() *MyStruct {
	a := &MyStruct{
		name: "jerry",
	}
	defer func() {
		a.name = "Tom"
	}()
	return a
}
