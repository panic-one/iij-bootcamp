package 
	shop

import (
	"fmt"
	"os"
	"./shop"
)

func Eat(name string) (bool, error) {
	if name == "" {
		return false, fmt.Errorf("name is empty")
	}
	fmt.Println(name)
	return true, nil
}

func main() {
	var name1 string = "GYUDON"
	if _, err := shop.Eat(name1); err != nil {
		fmt.Fprintf(os.Stderr, "cannot eat: '%s'\n" , err) //更新
	}

	var name2 string = ""
	if _, err := shop.Eat(name2); err != nil {
		fmt.Fprintf(os.Stderr, "cannot eat: '%s'\n" , err) //更新
	}
}