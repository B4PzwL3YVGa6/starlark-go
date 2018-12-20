# Tests of pointers.

load("assert.star", "assert")

tarray = go.array_of(5, go.int)
assert.eq(str(tarray), '[5]int')

# TODO: eliminate need for def.
def f():
  p = go.new(tarray)
  assert.eq(type(p), 'go.ptr<*[5]int>')
  assert.true(p)
  for i in range(len(p)):
    for j in range(i, len(p)):
       p[j] += j+1
  assert.eq(list(p), [1, 4, 9, 16, 25])
f()

# A nil *go.array is an empty sequence.
p = go.ptr_to(tarray)() # nil
assert.eq(type(p), 'go.ptr<*[5]int>')
assert.eq(str(p), '<nil>')
assert.true(not p)
assert.ne(p, None) # nil pointer is not None
assert.eq(len(p), 0) # note: len((*[k]T)(nil)) == 0, not k like in Go
assert.eq(list(p), []) # empty
assert.fails(lambda: p[0], 'out of range')

# Pointers to non-array types have no len.
# However, statically they satisfy Iterable, Indexable, and HasSetIndex,
# behaving like an empty sequence, hence the surprising results below.
# To fix this, we would have to split Ptr into two types, Ptr+IndexablePtr,
# as we do with vars.
p = go.new(go.int)
assert.fails(lambda: len(p), 'go.ptr<\*int> has no len') # no len
assert.eq(list(p), []) # iterable!!
assert.eq([x for x in p], []) # iterable!!
def f(): p[0] = 1
assert.fails(f, 'index 0 out of range') # support set-index!!
intslice = go.slice_of(go.int)
# Even though it's iterable, you can't convert a nil *T to a Go []T:
assert.fails(lambda: intslice(p), 'cannot convert go.ptr<\*int> to Go \[\]int')
