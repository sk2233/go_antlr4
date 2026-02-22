grammar Expr;

// 程序由多个语句组成
program: statement+ ;

// 语句由很多种
statement : varStatement
        | returnStatement
        | blockStatement
        | ifStatement
        | forStatement
        | breakStatement
        | continueStatement
        | expresstionStatement
        ;

varStatement: 'var' IDENTIFIER '=' expression ';' ;
returnStatement: 'return' expression ';' ;
blockStatement: '{' statement* '}' ;
ifStatement: 'if' '(' cond=expression ')' ifBody = blockStatement ('else' elseBody=blockStatement)? ;
forStatement: 'for' '(' init=statement cond=statement step=expression ')' body=blockStatement;
breakStatement: 'break' ';';
continueStatement: 'continue' ';';
expresstionStatement: expression ';' ;

expression : additionExpression    # AddInEq
            | left=expression op=(EQ|NEQ) right=additionExpression # Eq
            ;
additionExpression : multiplicationExpression                                          # MulInAdd
                    | left = additionExpression op=(PLUS|MINUS) right = multiplicationExpression       # Add
                    ;
multiplicationExpression : primary                                                          # PriInMul
                        | left = multiplicationExpression op=(MULTIPLY|DIVIDE) right = primary    # Mul
                        ;
// 单元或者递归的求属性、函数调用等
primary : unary                                             # UnaryExpr
        | left=primary '(' arguments ')'                    # CallExpr
;
unary
    : NUMBER     #Number
    | IDENTIFIER #Ident
    | '(' expression ')'     #Group
    | FUNCTION '(' params ')' blockStatement  #FunInUnary
;
params: (IDENTIFIER (',' IDENTIFIER)*)?;
arguments: (expression (',' expression)* ','?)?;

FUNCTION: 'function';
PLUS: '+';
MINUS: '-';
EQ: '==';
MULTIPLY: '*';
DIVIDE: '/';
NEQ: '!=';

NUMBER: [0-9]+('.' [0-9]+)? ;
IDENTIFIER: [a-zA-Z_][a-zA-Z_0-9]* ;

EMPTY: [\t\r\n ]+ -> skip;
// 单行注释
LINE_COMMENT: '//' ~[\r\n]* -> skip ;
// 多行注释
BLOCK_COMMENT: '/*' .*? '*/' -> skip ;