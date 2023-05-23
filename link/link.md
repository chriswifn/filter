The {{cmd .Name}} command replaces any apparent relative file path with
that same relative file path but wrapped in pandoc syntax.
If the relative file path is already in pandoc syntax it is skipped.

If no arguments are passed, assumes standard input.

If single argument is passed, assumes file name.
