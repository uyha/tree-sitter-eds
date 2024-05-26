package tree_sitter_eds_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-eds"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_eds.Language())
	if language == nil {
		t.Errorf("Error loading Eds grammar")
	}
}
