# 🌳 Go Filter Library

Inspired by [filter](https://github.com/rwxrob/filter)
by [rwxrob](https://github.com/rwxrob).
This is a very rudimentary implementation.

## Installation

You can grab the latest binary [release](https://github.com/chriswifn/filter/releases).
This filter command can be installed as a standalone program or composed into
a Bonzai command tree.

Standalone
```
go install github.com/chriswifn/filter/cmd/filter@latest
```

Composed

```go
package z

import (
    Z "github.com/rwxrob/bonzai/z"
    "github.com/chriswifn/filter"
)

var Cmd = &Z.Cmd{
    Name: `z`,
    Commands: []*Z.Cmd{help.Cmd, filter.Cmd},
}
```

## Tab Completion
To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C filter filter
```

