# GO Bug
This is a small repo to demonstrate a bug in cgo. The bug prevents C code that was compiled in go to be used in go.

## Instructions
```bash
$ make build && make run
# command-line-arguments
In file included from ./main.go:3:0:
cgo-builtin-prolog:7:48: error: conflicting types for '_GoString_'
cgo-builtin-prolog:6:44: note: previous declaration of '_GoString_' was here
In file included from ./main.go:3:0:
cgo-gcc-export-header-prolog:27:20: error: 'GoString' redeclared as different kind of symbol
cgo-builtin-prolog:8:12: note: previous declaration of 'GoString' was here
Makefile:9: recipe for target 'run' failed
make: *** [run] Error 2
```
