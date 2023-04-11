package evaluator

import (
	"testing"

	"github.com/nownabe/monkey-language-go/ast"
	"github.com/nownabe/monkey-language-go/lexer"
	"github.com/nownabe/monkey-language-go/object"
	"github.com/nownabe/monkey-language-go/parser"
)

func TestDefineMacros(t *testing.T) {
	input := `
	let number = 1;
	let function = fn(x, y) { x + y };
	let mymacro = macro(x, y) { x + y };`

	env := object.NewEnvironment()
	program := testParseProgram(input)

	DefineMacros(program, env)

	if len(program.Statements) != 2 {
		t.Fatalf("Wrong number of statements. got=%d", len(program.Statements))
	}

	if _, ok := env.Get("number"); ok {
		t.Fatalf("number should not be defined")
	}

	if _, ok := env.Get("function"); ok {
		t.Fatalf("function should not be defnined")
	}

	obj, ok := env.Get("mymacro")
	if !ok {
		t.Fatalf("macro not in invironment")
	}

	macro, ok := obj.(*object.Macro)
	if !ok {
		t.Fatalf("object is not Macro. got=%T (%+v)", obj, obj)
	}

	if len(macro.Parameters) != 2 {
		t.Fatalf("wrong number of macro parameters. got=%d", len(macro.Parameters))
	}

	if macro.Parameters[0].String() != "x" {
		t.Errorf("parameter is not 'x'. got=%q", macro.Parameters[0])
	}

	if macro.Parameters[1].String() != "y" {
		t.Errorf("parameter is not 'y'. got=%q", macro.Parameters[1])
	}

	expectedBody := "(x + y)"

	if macro.Body.String() != expectedBody {
		t.Errorf("body is not %q. got=%q", expectedBody, macro.Body.String())
	}
}

func testParseProgram(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}
