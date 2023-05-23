package comment_test

import (
	"fmt"
	"github.com/chriswifn/filter/comment"
)

func ExampleCommentTest() {
	text := `Test line
  Second test line
  `
	fmt.Println(comment.Comment(text))
	// Output:
	// # Test line
	// #   Second test line
}

func ExampleCommentTest2() {
	text := `Test line
  Second test line
  `
	fmt.Println(comment.Comment(text, `--`))
	// Output:
	// -- Test line
	// --   Second test line
}

func ExampleCommentTestEmptyLine() {
	text := `Test line

  Second test line
  `
	fmt.Println(comment.Comment(text, `--`))
	// Output:
	// -- Test line
	//
	// --   Second test line
}

func ExampleUnCommentTest() {
	text := `-- Test line
  -- Second test line
  `
	fmt.Println(comment.UnComment(text, `--`))
	// Output:
	// Test line
	//   Second test line
}

func ExampleHtitleTest() {
	text := `Das ist ein Test

Das ist auch ein Test`
	fmt.Println(comment.Htitle(text, `#`))
	// Output:
	// # ------------------------- Das ist ein Test -------------------------

	// # ----------------------- Das ist auch ein Test ----------------------

}
