package main

import "fmt"


func Eat(name string) bool {
	if name == "" {
		return false
	} else {
		fmt.Println(name)
		return true
	}
}

func main() {
	var name1 string = "GYUDON"
	if ok := Eat(name1); !ok {
		fmt.Println("cannt eat: ", name1)
	}

	var name2 string = ""
	if ok := Eat(name2); !ok {
		fmt.Println("cannt eat: ", name2)
		main2()
}
}

func Eat2(name string) (bool, error) {
	if name == "" {
		return false, fmt.Errorf("name is empty.")
	}
	fmt.Println(name)
	return true, nil
}

func main2() {
	var name1 string = "GYUDON"
	if _, err := Eat2(name1); err != nil {
		fmt.Println("cannot eat: ", err)
	}

	var name2 string = ""
	if _, err := Eat2(name2); err != nil {
		fmt.Println("cannot eat: ", err)
	}
}
