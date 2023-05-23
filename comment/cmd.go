package comment

import (
	_ "embed"
	"fmt"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
	"github.com/rwxrob/vars"
)

var comment = "--"

//go:embed comment.md
var commentDesc string

//go:embed uncomment.md
var uncommentDesc string

func init() {
	Z.Conf.SoftInit()
	Z.Vars.SoftInit()
	Z.Dynamic[`dcomment`] = func() string { return comment }
}

var Cmd = &Z.Cmd{
	Name:      `comment`,
	Aliases:   []string{`com`},
	Usage:     `[help|PATH]`,
	Summary:   `Comment lines`,
	Copyright: `Copyright 2023 Christian Hageloch`,
	Version:   `v0.1.0`,
	License:   `MIT`,
	Source:    `git@github.com:chriswifn/filter.git`,
	Issues:    `github.com/chriswifn/filter/issues`,
	Commands:  []*Z.Cmd{help.Cmd, vars.Cmd, conf.Cmd, CommentCmd, UncommentCmd, HtitleCmd, InitCmd},
}

var InitCmd = &Z.Cmd{
	Name:     `init`,
	Commands: []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, _ ...string) error {
		val, _ := x.Caller.C(`comment`)
		if val == "null" {
			val = comment
		}
		x.Caller.Set(`comment`, val)
		return nil
	},
}

var CommentCmd = &Z.Cmd{
	Name:        `comment`,
	Aliases:     []string{`com`},
	Usage:       `[help|PATH]`,
	Summary:     `add comment string to beginning of line`,
	Copyright:   `2023 Christian Hageloch`,
	License:     `MIT`,
	Source:      `git@github.com:chriswifn/filter.git`,
	Issues:      `github.com/chriswifn/filter/issues`,
	Commands:    []*Z.Cmd{help.Cmd},
	Description: commentDesc,

	Call: func(x *Z.Cmd, args ...string) error {
		s, err := x.Caller.Get("comment")
		if err != nil {
			return err
		}
		buf := Z.ArgsOrIn(args)
		fmt.Print(Comment(buf, s))
		return nil
	},
}

var UncommentCmd = &Z.Cmd{
	Name:        `uncomment`,
	Aliases:     []string{`ucom`},
	Usage:       `[help|PATH]`,
	Summary:     `remove comment string from line`,
	Copyright:   `2023 Christian Hageloch`,
	License:     `MIT`,
	Commands:    []*Z.Cmd{help.Cmd},
	Source:      `git@github.com:chriswifn/filter.git`,
	Issues:      `github.com/chriswifn/filter/issues`,
	Description: uncommentDesc,

	Call: func(x *Z.Cmd, args ...string) error {
		s, err := x.Caller.Get("comment")
		if err != nil {
			return err
		}
		buf := Z.ArgsOrIn(args)
		fmt.Print(UnComment(buf, s))
		return nil
	},
}

var HtitleCmd = &Z.Cmd{
	Name:        `htitle`,
	Aliases:     []string{`title`},
	Usage:       `[help|PATH]`,
	Summary:     `add title comment`,
	Copyright:   `2023 Christian Hageloch`,
	License:     `MIT`,
	Commands:    []*Z.Cmd{help.Cmd},
	Source:      `git@github.com:chriswifn/filter.git`,
	Issues:      `github.com/chriswifn/filter/issues`,
	Description: uncommentDesc,

	Call: func(x *Z.Cmd, args ...string) error {
		s, err := x.Caller.Get("comment")
		if err != nil {
			return err
		}
		buf := Z.ArgsOrIn(args)
		fmt.Print(Htitle(buf, s))
		return nil
	},
}
