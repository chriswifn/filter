package filter

import (
	"github.com/chriswifn/filter/comment"
	"github.com/chriswifn/filter/link"
	"github.com/chriswifn/filter/url"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:      `filter`,
	Summary:   `Collection of some filters for UNIX`,
	Commands:  []*Z.Cmd{help.Cmd, link.Cmd, url.Cmd, comment.Cmd},
	Copyright: `Copyright 2023 Christian Hageloch`,
	Version:   `v0.1.0`,
	License:   `MIT`,
	Source:    `git@github.com:chriswifn/filter.git`,
	Issues:    `github.com/chriswifn/filter/issues`,
}
