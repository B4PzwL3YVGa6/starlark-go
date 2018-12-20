package stargo

import (
	"fmt"
	"reflect"
	"sort"

	"go.starlark.net/starlark"
	"go.starlark.net/syntax"
)

// A goPtr represents a Go value of kind pointer.
//
// If the variable it points to contains a struct,
// the struct's fields may be accessed or updated:
// p.f returns the value of that field, and
// p.f = x updates it.
// Similarly, a pointer to an array allows elements to be updated:
// p[i] = 1 or p[i]+=1.
//
// Starlark has no unary * operator: it would be ambiguous in f(*args).
// Use the go.deref builtin to load the value from the pointed-to variable.
//
// goPtrs are comparable, but they do not compare equal to None as that
// would violate symmetry: None does not compare equal to any other
// value. To test whether a pointer is valid, use its truth-value.
type goPtr struct {
	v reflect.Value // Kind=Ptr; !CanAddr
}

var (
	_ Value                = goPtr{}
	_ starlark.Comparable  = goPtr{}
	_ starlark.HasAttrs    = goPtr{}
	_ starlark.HasSetField = goPtr{}
	_ starlark.HasSetIndex = goPtr{}
	_ starlark.Indexable   = goPtr{}
	_ starlark.Sequence    = goPtr{}
)

func (p goPtr) Attr(name string) (starlark.Value, error) {
	v := p.v

	// methods
	if m := v.MethodByName(name); m.IsValid() {
		return goFunc{m}, nil
	}

	// struct fields (including promoted ones)
	if v.Type().Elem().Kind() == reflect.Struct {
		if v.IsNil() {
			return nil, fmt.Errorf("nil dereference")
		}
		if f := v.Elem().FieldByName(name); f.IsValid() {
			if !f.CanInterface() {
				return nil, fmt.Errorf("access to unexported field .%s", name)
			}
			return varOf(f), nil // alias
		}
	}
	return nil, nil
}

func (p goPtr) Index(i int) starlark.Value {
	if p.v.Type().Elem().Kind() != reflect.Array {
		panic(fmt.Sprintf("cannot index %s", p.Type()))
	}
	return varOf(p.v.Elem().Index(i)) // alias
}

func (p goPtr) Len() int {
	if elem := p.v.Type().Elem(); elem.Kind() == reflect.Array {
		// In Go, len on a nil *array succeeds and returns the
		// array length, but in stargo we must treat len of a
		// nil *array pointer as zero; otherwise, the Index
		// method would be called, and it has no way to
		// indicate an error.
		//
		// (The Index method cannot be easily changed to
		// return an error because then the Iterator interface
		// would also need a way to report errors that occur
		// during iteration.)
		if p.v.IsNil() {
			return 0
		}
		return elem.Len()
	}
	return -1
}

func (p goPtr) Iterate() starlark.Iterator {
	if elem := p.v.Type().Elem(); elem.Kind() == reflect.Array && !p.v.IsNil() {
		return &indexIterator{p.v.Elem(), 0}
	}
	return starlark.Tuple(nil).Iterate() // empty iterator
}

func (p goPtr) SetIndex(i int, y starlark.Value) error {
	if p.v.IsNil() {
		return fmt.Errorf("nil dereference")
	}
	array := p.v.Elem()
	if array.Kind() != reflect.Array {
		return fmt.Errorf("can't set index of %s", array.Type())
	}
	elem := array.Index(i)
	x, err := toGo(y, elem.Type())
	if err != nil {
		return err
	}
	if !elem.CanSet() {
		return fmt.Errorf("can't set index of %s", array.Type())
	}
	elem.Set(x)
	return nil
}

func (p goPtr) AttrNames() []string {
	t := p.v.Type()

	var names []string
	names = appendMethodNames(names, t)
	names = appendFieldNames(names, t.Elem())
	sort.Strings(names)
	return names
}

func (p goPtr) SetField(name string, val starlark.Value) error {
	if p.v.IsNil() {
		return fmt.Errorf("nil dereference")
	}
	if elem := p.v.Elem(); elem.Kind() == reflect.Struct {
		if f := elem.FieldByName(name); f.CanSet() {
			x, err := toGo(val, f.Type())
			if err != nil {
				return err
			}
			f.Set(x)
			return nil
		}
	}
	return fmt.Errorf("can't set .%s field of %s", name, p.Type())
}

func (p goPtr) Freeze()                {} // unimplementable
func (p goPtr) Hash() (uint32, error)  { return ptrHash(p.v), nil }
func (p goPtr) Reflect() reflect.Value { return p.v }
func (p goPtr) String() string         { return str(p.v) }
func (p goPtr) Truth() starlark.Bool   { return p.v.IsNil() == false }
func (p goPtr) Type() string           { return fmt.Sprintf("go.ptr<%s>", p.v.Type()) }

func (x goPtr) CompareSameType(op syntax.Token, y_ starlark.Value, depth int) (bool, error) {
	y := y_.(goPtr)
	switch op {
	case syntax.EQL:
		return ptrsEqual(x.v, y.v), nil
	case syntax.NEQ:
		return !ptrsEqual(x.v, y.v), nil
	}
	return false, fmt.Errorf("invalid comparison: %s %s %s", x.Type(), op, y.Type())
}

// A goUnsafePointer represents a Go unsafe.Pointer value.
//
// goUnsafePointers are comparable. However, they do not compare equal to
// None (since that would violate symmetry). To test whether a pointer
// is valid, use bool(ptr).
type goUnsafePointer struct {
	v reflect.Value // Kind=UnsafePointer; !CanAddr
}

var (
	_ starlark.Comparable = goUnsafePointer{}
	_ Value               = goUnsafePointer{}
)

func (p goUnsafePointer) Freeze()                {} // immutable
func (p goUnsafePointer) Hash() (uint32, error)  { return ptrHash(p.v), nil }
func (p goUnsafePointer) Reflect() reflect.Value { return p.v }
func (p goUnsafePointer) String() string         { return str(p.v) }
func (p goUnsafePointer) Truth() starlark.Bool   { return isZero(p.v) == false }
func (p goUnsafePointer) Type() string           { return fmt.Sprintf("go.unsafepointer<%s>", p.v.Type()) }

func (x goUnsafePointer) CompareSameType(op syntax.Token, y_ starlark.Value, depth int) (bool, error) {
	y := y_.(goUnsafePointer)
	switch op {
	case syntax.EQL:
		return ptrsEqual(x.v, y.v), nil
	case syntax.NEQ:
		return !ptrsEqual(x.v, y.v), nil
	}
	return false, fmt.Errorf("invalid comparison: %s %s %s", x.Type(), op, y.Type())
}

// Hash and Equal functions for chan, func, ptr, unsafepointer.
// Pointers, even of different types, compare equal if they point to the same object;
// reflect.Value equality does not have this property.
func ptrHash(x reflect.Value) uint32    { return uint32(x.Pointer()) }
func ptrsEqual(x, y reflect.Value) bool { return x.Pointer() == y.Pointer() }

// -- builtins --

// new(type)
func go٠new(t reflect.Type) interface{} { return reflect.New(t).Interface() }

// deref(ptr) returns *ptr (as an rvalue).
func go٠deref(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (_ starlark.Value, err error) {
	var ptr goPtr
	if err := starlark.UnpackPositionalArgs(b.Name(), args, kwargs, 1, &ptr); err != nil {
		return nil, err
	}
	if ptr.v.IsNil() {
		return nil, fmt.Errorf("deref: nil pointer")
	}
	return toStarlark(copyVal(ptr.v.Elem())), nil
}
