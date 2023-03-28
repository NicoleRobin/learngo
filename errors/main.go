package main

import (
	"errors"
	"fmt"

	xerrors "github.com/pkg/errors"
)

// var ErrPermission = errors.New("permission denied")
var ErrPermission = errors.New("permission denied")

func DoSomething() error {
	// return fmt.Errorf("%w", ErrPermission)
	return xerrors.Wrapf(ErrPermission, "DoSomething failed")
}

func main() {
	if err := DoSomething(); err != nil && errors.Is(err, ErrPermission) {
		fmt.Printf("%+v\n", err)
	}
}
