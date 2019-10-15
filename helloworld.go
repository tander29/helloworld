package main

import (
	"fmt"

	"github.com/tander29/stringutil"
)

func main() {
	fmt.Println("hello world, I can type this in the command line anywhere, thanks goinstall")
	fmt.Println(stringutil.Reverse("hello world, I can type this in the command line anywhere, thanks goinstall"))
	fmt.Println(hoistTest())
	voidTest()
	sparkyTrick := TrickParameters{
		command: "speak",
	}
	sparky := makeDog("sparky", 3, true, func(TrickParameters) {
		fmt.Println(sparkyTrick.command)
	})
	fmt.Println(sparky.Name)
	sparky.trick(sparkyTrick)
}

func hoistTest() string {
	a := "Does hoisting work?"
	fmt.Println("hoist test")
	return a
}

func voidTest() {
	a := 1
	b := 2
	fmt.Println(a + b)
}

type Dog struct {
	Name  string
	age   int
	barks bool
	trick func(TrickParameters)
}

type TrickParameters struct {
	repeat  int
	command string
}

func makeDog(name string, age int, barks bool, trick func(TrickParameters)) Dog {
	newDog := Dog{
		Name:  name,
		age:   age,
		barks: barks,
		trick: trick,
	}
	return newDog
}
