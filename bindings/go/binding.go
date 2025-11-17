package tree_sitter_javascript

// // #cgo CFLAGS: -std=c11 -fPIC
// // #include "../../src/parser.c"
// // #if __has_include("../../src/scanner.c")
// // #include "../../src/scanner.c"
// // #endif

/*
#cgo CFLAGS: -I${SRCDIR}/../../src/tree_sitter
#include "../../src/parser.c"
#if __has_include("../../src/scanner.c")
#include "../../src/scanner.c"
#endif
*/
import "C"

import (
	"unsafe"

	_ "github.com/go-tree-sitter/tree-sitter-javascript/src"
	_ "github.com/go-tree-sitter/tree-sitter-javascript/src/tree_sitter"
)

// Get the tree-sitter Language for this grammar.
func Language() unsafe.Pointer {
	return unsafe.Pointer(C.tree_sitter_javascript())
}
