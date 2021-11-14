package main

import (
	"fmt"
	base26 "github.com/linexjlin/AZ-Base26"
)

func main() {
	b26 := base26.NewBase26("A")
	for i := 0; i < 200; i++ {
		fmt.Println(b26.String())
		b26.Add(1)
	}
}
