package main

import (
	"fmt"

	"github.com/Troxsoft/Furia2D-Engine/engine"
)

func main() {
	for i := 0; i < 10000; i++ {

		fmt.Println(engine.RandomRange(0, 10000))
	}

}
