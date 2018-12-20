package stargo

import (
	"fmt"
	"reflect"
	"strings"

	"go.starlark.net/starlark"
	"go.starlark.net/syntax"
)

// A goChan represents a Go value of kind chan.
//
// It implements Iterable (but not Sequence).
type goChan struct {
	v reflect.Value // kind=Chan; !CanAddr
}

var (
	_ Value               = goChan{}
	_ starlark.Comparable = goChan{}
	_ starlark.HasAttrs   = goChan{}
	_ starlark.Iterable   = goChan{}
)

func (c goChan) Attr(name string) (starlark.Value, error) { return method(c.v, name) }
func (c goChan) AttrNames() []string                      { return methodNames(c.v) }
func (c goChan) Freeze()                                  {} // unimplementable
func (c goChan) Hash() (uint32, error)                    { return ptrHash(c.v), nil }
func (c goChan) Iterate() starlark.Iterator               { return chanIterator{c.v} }
func (c goChan) Reflect() reflect.Value                   { return c.v }
func (c goChan) String() string                           { return str(c.v) }
func (c goChan) Truth() starlark.Bool                     { return c.v.IsNil() == false }
func (c goChan) Type() string                             { return fmt.Sprintf("go.chan<%s>", c.v.Type()) }

func (x goChan) CompareSameType(op syntax.Token, y_ starlark.Value, depth int) (bool, error) {
	y := y_.(goChan)
	switch op {
	case syntax.EQL:
		return ptrsEqual(x.v, y.v), nil
	case syntax.NEQ:
		return !ptrsEqual(x.v, y.v), nil
	}
	return false, fmt.Errorf("invalid comparison: %s %s %s", x.Type(), op, y.Type())
}

type chanIterator struct{ v reflect.Value }

func (it chanIterator) Next(x *starlark.Value) bool {
	v, ok := it.v.TryRecv()
	if ok {
		*x = toStarlark(v)
	}
	return ok
}

func (it chanIterator) Done() {}

// make_chan(type, cap=0)
func go٠make_chan(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (res starlark.Value, err error) {
	var t Type
	cap := 0
	if err := starlark.UnpackPositionalArgs(b.Name(), args, kwargs, 1, &t, &cap); err != nil {
		return nil, err
	}
	err = protect(thread, b.Name(), func() {
		res = toStarlark(reflect.MakeChan(t.t, cap))
	})
	return res, err
}

// recv(ch), try_recv(ch)
func go٠recv(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (res starlark.Value, err error) {
	var ch goChan
	if err := starlark.UnpackPositionalArgs(b.Name(), args, kwargs, 1, &ch); err != nil {
		return nil, err
	}

	err = protect(thread, b.Name(), func() {
		var x reflect.Value
		var ok bool
		if strings.HasPrefix(b.Name(), "try_") {
			x, ok = ch.v.TryRecv() // x may be invalid
		} else {
			x, ok = ch.v.Recv()
		}
		res = starlark.Tuple{toStarlark(x), starlark.Bool(ok)}
	})
	return res, err
}

// send(ch, v), try_send(ch, v)
func go٠send(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (res starlark.Value, err error) {
	var ch goChan
	var v starlark.Value
	if err := starlark.UnpackPositionalArgs(b.Name(), args, kwargs, 2, &ch, &v); err != nil {
		return nil, err
	}
	x, err := toGo(v, ch.v.Type().Elem())
	if err != nil {
		return nil, fmt.Errorf("%s: %v", b.Name(), err)
	}
	err = protect(thread, b.Name(), func() {
		if strings.HasPrefix(b.Name(), "try_") {
			res = starlark.Bool(ch.v.TrySend(x))
		} else {
			ch.v.Send(x)
			res = starlark.None
		}
	})
	return res, err
}

// close(chan)
func go٠close(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var ch goChan
	if err := starlark.UnpackPositionalArgs(b.Name(), args, kwargs, 1, &ch); err != nil {
		return nil, err
	}
	return starlark.None, protect(thread, "close", ch.v.Close)
}
