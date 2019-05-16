package ast

import (
	"fmt"
	"strconv"
)

type Float float64

func (d Float) Interface() interface{} {
	return float64(d)
}

func (n Float) String() string {
	return fmt.Sprint(float64(n))
}

func (n Float) Exec(c *Context) (interface{}, error) {
	return float64(n), nil
}

func NewFloat(b []byte) (Float, error) {
	f, err := strconv.ParseFloat(string(b), 64)
	ft := Float(f)
	if err != nil {
		return ft, err
	}
	return ft, nil
}
