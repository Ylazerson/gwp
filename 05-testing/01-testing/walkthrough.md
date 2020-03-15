# B"H

```sh

# -- --------------------------------
go test

# -- --------------------------------
# -v    : verbose 
# -cover: show test coverage
go test -v -cover

# -- --------------------------------
# -short: skip "long running" tests
go test -v -cover -short

# -- --------------------------------
go test -v -cover -short -bench .

# -- --------------------------------
# Run just benchmark tests
go test -run x -bench .
```