package shop

import (
	"fmt"
	"time"
)

type Gyudon struct {
	menu string
}

func NewGyudon() Gyudon { 
	return Gyudon{
		menu: "NegitamaGyudon",
	}
}

func (self *Gyudon) Eat() (bool, error) {
	if self.menu == "" {
		return false, fmt.Errorf("name is empty.")
	}

	time.Sleep(time.Second * 10) 
	fmt.Println(self.menu)
	return true, nil
}
 