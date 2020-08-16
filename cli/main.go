package main

import (
	"fmt"
	"log"

	"github.com/ryomak/qrme/cli/src"
	"github.com/ryomak/qrme/logic"
)

func main() {
	me, err := logic.NewMeBuilder().NickName("rymak").Introduction("暇です").Build()
	if err != nil {
		log.Fatal(err)
	}
	if err := src.Create(me); err != nil {
		log.Fatal(err)
	}
	fmt.Println(me.GetWebURL())
}
