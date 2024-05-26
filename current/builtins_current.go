package current

import (
	"github.com/refaktor/rye-front/current/ebitengine"
	"github.com/refaktor/rye-front/current/fyne"
	"github.com/refaktor/rye-front/current/fynegen"
	"github.com/refaktor/rye-front/current/webview"
	"github.com/refaktor/rye/env"
	"github.com/refaktor/rye/evaldo"
)

var Builtins_current = map[string]*env.Builtin{

	"get": { // *** currently a concept in testing ... for getting a code of a function, maybe same would be needed for context?
		Argsn: 1,
		Doc:   "Returns value of the word in context",
		Pure:  true,
		Fn: func(ps *env.ProgramState, arg0 env.Object, arg1 env.Object, arg2 env.Object, arg3 env.Object, arg4 env.Object) env.Object {
			switch w := arg0.(type) {
			case env.Word:
				object, found := ps.Ctx.Get(w.Index)
				if found {
					return object
				} else {
					return evaldo.MakeBuiltinError(ps, "Word not found in contexts	", "get")
				}
			case env.Opword:
				object, found := ps.Ctx.Get(w.Index)
				if found {
					return object
				} else {
					return evaldo.MakeBuiltinError(ps, "Word not found in contexts	", "get")
				}
			default:
				return evaldo.MakeArgError(ps, 1, []env.Type{env.WordType}, "set")
			}
		},
	},
}

func RegisterBuiltins(ps *env.ProgramState) {
	evaldo.RegisterBuiltins2(Builtins_current, ps, "current")
	evaldo.RegisterBuiltinsInContext(ebitengine.Builtins_ebitengine, ps, "ebitengine")
	evaldo.RegisterBuiltinsInContext(fyne.Builtins_fyne, ps, "fyne")
	evaldo.RegisterBuiltinsInContext(fynegen.Builtins_fynegen, ps, "fyne")
	evaldo.RegisterBuiltinsInContext(webview.Builtins_webview, ps, "webview")
}
