###CALCULATOR

To build calculator in go, there is go standard libraries which evaluate an expression

```Code snippet```
```
import (
       "go/ast"
       "go/parser"
       "go/token"
   )
   
   func main() {
       fs := token.NewFileSet()
       tr, _ := parser.ParseExpr("(4-2) * 5")
       ast.Print(fs, tr)
   }
```

You could use this subpackages of the go package, but they might be overkill for your application.

This Application will help you to build simple calculator.