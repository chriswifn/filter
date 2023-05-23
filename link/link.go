package link

import (
	"github.com/rwxrob/to"
	"regexp"
)

var LinkExp = regexp.MustCompile(`([^<]|^)(https?:\/\/\S+)`)

// This will only work on relative paths which is good practice
// because absolute paths are useless when deploying or sending
// somewhere else that is not the local machine
var FileExp = regexp.MustCompile(`([^((\.]|^)(\.{1,2}\/\S+)`)

// Linkify converts any YouTube URL within the input into a valid
// markdown URL (wrapped with angle brackets) and prefixed with
// the optionally passed prefix (default is video)
// returns an empty string if no arguments are passed.
func LinkifyUrl(args ...any) string {
	var in, pre string
	ln := len(args)
	pre = "🌐"
	switch {
	case ln >= 2:
		pre = to.String(args[1])
		fallthrough
	case ln == 1:
		in = to.String(args[0])
	case ln == 0:
		return ""
	}
	return LinkExp.ReplaceAllString(in, `$1`+pre+` <$2>`)
}

// LinkifyFile converts any apparent file path into a pandoc compatible
// file link.
func LinkifyFile(args ...any) string {
	var in, pre string
	ln := len(args)
	pre = "📁"
	switch {
	case ln >= 2:
		pre = to.String(args[1])
		fallthrough
	case ln == 1:
		in = to.String(args[0])
	case ln == 0:
		return ""
	}
	return FileExp.ReplaceAllString(in, `$1`+`![`+pre+`]`+`($2)`)
}
