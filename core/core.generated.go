//line ./core/core.gop:1
package core
//line ./core/core.gop:4

//line ./core/core.gop:3
func NewDefaultEnv() *Env {
	env := NewBlankEnv()

	// Specials
	env.LetSpecial("=", f_assign)
	env.LetSpecial("quote", f_quote)
	env.LetSpecial("fn", f_fn)
	env.LetSpecial("if", f_if)

	// I/O
	env.LetFn("prn", f_prn)

	// Types
	env.LetFn("type", f_type)

	// Control
	env.LetFn("catch", f_catch)
	env.LetFn("finally", f_finally)

	// Tags
	env.LetFn("tag", f_tag)
	env.LetFn("untag", f_untag)

	// Maths
	env.LetFn("+", f_add)
	env.LetFn("-", f_sub)
	env.LetFn("*", f_mul)
	env.LetFn("/", f_div)

	// Tables
	env.LetFn("table", f_table)

				// Constants
				env.Let("nil", nil)
				env.Let("t", Intern("t"))
//line ./core/core.gop:40

//line ./core/core.gop:39
	return env
}
