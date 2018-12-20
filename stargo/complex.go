package stargo

import (
	"fmt"
	"reflect"

	"go.starlark.net/starlark"
	"go.starlark.net/syntax"
)

// A goComplex represents a Go value of kind complex.
// (Named complex types are represented by goNamedBasic.)
type goComplex struct {
	v reflect.Value // Kind=Complex; !CanAddr
}

var (
	_ Value               = goComplex{}
	_ starlark.Comparable = goComplex{}
	_ starlark.HasBinary  = goComplex{}
	_ starlark.HasUnary   = goComplex{} // - +
)

func (c goComplex) Attr(name string) (starlark.Value, error) { return method(c.v, name) }
func (c goComplex) AttrNames() []string                      { return methodNames(c.v) }
func (c goComplex) Freeze()                                  {}                                                   // immutable
func (c goComplex) Hash() (uint32, error)                    { return 0, fmt.Errorf("unhashable: %s", c.Type()) } // TODO
func (c goComplex) Reflect() reflect.Value                   { return c.v }
func (c goComplex) String() string                           { return str(c.v) }
func (c goComplex) Truth() starlark.Bool                     { return c.v.Complex() != 0 }
func (c goComplex) Type() string                             { return fmt.Sprintf("go.complex<%s>", c.v.Type()) }

func (x goComplex) CompareSameType(op syntax.Token, y starlark.Value, depth int) (bool, error) {
	// Different kinds of goComplex may compare equal, regardless
	// of width or "namedness". But a goComplex cannot compare
	// equal to any other number because it would be asymmetric.
	// TODO: Starlark may need a mechanism for comparing all kinds
	// of numbers symmetrically; see what Python does.
	xc := x.v.Complex()
	yc := y.(goComplex).v.Complex()
	switch op {
	case syntax.EQL:
		return xc == yc, nil
	case syntax.NEQ:
		return xc != yc, nil
	}
	return false, fmt.Errorf("invalid comparison: %s %s %s", x.Type(), op, y.Type())
}

func (c goComplex) Binary(op syntax.Token, y starlark.Value, side starlark.Side) (starlark.Value, error) {
	// TODO: + - * / // %. Operands must have same type.
	return nil, nil
}

func (c goComplex) Unary(op syntax.Token) (starlark.Value, error) {
	switch op {
	case syntax.PLUS:
		return c, nil
	case syntax.MINUS:
		return toStarlark(reflect.ValueOf(-c.v.Complex())), nil
	}
	return nil, nil
}

// -- builtins --

// complex(re, im)
func go٠complex(x, y float64) complex128 { return complex(x, y) }

// real(complex)
func go٠real(x complex128) float64 { return real(x) }

// imag(complex)
func go٠imag(x complex128) float64 { return imag(x) }
