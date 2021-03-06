package golem

type Env struct {
	maps [NamespaceCount]map[*Sym]Value
	up   *Env
}

func NewEnv(up *Env) *Env {
	env := &Env{up: up}
	m := make(map[*Sym]Value)
	for i := 0; i < int(NamespaceCount); i++ {
		env.maps[i] = m
	}
	return env
}

func (e *Env) Get(ns Namespace, name *Sym) Value {
	m := e.maps[ns]
	if val, ok := m[name]; ok {
		return val
	} else if e.up != nil {
		return e.up.Get(ns, name)
	} else {
		panic(Signal("unbound", "name", name, "ns", NamespaceName(ns)))
	}
}

func (e *Env) Set(ns Namespace, name *Sym, value Value) {
	e2 := e
again:
	m := e2.maps[ns]
	if _, ok := m[name]; ok {
		m[name] = value
	} else if e2.up != nil {
		e2 = e2.up
		goto again
	} else {
		e.Bind(ns, name, value)
	}
}

func (e *Env) Bind(ns Namespace, name *Sym, v Value) {
	m := e.maps[ns]
	m[name] = v
}

func (e *Env) DestructuringBind(ns Namespace, places Value, values Value) {
	switch places := places.(type) {
	case *Sym:
		e.Bind(ns, places, values)
	case *Cons:
		if values, ok := values.(*Cons); ok {
			for {
				e.DestructuringBind(ns, places.car, values.car)
				if places.cdr == Nil {
					if values.cdr != Nil {
						panic(Signal("excess-values", "what", values))
					}
					break
				} else if next, ok := places.cdr.(*Cons); ok {
					places = next
				} else if next, ok := places.cdr.(*Sym); ok {
					e.DestructuringBind(ns, next, values.cdr)
					break
				} else {
					panic(Signal("bad-type", "expected", List(Intern("or"), Intern("cons"), Intern("sym")), "got", places.cdr))
				}
				if values.cdr == Nil {
					panic(Signal("excess-values", "what", values))
				} else if next, ok := values.cdr.(*Cons); ok {
					values = next
				} else {
					panic(Signal("bad-type", "expected", Intern("cons"), "got", values.cdr))
				}
			}
		} else {
			panic(Signal("bad-type", "expected", Intern("cons"), "got", values))
		}
	case NilType:
		if values != Nil {
			panic(Signal("excess-values", "what", places))
		}
	default:
		panic(Signal("bad-type", "expected", List(Intern("or"), Intern("cons"), Intern("sym")), "got", places))
	}
	return
}

func (e *Env) Function(name string, proc func(*Env, []Value) Value) {
	e.Bind(Functions, Intern(name), NewFn(Intern(name), proc))
}

func (e *Env) Special(name string, proc func(*Env, []Value) Value) {
	e.Bind(Functions, Intern(name), NewSpecial(Intern(name), proc))
}

func (e *Env) Variable(name string, value Value) {
	e.Bind(Variables, Intern(name), value)
}

func (e *Env) String() string {
	return "<env>"
}

func (e *Env) Type() Value {
	return Intern("env")
}

func init() {
	Core.Function("env?", func(env *Env, args []Value) (result Value) {
		CheckArgs("env?", 1, 1, args)
		if _, ok := args[0].(*Env); ok {
			return T
		} else {
			return Nil
		}
	})
}

func init() {
	Core.Function("env", func(env *Env, args []Value) (result Value) {
		CheckArgs("env", 1, 1, args)
		if args[0] == Nil {
			return NewEnv(nil)
		}
		arg0, ok := args[0].(*Env)
		e := arg0
		if !ok {
			panic("bad-type")
		}
		return NewEnv(e)
	})
}
