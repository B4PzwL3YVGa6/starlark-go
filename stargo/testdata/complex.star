load("assert.star", "assert")

# TODO: implement complex arithmetic.
# TODO: test comparison

# complex built-in function
assert.eq('go.func<func(float64, float64) complex128>', type(go.complex))
assert.eq('go.starlark.net/stargo.go٠complex', str(go.complex))
# assert.eq(1, go.complex(1, 2)) # TODO: convert int to float
assert.eq('(1+2i)', str(go.complex(1.0, 2.0)))
assert.eq('go.complex<complex128>', type(go.complex(1.0, 2.0)))
assert.eq(go.complex128, go.typeof(go.complex(1.0, 2.0)))

# complex128 type
assert.eq('go.type', type(go.complex128))
assert.eq('(0+0i)', str(go.complex128())) # zero value
assert.eq('(0+0i)', str(go.complex128(go.complex64()))) # conversion

c = go.complex(4.0, 5.0)
assert.eq('(4+5i)', str(c))
assert.eq('(-4-5i)', str(-c))

# real + imag
assert.eq(go.real(c), 4)
assert.fails(lambda: go.real(1.0), 'cannot convert float to Go complex128')
assert.eq(go.imag(c), 5)
assert.eq('float', type(go.imag(c)))
assert.eq(go.real(go.complex64(c)), 4)

# equality
assert.eq(c, c)
assert.eq(go.complex(4.0, 5.0), c)
# assert.eq(go.complex(4, 5), c) # TODO: can't convert int to float

# complex values are no ordered
assert.fails(lambda: c < c, 'invalid comparison')




