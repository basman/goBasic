# Hacking the calculator
## Introduction

In this exercise you are looking at existing code of a basic calculator.
Currently, it only supports arithmetic operations:

`+, -, *, /`

Test it out by running your own calculation:

`$ go run main.go '1 + 2 * 3'`

From the output you can see, the calculator honors precedence of * and / over + and -.

Parentheses are supported as well:

`$ go run main.go '(1 + 2) * 3'`

Run the unit tests in calc/evaluate_test.go. One unit test will fail, because the calculator
does not know the *bin(...)* function yet.

## Exercise: add and implement the *bin(...)* function

In this exercise you will extend the existing calculator to support conversion of decimal numbers
to binary.

The calculator uses a generated parser. The calculator's grammar is defined in **Expr.g4**.

### Step 1: extend the grammar in **Expr.g4**

Open **Expr.g4** and study the grammar definition. Be aware of the trailing **# labels** in some rules.
Try to duplicate and adapt the **# AddSub** line, so the operator **bin(...)** becomes part of the
language understood by the calculator.

Make sure the grammar rule you add ends with the label `# ToBinary`.

Based on file **Expr.g4**, ANTLR is used to generate a lexer and a parser.
You need antlr4 to be installed in order to generate a parser based on the grammar you just changed.
`$ brew install antlr`

The Makefile has all required options for the generator. Simply run **make** and the files in parser/
will be re-created:
   `$ make`

Only proceed if the generated file **parser/expr_listener.go** now lists `ExitToBinary()` and `EnterToBinary()`
as part of the interface `ExprListener`.
The smallest change in the grammar file has big effects on the generated code, so be careful to only
add one single line for the binary conversion and leave everything else as it was.

Ideally, you create a rule that is as powerful as the existing `# MulDiv` and `# AddSub` rules, for
the calculator to be able to evaluate all combinations with other expressions, such as:
```
bin(2)

bin(2 * 3)

1 + bin(2 - 1)

bin(4*2) + bin(2-1)

( bin(4*2) + bin(2 - 1) ) * 2

etc.
```
For this exercise it is good enough to convert the integer part of positive numbers. We will not care for
negative numbers.

### Step 2: use the new token from the parser

Open **calc/evaluate.go** and add a new method:

`func (l *exprListener) ExitToBinary(c *parser.ToBinaryContext)`

This method will be called automatically whenever the calculator finds *bin(...)* in the input.

The method *ExitToBinary* now is to be filled with useful code by you.
It needs to
1. get the last value from the stack using `l.pop()`, then
2. convert the value to its binary representation (e.g. `5` becomes `101`) and
3. finally `push` that value back onto the stack.

If you have never used a generated parser before, this may look very unusual. The method ExitToBinary
is part of the generated code on one hand and on the other hand is overridden by our custom `exprListener`
implementation. This allows us to fill the empty default implementation with instructions what to do with
the parsed input.

As the parser is executed in our `func Eval()` using
```
antlr.ParseTreeWalkerDefault.Walk(treeListener, p.Prog())
```
the treeListener's methods are called by the Walk method, each time the parser reads an entire rule
from the input.

### Step 3: rinse and repeat
Keep adapting the grammar and the `ExitToBinary` method until all unit tests succeed.
