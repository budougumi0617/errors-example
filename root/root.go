// Copyright © 2018 budougumi0617 All Rights Reserved.

package root

import (
	"github.com/budougumi0617/errors-example/root/sub"
)

// FirstFunc calles method in sub package.
func FirstFunc() error {
	return secondFunc()
}

func secondFunc() error {
	s := &sub.Sub{}
	return s.ReturnError()
}

// FirstFuncWith calles method in sub package.
func FirstFuncWithPkg() error {
	return secondFuncWithPkg()
}

func secondFuncWithPkg() error {
	s := &sub.Sub{}
	return s.ReturnErrorByPkg()
}
