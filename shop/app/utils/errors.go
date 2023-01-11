package utils

import (
	"fmt"
	"net/http"
)

type Op string
type Kind int

type Error struct {
	Op   Op
	Kind Kind
	Err  error
}

func (e Error) Error() string {
	return fmt.Sprintf("op=%s kind=%v err=%s", e.Op, e.Kind, e.Err.Error())
}

func E(args ...interface{}) error {
	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case Op:
			e.Op = arg
		case Kind:
			e.Kind = arg
		case error:
			e.Err = arg
		default:
			panic("bad call to error constructor E")
		}
	}
	return e
}

const (
	KindBadRequest Kind = http.StatusBadRequest
	KindNotFound   Kind = http.StatusNotFound
	KindInternal   Kind = http.StatusInternalServerError
)
