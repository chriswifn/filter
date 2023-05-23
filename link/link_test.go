package link_test

import (
	"fmt"
	"github.com/chriswifn/filter/link"
)

func ExampleLinkExp() {
	fmt.Println(link.LinkExp.MatchString(`https://youtu.be/8st1NhaKDCA`))
	fmt.Println(link.LinkExp.MatchString(`https://youtube/8st1NhaKDCA`))
	fmt.Println(link.LinkExp.MatchString(`https://youtube/`))
	fmt.Println(link.LinkExp.MatchString(`Test`))
	fmt.Println(link.LinkExp.MatchString(`another Test`))
	// Output:
	// true
	// true
	// true
	// false
	// false
}

func ExampleLink() {
	text := `
  This is a paragraph with https://youtu.be/inparagraph in it.

  * https://youtu.be/inlist

  https://youtu.be/ownblock
  http://gist.github.com/Xeoncross/dd18aa4bbfba8e0f5d24a44604c911b7
  https://www.google.com/
  `
	fmt.Printf("%q\n", link.LinkifyUrl(text))
	// Output:
	// "\n  This is a paragraph with 🌐 <https://youtu.be/inparagraph> in it.\n\n  * 🌐 <https://youtu.be/inlist>\n\n  🌐 <https://youtu.be/ownblock>\n  🌐 <http://gist.github.com/Xeoncross/dd18aa4bbfba8e0f5d24a44604c911b7>\n  🌐 <https://www.google.com/>\n  "
}

func ExampleLink_ignore_Angled() {
	text := `
  This is a paragraph with <https://youtu.be/inpara> in it.
  `
	fmt.Println(link.LinkifyUrl(text))
	// Output:
	// This is a paragraph with <https://youtu.be/inpara> in it.
}

func ExampleFileExp() {
	fmt.Println(link.FileExp.MatchString(`./tmp/test.txt`))
	fmt.Println(link.FileExp.MatchString(`../tmp/relative/path/to/test.txt`))
	fmt.Println(link.FileExp.MatchString(`./test.txt`))
	fmt.Println(link.FileExp.MatchString(`/absolute/path/wont/match/text.txt`))
	fmt.Println(link.FileExp.MatchString(`another Test`))
	// Output:
	// true
	// true
	// true
	// false
	// false
}

func ExampleFile() {
	text := `
  This is a paragraph with ./tmp/test.png in it.
  * This is the file in a list ./tmp/test.png
  ./tmp/test.png
  ../tmp/test.png
  ![test](./tmp/test.png)
  ![test](../tmp/test.png)
  `
	fmt.Printf("%q\n", link.LinkifyFile(text))
	// Output:
	// "\n  This is a paragraph with ![📁](./tmp/test.png) in it.\n  * This is the file in a list ![📁](./tmp/test.png)\n  ![📁](./tmp/test.png)\n  ![📁](../tmp/test.png)\n  ![test](./tmp/test.png)\n  ![test](../tmp/test.png)\n  "
}
