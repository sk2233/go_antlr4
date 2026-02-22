grammar Expr;

// 语法规则（Parser Rules）
expr: expr op=('*'|'/') expr # MulDiv
    | expr op=('+'|'-') expr # AddSub
    | INT                    # Int
    | '(' expr ')'           # Paren
    ;

// 词法规则（Lexer Rules）
INT: [0-9]+ ;          // 整数
WS: [ \t\r\n]+ -> skip; // 忽略空白符