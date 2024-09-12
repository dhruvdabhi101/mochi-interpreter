# How does this parser works?

`1 + 2 + 3;` is a simple that expression that we are going to parse.

Assuming we have gotten the tokenized input from lexer and we are just parsing.

so, `p.curToken` is `1` and `p.peekToken` is `+`.
we check if we have prefixFn associated with the `p.curToken`, and if we find the function, then then we call that function which returns us some ast node.

after that we run the loop till we find SEMICOLON or the precendec < precedence of current token,

if the loop evaluates to true, then we parse the infix expression, and check if we have the infixFn mapped to the operator, and we move to the next token, and we execute that infix function which we got from the infixParseFn map.

and we recursivly parse the expression till we find the SEMICOLON.

> [info] The above explanation is a high level overview of how the Pratt parser works, and it is not the exact implementation of the parser.


