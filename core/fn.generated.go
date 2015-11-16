package core

type Fn func(*Env, *Cons) (Value, error)
type Special func(*Env, *Cons) (Value, error)
type Macro func(*Env, *Cons) (Value, error)
