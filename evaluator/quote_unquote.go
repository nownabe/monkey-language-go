package evaluator

import (
	"github.com/nownabe/monkey-language-go/ast"
	"github.com/nownabe/monkey-language-go/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
