# Tests of strings.

load("go", "fmt")
load("assert.star", "assert")

assert.eq("[1, 2, 3]", fmt.Sprintf("%s", [1, 2, 3]))
assert.eq("*starlark.List", fmt.Sprintf("%T", [1, 2, 3]))
assert.eq("1 2 3", fmt.Sprintf("%d %d %d", *[1, 2, 3]))
