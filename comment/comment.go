package comment

import (
	"github.com/rwxrob/to"
	"regexp"
	"strings"
)

// This is just here to filter input
// var LineExp = regexp.MustCompile(`.*\n`)
var LineExp = regexp.MustCompile(`(.*\S.*)`)

func Comment(args ...any) string {
	var in, comment string
	comment = "#"
	ln := len(args)
	switch {
	case ln >= 2:
		comment = to.String(args[1])
		fallthrough
	case ln == 1:
		in = to.String(args[0])
	case ln == 0:
		return ""
	}
	return LineExp.ReplaceAllString(in, comment+` `+`$1`)
}

func UnComment(args ...any) string {
	var in, comment string
	comment = "#"
	ln := len(args)
	switch {
	case ln >= 2:
		comment = to.String(args[1])
		fallthrough
	case ln == 1:
		in = to.String(args[0])
	case ln == 0:
		return ""
	}
	return strings.Replace(in, comment+" ", "", -1)
}

func Htitle(args ...any) string {
	var in, comment string
	var hrulewidth int
	hrulewidth = 72
	comment = "#"
	ln := len(args)
	switch {
	case ln >= 2:
		comment = to.String(args[1])
		fallthrough
	case ln == 1:
		in = to.String(args[0])
	case ln == 0:
		return ""
	}
	lines := strings.Split(in, "\n")
	var arr []string
	for _, line := range lines {
		length := len(line)
		side := (hrulewidth/2 - length/2) - 3
		left := side
		right := side
		if length%2 == 1 {
			right--
		}
		arr = append(arr, LineExp.ReplaceAllString(line, comment+` `+strings.Repeat(`-`, left)+` `+`$1`+` `+strings.Repeat(`-`, right)))
	}
	return strings.Join(arr, "\n")
}
