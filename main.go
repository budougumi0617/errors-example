// Copyright © 2018 budougumi0617 All Rights Reserved.

package main

import (
	"fmt"

	"github.com/budougumi0617/errors-example/root"
)

func main() {
	if err := root.FirstFunc(); err != nil {
		fmt.Printf("大雑把にスタックトレースを出力する独自error:\n%+v\n", err)
	}

	if err := root.FirstFuncWithPkg(); err != nil {
		fmt.Printf("pkg/errorserror:\n%+v\n", err)
	}
}
