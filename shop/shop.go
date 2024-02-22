package shop

import (
	"fmt"
	"time"
	"net/http"
)

type Gyudon struct {
	menu string
}

func NewGyudon() Gyudon {
	return Gyudon{
		menu: "NegitamaGyudon",
	}
}

func (self *Gyudon) Eat(w http.ResponseWriter, r *http.Request) { //引数をhttpdのセッション状態を受け取れるように追加
	if self.menu == "" {
		return
	}

	time.Sleep(time.Second * 10) //擬似食べてる時間
	fmt.Fprintf(w, "'%s'\n", self.menu) //食べた事を報告
	return
}
 
 