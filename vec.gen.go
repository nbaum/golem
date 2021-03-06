package golem

import (
	"fmt"
)

type Vec []Value

func (v Vec) String() string {
	return fmt.Sprint([]Value(v))
}

func (v Vec) Type() Value {
	return Intern("array")
}

func (v Vec) First() Value {
	return v[0]
}

func (v Vec) Empty() bool {
	return len(v) == 0
}

func (v Vec) Length() int {
	return len(v)
}

func (v Vec) Take(i int, s Sequence) Sequence {
	return s.Prepend(v[:i]...)
}

func (v Vec) Drop(i int) Value {
	return Vec(v[i:])
}

func (v Vec) Prepend(d ...Value) Sequence {
	return Vec(append(d, v...))
}

func ToVec(v Value) Vec {
	a := []Value{}
again:
	switch w := v.(type) {
	case Vec:
		return append(a, w...)
	case *Cons:
		a = append(a, w.car)
		v = w.cdr
		goto again
	case NilType:
		return a
	default:
		panic(Signal("bad-type"))
	}
}

func init() {
	Core.Function("vector?", func(env *Env, args []Value) (result Value) {
		CheckArgs("vector?", 1, 1, args)
		if _, ok := args[0].(Vec); ok {
			return args[0]
		} else {
			return Nil
		}
	})
}

func init() {
	Core.Function("make-vector", func(env *Env, args []Value) (result Value) {
		CheckArgs("make-vector", 1, 1, args)
		len, ok := args[0].(Int)
		if !ok {
			panic("bad-type")
		}

		return Vec(make([]Value, int(len)))
	})
}
