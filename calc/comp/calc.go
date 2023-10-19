package comp

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"

	"calculator/parser"
)

type exprListener struct {
	*parser.BaseExprListener

	stack []float64
}

func (l *exprListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println("EnterEveryRule", ctx.GetText())
}

func (l *exprListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	if ctx.GetStart().GetTokenType() == parser.ExprLexerNUMBER {
		l.push(v)
	}
	fmt.Println("ExitEveryRule", ctx.GetText())
}

/*
func evalOp(num1, op, num2 antlr.Token) float64 {
	n1, err := strconv.ParseFloat(num1.GetText(), 64)
	if err != nil {
		panic("number parsing failed (n1)")
	}

	n2, err := strconv.ParseFloat(num2.GetText(), 64)
	if err != nil {
		panic("number parsing failed (n2)")
	}

	switch op.GetTokenType() {
	case parser.ExprLexerADD:
		return newNumberToken(n1 + n2)
	}


}

func isOp(t antlr.Token) bool {
	typ := t.GetTokenType()
	if typ == parser.ExprLexerMUL ||
		typ == parser.ExprLexerDIV ||
		typ == parser.ExprLexerADD ||
		typ == parser.ExprLexerSUB {
		return true
	}

	return false
}

func isNumber(t antlr.Token) bool {
	return t.GetTokenType() == parser.ExprLexerNUMBER
}

func isOpen(t antlr.Token) bool {
	return t.GetTokenType() == parser.ExprLexerPAROPEN
}

func isClose(t antlr.Token) bool {
	return t.GetTokenType() == parser.ExprLexerPARCLOSE
}
*/

func (s *exprListener) push(v float64) {
	s.stack = append(s.stack, v)
}

func (s *exprListener) pop() float64 {
	if len(s.stack) == 0 {
		panic("pop failed, stack is empty")
	}

	v := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

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
