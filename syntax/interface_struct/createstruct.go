package interface_struct

import "fmt"

type User struct {
	Name string
	Age  int
}

func Create_user() {
	//初始化结构体
	u := User{}
	fmt.Printf("\n%+v\n", u)

	//up是指针
	up := &User{}
	fmt.Printf("%+v\n", *up)
	up2 := new(User)
	fmt.Printf("%+v\n", up2)

	//赋值方式
	u4 := User{Name: "tom", Age: 12}
	u4.Age = 15
	fmt.Println(u4)

	u5 := &User{Name: "tom", Age: 12}
	u5.Name = "jerry"
	fmt.Println(u5)

	var u6 *User
	//因为u6是指针，指针的作用是存储一个内存地址，所以需要重新声明一个变量，并把变量的地址赋值给指针
	/*
		不可以使用   u7 := &UUser{Name: "charli", Age: 14},u7是指针类型
	*/
	u7 := User{Name: "charli", Age: 14}
	u6 = &u7
	fmt.Printf("%T\n", u6)
	fmt.Printf("%p,%T\n", &u7, u7)
	fmt.Println(&u7)
}
