package main

import "fmt"

type MadMax struct {
	On          bool
	Ammo, Power int
}

func (str *MadMax) Shoot() (bo bool) {
	if str.On == true {
		if str.Ammo > 0 {
			str.Ammo--
			bo = true
		}
	} else {
		bo = false
	}
	return
}

func (str *MadMax) RideBike() (bo bool) {
	if str.On == true {
		if str.Power > 0 {
			str.Power--
			bo = true
		}
	} else {
		bo = false
	}
	return
}

func main() {

	testStruct := &MadMax{true, 1, 2}
	fmt.Println(testStruct.RideBike())
	fmt.Println(testStruct.Shoot())

	fmt.Println(testStruct)

}
