package golem

func init() {
	Core.Function("<", func(env *Env, args []Value) (result Value) {
		CheckArgs("<", 1, -1, args)
		result = args[0]
		for _, arg := range args[1:] {
			switch a := result.(type) {
			case Int:
				switch b := arg.(type) {
				case Int:
					if !((a) < (b)) {
						return Nil
					}
				case Float:
					if !(Float(a) < (b)) {
						return Nil
					}
				default:
					panic("bad-type")
				}
			case Float:
				switch b := args[1].(type) {
				case Int:
					if !((a) < Float(b)) {
						return Nil
					}
				case Float:
					if !((a) < (b)) {
						return Nil
					}
				default:
					panic("bad-type")
				}
			default:
				panic("bad-type")
			}
			result = arg
		}
		return
	})
}
func init() {
	Core.Function(">", func(env *Env, args []Value) (result Value) {
		CheckArgs(">", 1, -1, args)
		result = args[0]
		for _, arg := range args[1:] {
			switch a := result.(type) {
			case Int:
				switch b := arg.(type) {
				case Int:
					if !((a) > (b)) {
						return Nil
					}
				case Float:
					if !(Float(a) > (b)) {
						return Nil
					}
				default:
					panic("bad-type")
				}
			case Float:
				switch b := args[1].(type) {
				case Int:
					if !((a) > Float(b)) {
						return Nil
					}
				case Float:
					if !((a) > (b)) {
						return Nil
					}
				default:
					panic("bad-type")
				}
			default:
				panic("bad-type")
			}
			result = arg
		}
		return
	})
}
func init() {
	Core.Function("<=", func(env *Env, args []Value) (result Value) {
		CheckArgs("<=", 1, -1, args)
		result = args[0]
		for _, arg := range args[1:] {
			switch a := result.(type) {
			case Int:
				switch b := arg.(type) {
				case Int:
					if !((a) <= (b)) {
						return Nil
					}
				case Float:
					if !(Float(a) <= (b)) {
						return Nil
					}
				default:
					panic("bad-type")
				}
			case Float:
				switch b := args[1].(type) {
				case Int:
					if !((a) <= Float(b)) {
						return Nil
					}
				case Float:
					if !((a) <= (b)) {
						return Nil
					}
				default:
					panic("bad-type")
				}
			default:
				panic("bad-type")
			}
			result = arg
		}
		return
	})
}
func init() {
	Core.Function(">=", func(env *Env, args []Value) (result Value) {
		CheckArgs(">=", 1, -1, args)
		result = args[0]
		for _, arg := range args[1:] {
			switch a := result.(type) {
			case Int:
				switch b := arg.(type) {
				case Int:
					if !((a) >= (b)) {
						return Nil
					}
				case Float:
					if !(Float(a) >= (b)) {
						return Nil
					}
				default:
					panic("bad-type")
				}
			case Float:
				switch b := args[1].(type) {
				case Int:
					if !((a) >= Float(b)) {
						return Nil
					}
				case Float:
					if !((a) >= (b)) {
						return Nil
					}
				default:
					panic("bad-type")
				}
			default:
				panic("bad-type")
			}
			result = arg
		}
		return
	})
}
func init() {
	Core.Function("==", func(env *Env, args []Value) (result Value) {
		CheckArgs("==", 1, -1, args)
		result = args[0]
		for _, arg := range args[1:] {
			switch a := result.(type) {
			case Int:
				switch b := arg.(type) {
				case Int:
					if !((a) == (b)) {
						return Nil
					}
				case Float:
					if !(Float(a) == (b)) {
						return Nil
					}
				default:
					panic("bad-type")
				}
			case Float:
				switch b := args[1].(type) {
				case Int:
					if !((a) == Float(b)) {
						return Nil
					}
				case Float:
					if !((a) == (b)) {
						return Nil
					}
				default:
					panic("bad-type")
				}
			default:
				panic("bad-type")
			}
			result = arg
		}
		return
	})
}
func init() {
	Core.Function("!=", func(env *Env, args []Value) (result Value) {
		CheckArgs("!=", 1, -1, args)
		result = args[0]
		for _, arg := range args[1:] {
			switch a := result.(type) {
			case Int:
				switch b := arg.(type) {
				case Int:
					if !((a) != (b)) {
						return Nil
					}
				case Float:
					if !(Float(a) != (b)) {
						return Nil
					}
				default:
					panic("bad-type")
				}
			case Float:
				switch b := args[1].(type) {
				case Int:
					if !((a) != Float(b)) {
						return Nil
					}
				case Float:
					if !((a) != (b)) {
						return Nil
					}
				default:
					panic("bad-type")
				}
			default:
				panic("bad-type")
			}
			result = arg
		}
		return
	})
}
