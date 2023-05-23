package url

import (
	_ "embed"
	"fmt"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:      `url`,
	Aliases:   []string{`url`},
	Summary:   `url filters`,
	Copyright: `Copyright 2023 Christian Hageloch`,
	Version:   `v0.1.0`,
	License:   `MIT`,
	Site:      `tba`,
	Source:    `git@github.com:chriswifn/filter.git`,
	Issues:    `github.com/chriswifn/filter/issues`,
	Commands:  []*Z.Cmd{help.Cmd, UrlCmd, UrlQrCmd},
}

var UrlCmd = &Z.Cmd{
	Name:      `get`,
	Aliases:   []string{`get`},
	Usage:     `[help|PATH]`,
	Summary:   `add pandoc url link syntax to url`,
	Copyright: `Copyright 2023 Christian Hageloch`,
	Version:   `v0.1.0`,
	License:   `MIT`,
	Source:    `git@github.com:chriswifn/filter.git`,
	Issues:    `github.com/chriswifn/filter/issues`,
	Commands:  []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		buf := Z.ArgsOrIn(args)
		fmt.Print(GetUrl(buf))
		return nil
	},
}

var UrlQrCmd = &Z.Cmd{
	Name:      `qr`,
	Aliases:   []string{`get`},
	Usage:     `[help|PATH]`,
	Summary:   `add pandoc url link syntax to url`,
	Copyright: `Copyright 2023 Christian Hageloch`,
	Version:   `v0.1.0`,
	License:   `MIT`,
	Source:    `git@github.com:chriswifn/filter.git`,
	Issues:    `github.com/chriswifn/filter/issues`,
	Commands:  []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		buf := Z.ArgsOrIn(args)
		fmt.Print(GetUrlQr(buf))
		return nil
	},
}
