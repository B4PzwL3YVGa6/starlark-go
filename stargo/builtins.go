package stargo

import (
	"reflect"

	"go.starlark.net/starlark"
)

// Builtins is a module, typically predeclared under the name "go",
// that provides access to all the Go builtin functions and types.
// Because reflect.Type is effectively a built-in type in stargo,
// many of the methods of reflect.Type are included here too.
//
var Builtins = &starlark.Module{
	Name: "go",
	Members: starlark.StringDict{
		// We use an Arabic zero '٠' (U+0660) in these function names.

		// built-in functions
		"cap":     starlark.NewBuiltin("cap", go٠cap),
		"close":   starlark.NewBuiltin("close", go٠close),
		"complex": ValueOf(go٠complex),
		"new":     ValueOf(go٠new),
		"panic":   ValueOf(go٠panic),
		"typeof":  ValueOf(reflect.TypeOf),

		// ptr
		"deref": starlark.NewBuiltin("deref", go٠deref), // deref(ptr) returns *ptr

		// map
		"make_map": starlark.NewBuiltin("make_map", go٠make_map),
		"delete":   starlark.NewBuiltin("delete", go٠delete),

		// slice
		"make_slice": starlark.NewBuiltin("make_slice", go٠make_slice),
		"append":     starlark.NewBuiltin("append", go٠append),
		"copy":       starlark.NewBuiltin("copy", go٠copy),
		"slice":      starlark.NewBuiltin("slice", go٠slice),

		// complex
		"real": ValueOf(go٠real),
		"imag": ValueOf(go٠imag),

		// chan
		"make_chan": starlark.NewBuiltin("make_chan", go٠make_chan),
		"send":      starlark.NewBuiltin("send", go٠send),
		"recv":      starlark.NewBuiltin("recv", go٠recv),
		"try_recv":  starlark.NewBuiltin("try_recv", go٠recv),
		"try_send":  starlark.NewBuiltin("try_send", go٠send),
		"ChanDir":   TypeOf(new(reflect.ChanDir)),
		"BothDir":   ValueOf(reflect.BothDir),
		"RecvDir":   ValueOf(reflect.RecvDir),
		"SendDir":   ValueOf(reflect.SendDir),

		// type constructors
		"map_of":   ValueOf(reflect.MapOf),
		"array_of": ValueOf(reflect.ArrayOf),
		"slice_of": ValueOf(reflect.SliceOf),
		"ptr_to":   ValueOf(reflect.PtrTo),
		"chan_of":  ValueOf(reflect.ChanOf),
		"func_of":  ValueOf(reflect.FuncOf),

		// built-in types
		"bool":       TypeOf(new(bool)),
		"int":        TypeOf(new(int)),
		"int8":       TypeOf(new(int8)),
		"int16":      TypeOf(new(int16)),
		"int32":      TypeOf(new(int32)),
		"int64":      TypeOf(new(int64)),
		"uint":       TypeOf(new(uint)),
		"uint8":      TypeOf(new(uint8)),
		"uint16":     TypeOf(new(uint16)),
		"uint32":     TypeOf(new(uint32)),
		"uint64":     TypeOf(new(uint64)),
		"uintptr":    TypeOf(new(uintptr)),
		"rune":       TypeOf(new(rune)),
		"byte":       TypeOf(new(byte)),
		"float32":    TypeOf(new(float32)),
		"float64":    TypeOf(new(float64)),
		"complex64":  TypeOf(new(complex64)),
		"complex128": TypeOf(new(complex128)),
		"string":     TypeOf(new(string)),
		"error":      TypeOf(new(error)),
	},
}

// panic(any)
func go٠panic(x interface{}) { panic(x) }
