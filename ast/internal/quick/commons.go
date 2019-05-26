package quick

import "github.com/gobuffalo/lush/ast"

var (
	STRING  = NewString("hi")
	ARRAY   = NewArray(1, NewInteger(2), "3")
	IDENT   = NewIdent("foo")
	INT     = NewInteger(42)
	FLOAT   = NewFloat(3.14)
	ASSIGN  = NewAssign(IDENT, INT)
	VAR     = NewVar(IDENT, INT)
	ACCESS  = NewAccess(IDENT, 42)
	BLOCK   = NewBlock(ASSIGN, VAR)
	COMMENT = NewComment("i've got blisters on my fingers")
	MAP     = NewMap(map[string]interface{}{
		IDENT.String(): INT,
	})
	CALL   = NewCall(IDENT, NewIdent("Bar"), ast.Statements{INT, FLOAT, STRING}, BLOCK)
	IF     = NewIf(nil, ast.True, BLOCK, nil)
	ELSE   = NewElse(BLOCK)
	ELSEIF = NewElseIf(IF)
	FOR    = NewFor(ARRAY, []interface{}{NewIdent("i"), NewIdent("n")}, BLOCK)
	RANGE  = NewRange(ARRAY, []interface{}{NewIdent("i"), NewIdent("n")}, BLOCK)
	FUNC   = NewFunc([]interface{}{IDENT}, BLOCK)
	LET    = NewLet(IDENT, INT)
	OPEXPR = NewOpExpression(INT, "==", FLOAT)
	RETURN = NewReturn(INT)
	IMPORT = NewImport("fmt")
)
