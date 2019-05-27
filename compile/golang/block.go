package golang

import (
	"fmt"

	"github.com/gobuffalo/lush/ast"
)

func (c Compiler) astBlock(b *ast.Block) error {
	if b == nil {
		return nil
	}

	if len(b.Statements) > 0 {
		fmt.Fprintln(c)
		// x := b.Statements.GoString()
		// x = strings.TrimSpace(x)
		// scan := bufio.NewScanner(strings.NewReader(x))
		// for scan.Scan() {
		// 	s := scan.Text()
		// 	if len(strings.TrimSpace(s)) == 0 {
		// 		fmt.Fprintf(c, "\n")
		// 		continue
		// 	}
		// 	fmt.Fprintf(c, fmt.Sprintf("\t%s\n", s))
		// }
	}
	return c.astStatements(b.Statements)
}
