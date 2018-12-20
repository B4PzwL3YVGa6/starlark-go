package stargo

import (
	"fmt"
	"reflect"
	"strings"

	"go.starlark.net/starlark"
	"go.starlark.net/syntax"
)

// A goNamedBasic represents a Go value of a named type whose
// underlying type is bool, number, or string.
//
// Values of the predeclared Go bool, string, and non-complex numeric
// types are represented by starlark.Bool, starlark.Int, and
// starlark.Float. Values of the predeclared Go complex types
// complex{64,128} are represented by goComplex.
//
// Named string types do not possess the methods of starlark.String as
// that would conflict with the named type's methods.
type goNamedBasic struct {
	v reflect.Value // Kind=Bool | String | Number; !CanAddr
}

var (
	_ Value               = goNamedBasic{}
	_ starlark.Comparable = goNamedBasic{}
	_ starlark.HasAttrs   = goNamedBasic{}
	_ starlark.HasBinary  = goNamedBasic{} // + - / // %
	_ starlark.HasUnary   = goNamedBasic{} // + -
)

func (b goNamedBasic) Attr(name string) (starlark.Value, error) { return method(b.v, name) }
func (b goNamedBasic) AttrNames() []string                      { return methodNames(b.v) }

func (b goNamedBasic) Hash() (uint32, error) {
	return 0, fmt.Errorf("unhashable: %s", b.Type()) // TODO
}

func (b goNamedBasic) Freeze()                {} // immutable
func (b goNamedBasic) Reflect() reflect.Value { return b.v }
func (b goNamedBasic) String() string         { return str(b.v) }
func (b goNamedBasic) Truth() starlark.Bool   { return isZero(b.v) == false }
func (b goNamedBasic) Type() string {
	// e.g. "go.uint<parser.Mode>".
	return fmt.Sprintf("go.%s<%s>", strings.ToLower(b.v.Kind().String()), b.v.Type())
}

func (x goNamedBasic) CompareSameType(op syntax.Token, y_ starlark.Value, depth int) (bool, error) {
	y := y_.(goNamedBasic)

	// Reject comparisons where the Go types are not equal,
	// just as the Go compiler does statically for x==y.
	// This restriction may prove too onerous for an untyped language.
	// Also, it is not consistent with Starlark's (mathematical)
	// equivalence relation for numbers.
	// TODO: more design required.
	if x.v.Type() != y.v.Type() {
		return false, fmt.Errorf("unsupported comparison %s %s %s", x.Type(), op, y.Type())
	}

	return x.v == y.v, nil
}

// TODO: support arithmetic on numeric primitives, so long as both
// operands have the same kind (e.g. int32) and/or type (e.g. Celsius).

func (b goNamedBasic) Binary(op syntax.Token, y starlark.Value, side starlark.Side) (starlark.Value, error) {
	// TODO: + - * / // %. Operands must have same kind and type.
	return nil, nil
}

func (b goNamedBasic) Unary(op syntax.Token) (starlark.Value, error) {
	switch op {
	case syntax.PLUS:
		return b, nil
	case syntax.MINUS:
		// TODO
	}
	return nil, nil
}
