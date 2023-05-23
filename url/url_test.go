package url_test

import (
	"fmt"
	"github.com/chriswifn/filter/url"
)

func ExampleLinkExp() {
	fmt.Println(url.LinkExp.MatchString(`https://youtu.be/8st1NhaKDCA`))
}

func Example() {
	text := `
  This is a paragraph with some links:

  https://youtu.be/ownblock
  http://gist.github.com/Xeoncross/dd18aa4bbfba8e0f5d24a44604c911b7
  https://www.google.com/
  `
	fmt.Println(url.GetUrl(text))

	// Output:
	// https://youtu.be/ownblock
	// http://gist.github.com/Xeoncross/dd18aa4bbfba8e0f5d24a44604c911b7
	// https://www.google.com/

}

func QrExample() {
  text := `https://youtu.be/ownblock`
  fmt.Println(url.GetUrlQr(text))

  // Output
}
