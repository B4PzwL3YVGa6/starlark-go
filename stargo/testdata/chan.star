# Tests of channels.

# TODO: improve unwieldy syntax for go.make_chan(go.chan_of(go.BothDir, T))
# BothDir is redundant.  However, permitting go.make_chan(T), where T
# is the elemnt type, would be# a mistake because it would allow no
# way to create a named channel type (though admittedly this is not a
# common need). Perhaps: go.make_chan_of(T)?

load("assert.star", "assert")

# chan type
sendch = go.chan_of(go.SendDir, go.string)
assert.eq('chan<- string', str(sendch))
assert.eq('go.type', type(sendch))
assert.true(not sendch()) # nil

# go.make_chan
assert.fails(lambda: go.make_chan(sendch, 1), 'unidirectional channel type')
assert.fails(lambda: go.make_chan(go.int, -10), 'non-chan type')

# go.cap
tch = go.chan_of(go.BothDir, go.string)
assert.eq('go.type', type(tch))
assert.eq(0, go.cap(go.make_chan(tch))) # unbuffered
ch = go.make_chan(tch, 10) # buffered
assert.eq(go.cap(ch), 10)
assert.fails(lambda: len(ch), 'has no len')

# send & recv (buffered)
assert.fails(lambda: go.send(ch, 1), 'send: cannot convert int to Go string')
assert.eq(go.send(ch, "hello"), None)
assert.eq(go.recv(ch), ("hello", True))
assert.eq(go.send(ch, "world"), None)
assert.eq(go.try_recv(ch), ("world", True))
assert.eq(go.try_recv(ch), (None, False))

# close
go.close(ch)
assert.eq(go.recv(ch), ("", False))
assert.eq(go.try_recv(ch), ("", False))
assert.fails(lambda: go.send(ch, "hello"), 'send: send on closed channel')
assert.fails(lambda: go.try_send(ch, "hello"), 'try_send: send on closed channel')
assert.fails(lambda: go.close(ch), 'close of closed channel')
assert.fails(lambda: go.close(tch()), 'close of nil channel')

# iteration
ch = go.make_chan(tch, 10)
go.send(ch, "one")
go.send(ch, "two")
go.send(ch, "three")
go.close(ch)
assert.eq([x for x in ch], ["one", "two", "three"])
