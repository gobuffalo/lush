package quick

import "github.com/gobuffalo/lush/ast"

var (
	STRING  = NewString("hi")
	ARRAY   = NewArray(1, NewInteger(2), "3")
	IDENT   = NewIdent("foo")
	INT     = NewInteger(42)
	FLOAT   = NewFloat(3.14)
	ASSIGN  = NewAccess(IDENT, INT)
	VAR     = NewVar(IDENT, INT)
	ACCESS  = NewAccess(IDENT, 42)
	BLOCK   = NewBlock(ASSIGN, VAR)
	COMMENT = NewComment("i've got blisters on my fingers")
	MAP     = NewMap(map[ast.Statement]interface{}{
		IDENT: INT,
	})
	CALL = NewCall(IDENT, NewIdent("Bar"), ast.Statements{INT, FLOAT, STRING}, BLOCK)
	ELSE = NewElse(BLOCK)
)
