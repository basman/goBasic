grammar Expr;
prog:	expr EOF ;
expr:	expr op=('*'|'/') expr
    |	expr op=('+'|'-') expr
    |	NUMBER
    |	'(' expr ')'
    ;
WHITESPACE: [ \r\n\t]+ -> skip;
NUMBER     : [0-9]+.?[0-9]*;
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
PAROPEN: '(';
PARCLOSE: ')';
