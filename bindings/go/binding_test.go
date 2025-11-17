package tree_sitter_javascript_test

import (
	"testing"

	tree_sitter_javascript "github.com/go-tree-sitter/tree-sitter-javascript/bindings/go"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_javascript.Language())
	if language == nil {
		t.Errorf("Error loading JavaScript grammar")
	}
}
