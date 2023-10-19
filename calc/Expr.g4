grammar Expr;
prog:	expr EOF ;
expr:	expr op=('*'|'/') expr # MulDiv
    |	expr op=('+'|'-') expr # AddSub
    |	NUMBER                 # Number
    |	'(' expr ')'           # Parens
    ;
WHITESPACE: [ \r\n\t]+ -> skip;
NUMBER     : [-]?[0-9]+ ('.' [0-9]+)?;
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
PAROPEN: '(';
PARCLOSE: ')';
