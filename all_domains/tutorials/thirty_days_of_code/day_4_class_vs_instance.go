package thirty_days_of_code

import (
	"fmt"

	"github.com/ereminIvan/hackerrank/utils"
)

type person struct {
	age int
}

func (p person) NewPerson(initialAge int) person {
	//Add some more code to run some checks on initialAge
	if initialAge < 0 {
		initialAge = 0
		fmt.Println("Age is not valid, setting age to 0.")
	}
	p.age = initialAge
	return p
}

func (p person) amIOld() {
	if p.age < 13 {
		fmt.Println("You are young.")
	} else if p.age >= 13 && p.age < 18 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are old.")
	}
}

func (p person) yearPasses() person {
	//Increment the age of the person in here
	p.age++
	return p
}

func Day4ClassVsInstance(getInput utils.GetInput) {
	input := getInput()
	var T, age int
	fmt.Fscanf(input, "%d", &T)

	for i := 0; i < T; i++ {
		fmt.Fscanf(input, "%v", &age)
		p := person{age: age}
		p = p.NewPerson(age)
		p.amIOld()
		for j := 0; j < 3; j++ {
			p = p.yearPasses()
		}
		p.amIOld()
		fmt.Println()
	}
}
