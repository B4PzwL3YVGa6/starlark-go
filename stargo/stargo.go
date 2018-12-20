// Package stargo provides Starlark bindings for Go values, variables,
// and types, allowing you to write Starlark scripts that interact with
// the full richness of Go functions and data types.
//
// See the cmd/stargo subdirectory for examples.
//
// This was inspired by Nate Finch's work on Starlight.
//
// THIS IS ALL EXPERIMENTAL AND MAY CHANGE RADICALLY.
// In particular, the way in which Go packages are loaded
// from Starlark needs more thought.
//
package stargo

import (
	"reflect"

	"go.starlark.net/starlark"
)

// NOTES for docs (TODO):
//
// - beware: stargo undermines many Starlark properties:
//   programs may be nondeterministic (e.g. due to map iteration)
//   programs may block (e.g. due to channels);
//   programs may panic (though we attempt to catch them);
//   values are not frozen, so data races are possible;
///  the lvalue/rvalue concept becomes important.
//
// - lvalue and rvalue modes, &v, and *ptr and the
//   changes in the Starlark compiler and runtime.
//   TODO unary *x won't work: ambiguous wrt f(*args).
//   Use go.deref builtin. Shouldn't be needed often (?).
//
// - API clients should not expect Index and Attr to work as usual.
//   They must be followed by an UAMP or VALUE op.
//
// - types are first-class values, as if reflect.Type was built-in.
//   e.g. "bytes.Buffer" is the reflect.Type for it.
//   bytes.Buffer() instantiates one
//
// - anonymous fields are promoted. (This creates potential ambiguity
//   because names are unqualified so a type can have two methods or
//   fields called f.)
//
// - there is no reasonable interpretation we can give to Freeze, so I'm
//   completely ignoring it. Caveat usor.
//
// - when an expression retrieves a primitive value of a named type
//   (e.g. token.Pos, which is a uint), we must preserve the namedness
//   because it provides all the methods. (We don't just want to conver
//   it to a uint.)  However, the primitive wrappers do not (yet)
//   participate in arithmetic.  We could support that, but we should
//   require that both operands of binary ops (+, etc) have the same
//   type. Otherwise what does celcius + fahrenheit return?
//   Named string types will not have the methods of string.
//
// - don't use "if err" (truth value) because err is
//   the concrete type, which could be zero. Use err != None.
//
// - there is no go.interface. An interface is a type, and this language is untyped.
//   All Starlark reflect.Values are rvalues or pointers; there are no lvalues
//   that might have a kind of Interface.
//
// - slice a[i:j:k] operator intentionally unimplemented to avoid
//   confusion. (It would be trivial to implement on Slice and Array.)
//   We definitely don't want to use the starlark operator with
//   Go semantics. Use go.slice operator.
//
// - there are many checks done by the reflect package that we cannot
//   hope to replicate exactly, so you should assume that more obscure
//   mistakes can crash the program.  goFunc.CallInternal handles
//   them, but many of the builtin functions do not.

// A Value is a Starlark value that wraps a Go value using reflection.
type Value interface {
	starlark.Value
	Reflect() reflect.Value
}

// ValueOf returns a Starlark value corresponding to the Go value x.
func ValueOf(x interface{}) starlark.Value {
	return toStarlark(reflect.ValueOf(x))
}

// VarOf returns a Starlark value that wraps the Go variable *ptr.
// VarOf panics if ptr is not a pointer.
func VarOf(ptr interface{}) starlark.Variable {
	v := reflect.ValueOf(ptr)
	if v.Kind() != reflect.Ptr {
		panic("VarOf: not a pointer")
	}
	return varOf(v.Elem())
}

// TypeOf returns a Starlark value that represents the
// type T of the variable pointed to by *ptr.
// TypeOf panics if ptr is not a pointer.
// It is typically invoked as TypeOf(new(T)).
func TypeOf(ptr interface{}) Type {
	if ptr == nil {
		panic("TypeOf(nil)")
	}
	tptr := reflect.TypeOf(ptr)
	if tptr.Kind() != reflect.Ptr {
		panic("TypeOf: not a pointer")
	}
	return Type{tptr.Elem()}
}
