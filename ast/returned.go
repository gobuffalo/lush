package ast

import (
	"fmt"

	"github.com/gobuffalo/lush/faces"
)

type Returned struct {
	Value interface{}
	err   error
}

func (r Returned) Err() error {
	return r.err
}

func (r Returned) String() string {
	if st, ok := r.Value.(fmt.Stringer); ok {
		return st.String()
	}
	return fmt.Sprint(r.Value)
}

func NewReturned(i interface{}) (ret Returned) {
	defer func() {
		if err, ok := ret.Value.(error); ok {
			ret.err = err
		}
	}()
	if r, ok := i.(Returned); ok {
		return r
	}
	if i == nil {
		return Returned{}
	}

	switch t := i.(type) {
	case Returned:
		return t
	case nil:
		return Returned{}
	case interfacer:
		return Returned{Value: t.Interface()}
	case []interface{}:
		ii := t
		if len(ii) == 0 {
			return Returned{}
		}
		if len(ii) == 1 {
			i = ii[0]
			if r, ok := i.(Returned); ok {
				return r
			}
			if ri, ok := i.(interfacer); ok {
				i = ri.Interface()
			}
			if ri, ok := i.([]interface{}); ok {
				return NewReturned(ri)
			}
			return Returned{Value: i}
		}
		return Returned{Value: ii}
	case faces.Slice:
		ii := t.Slice()
		if len(ii) == 0 {
			return Returned{}
		}
		if len(ii) == 1 {
			i = ii[0]
			if r, ok := i.(Returned); ok {
				return r
			}
			if ri, ok := i.(interfacer); ok {
				i = ri.Interface()
			}
			if ri, ok := i.([]interface{}); ok {
				return NewReturned(ri)
			}
			return Returned{Value: i}
		}
		return Returned{Value: ii}
	}

	return Returned{Value: i}
}
