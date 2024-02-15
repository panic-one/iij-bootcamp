package main

import "fmt"

func main() {
	var Watashi_no_Hensu string = "GYUDON"
	fmt.Println(Watashi_no_Hensu)
	main2()
}

func Eat(name string) bool {
	if name == "" {
		return false
	} else {
		fmt.Println(name)
		return true
	}
}

func main2() {
	var name1 string = "GYUDON"
	if ok := Eat(name1); !ok {
		fmt.Println("cannot eat: ", name1)
	}

	var name2 string = ""
	if ok :=Eat(name2); !ok {
		fmt.Println("cannot eat: ", name2)
	}
}
