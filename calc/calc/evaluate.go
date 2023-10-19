package calc

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"

	"calculator/parser"
)

type exprListener struct {
	*parser.BaseExprListener

	stack []float64
}

func (l *exprListener) ExitMulDiv(c *parser.MulDivContext) {
	right, left := l.pop(), l.pop()
	op := c.GetOp()
	switch op.GetTokenType() {
	case parser.ExprLexerMUL:
		l.push(left * right)
	case parser.ExprLexerDIV:
		l.push(left / right)
	default:
		panic(fmt.Errorf("unexpected MulDiv op token type: %v(%v)", op.GetTokenType(), c.GetParser().GetSymbolicNames()[op.GetTokenType()]))
	}
}

func (l *exprListener) ExitAddSub(c *parser.AddSubContext) {
	right, left := l.pop(), l.pop()
	op := c.GetOp()
	switch op.GetTokenType() {
	case parser.ExprLexerADD:
		l.push(left + right)
	case parser.ExprLexerSUB:
		l.push(left - right)
	default:
		panic(fmt.Errorf("unexpected AddSub op token type: %v(%v)", op.GetTokenType(), c.GetParser().GetSymbolicNames()[op.GetTokenType()]))
	}
}

func (l *exprListener) ExitNumber(c *parser.NumberContext) {
	v, err := strconv.ParseFloat(c.GetText(), 64)
	if err != nil {
		panic(fmt.Errorf("number conversion of '%v' failed: %v", c.GetText(), err))
	}

	l.push(v)
}

func (l *exprListener) push(v float64) {
	//fmt.Println("(push", v, ")")
	l.stack = append(l.stack, v)
}

func (l *exprListener) pop() float64 {
	if len(l.stack) == 0 {
		panic("pop failed, stack is empty")
	}

	v := l.stack[len(l.stack)-1]
	l.stack = l.stack[:len(l.stack)-1]

	//fmt.Println("(pop", v, ")")
	return v
}

func Eval(input string) (float64, error) {
	is := antlr.NewInputStream(input)
	lexer := parser.NewExprLexer(is)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewExprParser(stream)

	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	treeListener := &exprListener{}
	antlr.ParseTreeWalkerDefault.Walk(treeListener, p.Prog())

	return treeListener.pop(), nil
}
