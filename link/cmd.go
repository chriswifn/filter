package link

import (
	_ "embed"
	"fmt"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
	"github.com/rwxrob/vars"
)

var (
	fileicon = "📁"
	urlicon  = "🌐"
)

//go:embed url.md
var urlDesc string

//go:embed link.md
var linkDesc string

func init() {
	Z.Conf.SoftInit()
	Z.Vars.SoftInit()
	Z.Dynamic[`durlicon`] = func() string { return urlicon }
	Z.Dynamic[`dfileicone`] = func() string { return fileicon }
}

var Cmd = &Z.Cmd{
	Name:      `link`,
	Aliases:   []string{`lk`},
	Summary:   `collection of a bunch of UNIX filters for links`,
	Copyright: `Copyright 2023 Christian Hageloch`,
	Version:   `v0.1.0`,
	License:   `MIT`,
	Source:    `git@github.com:chriswifn/filter.git`,
	Issues:    `github.com/chriswifn/filter/issues`,
	Commands:  []*Z.Cmd{help.Cmd, vars.Cmd, conf.Cmd, LinkUrlCmd, LinkFileCmd, InitCmd},
}

var InitCmd = &Z.Cmd{
	Name:     `init`,
	Commands: []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, _ ...string) error {
		val, _ := x.Caller.C(`fileicon`)
		if val == "null" {
			val = fileicon
		}
		x.Caller.Set(`fileicon`, val)

		val, _ = x.Caller.C(`urlicon`)
		if val == "null" {
			val = urlicon
		}
		x.Caller.Set(`urlicon`, val)
		return nil
	},
}

var LinkUrlCmd = &Z.Cmd{
	Name:        `url`,
	Aliases:     []string{`u`},
	Usage:       `[help|PATH]`,
	Summary:     `add pandoc url link syntax to url`,
	Copyright:   `Copyright 2023 Christian Hageloch`,
	Version:     `v0.1.0`,
	License:     `MIT`,
	Source:      `git@github.com:chriswifn/filter.git`,
	Issues:      `github.com/chriswifn/filter/issues`,
	Commands:    []*Z.Cmd{help.Cmd},
	Description: urlDesc,

	Call: func(x *Z.Cmd, args ...string) error {
		s, err := x.Caller.Get("urlicon")
		if err != nil {
			return err
		}
		buf := Z.ArgsOrIn(args)
		fmt.Print(LinkifyUrl(buf, s))
		return nil
	},
}

var LinkFileCmd = &Z.Cmd{
	Name:        `file`,
	Aliases:     []string{`f`},
	Usage:       `[help|PATH]`,
	Summary:     `add pandoc file link syntax to relative file path`,
	Copyright:   `Copyright 2023 Christian Hageloch`,
	Version:     `v0.1.0`,
	License:     `MIT`,
	Site:        `tba`,
	Source:      `git@github.com:chriswifn/filter.git`,
	Issues:      `github.com/chriswifn/filter/issues`,
	Commands:    []*Z.Cmd{help.Cmd},
	Description: linkDesc,

	Call: func(x *Z.Cmd, args ...string) error {
		s, err := x.Caller.Get("fileicon")
		if err != nil {
			return err
		}
		buf := Z.ArgsOrIn(args)
		fmt.Print(LinkifyFile(buf, s))
		return nil
	},
}
