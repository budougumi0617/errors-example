// Copyright © 2018 budougumi0617 All Rights Reserved.

package sub

import (
	"github.com/budougumi0617/errors-example/errors"
)

// Sub is struct in nested pkg.
type Sub struct {
}

// ReturnError returns custom error.
func (s *Sub) ReturnError() error {
	return errors.New("custom error")
}

// ReturnErrorByPkg returns error by pkg/errors.
func (s *Sub) ReturnErrorByPkg() error {
	return errors.NewByPkg("pkg/errors")
}
