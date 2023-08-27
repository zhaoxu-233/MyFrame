package interface_struct

import "fmt"

type fish struct {
	name string
}

type fakefish fish

type pfish struct {
	*fish
}

func UseFish() {
	f1 := fish{
		name: "1",
	}
	f2 := fakefish{}
	f3 := fish(f2)
	fmt.Println(f1, f2, f3)
	f2.name = "bb"
	f3 = fish(f2)
	fmt.Println(f1, f2, f3)

}

func (f fakefish) Swim() {
	fmt.Println("fish can swim in sea")
}

func CreateFish() {
	var Fish fakefish
	Fish = fakefish{name: "sss"}
	Fish.Swim()
}

func CreateFishP() {
	var PFish pfish
	PFish = pfish{&fish{name: "qw"}}
	fmt.Printf("PFish type is %T  value is %+v\n", PFish, &PFish)
}
