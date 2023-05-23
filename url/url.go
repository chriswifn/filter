package url

import (
	"github.com/rwxrob/to"
	"github.com/skip2/go-qrcode"
	"log"
	"regexp"
	"strings"
)

var LinkExp = regexp.MustCompile(`(((http|https|gopher|gemini|ftp|ftps|git)://|www\\.)[a-zA-Z0-9.]*[:;a-zA-Z0-9./+@$&%?$\#=_~-]*)|((magnet:\\?xt=urn:btih:)[a-zA-Z0-9]*)`)

func GetUrl(args ...any) string {
	var in string
	ln := len(args)
	switch {
	case ln >= 2:
		fallthrough
	case ln == 1:
		in = to.String(args[0])
	case ln == 0:
		return ""
	}
	return strings.Replace(strings.Join(LinkExp.FindAllString(in, -1), " "), " ", "\n", -1)
}

func GetUrlQr(args ...any) string {
	var in string
	ln := len(args)
	switch {
	case ln >= 2:
		fallthrough
	case ln == 1:
		in = to.String(args[0])
	case ln == 0:
		return ""
	}

	str := strings.Replace(strings.Join(LinkExp.FindAllString(in, -1), " "), " ", "\n", -1)
	q, err := qrcode.New(str, qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	}
	return q.ToSmallString(false)
}
