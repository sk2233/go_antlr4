grammar Expr;

// 语法规则（Parser Rules）（使用小写字母）
expr: expr op=(MUL|DIV) expr # MulDiv // # Xxx 为这个分支单独生成 VisitXxx 函数
    | expr op=(ADD|SUB) expr # AddSub
    | INT                    # Int
    // 语法规则中的 'xxx' 会自动生成匿名的 token 枚举为 ExprLexerT__{Num} 若是需要后续使用最好显示定义
    | '(' expr ')'           # Paren
    ;

// 词法规则（Lexer Rules）（使用大写字母）
INT: [0-9]+ ;          // 整数
WS: [ \t\r\n]+ -> skip; // 忽略空白符
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';