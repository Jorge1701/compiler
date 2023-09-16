# `to-be-named`
This project is a hobby languaje, I'm just trying to learn a little about asembler and what it takes to create a compiler.

In this project you can find code for a tokenizer, parser and generator, the latter one outputs [Netwide Asembler](https://www.nasm.us/) from which you can later generate a binary for many platforms ...apparently, haven't tried it, works on my machine though.

## Usage
Write a text file with valid syntax (see next section) and extension `.tbd` (will change when I find a good name ...or not).

1. Build the project with `go build` it will procude an executable named `compiler`.
2. Run `./compiler path-to-source-file.tbd`.
    - It will produce an `output.asm`.
3. Use [nasm](https://www.nasm.us/) to generate the object file, it depends on the target platform.
    - It will produce an `output.o`.
    - Example for x86-64 linux `nasm -f elf64 output.asm`
    - Use `nasm --help` and see `-f format` section for other systems.
    - I don't know how other systems react to the asembler my generator outputs, so fingers crossed.
4. Run `ld output.o -o runnable`.
    - It will generate an executable named `runnable`.
    - (I don't remember where I installed `ld`)
5. Then you can just run like `./runnable`.

For now the only thing you can do with the language is return diferent code errores, to see that run `echo $?`, this prints the error code of the last command you executed. Idk how to do that in other platforms...

## Syntax
For now you can only initialize variables of type `int` and return that as the exit code, here's an example:

```
int num = 420
exit num
```

Or simply return some value like this:

```
exit 69
```

Will expand when new things get added.

# How does it work?

## Tokenizer
The tokenizer reads a sequence of runes that contains the source code and returns a list of tokens.

This does not check for gramatical errors like a missing closing parentheses, it just returns the list.

### Tokens
|Separators|Value|Description
|-|-|-
|SEP|`\n`|Separator
|P_L|`(`|Parentheses left
|P_R|`)`|Parentheses right
|B_L|`{`|Brace left
|B_R|`}`|Brace right
|SB_L|`[`|Square brace left
|SB_R|`]`|Square brace right

|Operations|Value|Description
|-|-|-
|ADD|`+`|Addition
|SUB|`-`|Subtraction
|MUL|`*`|Multiplication
|DIV|`/`|Division

|Others|Value|Description
|-|-|-
|EQ|`=`|Assignment operator

|Keyword|Value|Description
|-|-|-
|INT|`int`|Integer type
|EXIT|`exit`|Exit command

|Matchers|Regex|Description
|-|-|-
|IDENTIFIER|`[a-zA-Z][a-zA-Z0-9_]*`|Identifies names of variables
|LITERAL|`[0-9]+`|Represents literal numbers

## Parser
`[To be refactored]`

Takes a list of tokens and generates an abstract syntax tree (AST), which represents the syntactic structure of the source code.

This process drops separators as they are only needed to delimit different parts of the code but do not need to be saved in the AST, since these limits would be implied by the structure itself.

Some nodes of the tree do refer to values from the tokens. For example a node of the tree that represents assignment might want to save the type to be assigned, the name of the variable to create and the value, which in turn might be a node of an expresion that represents addition, since you could assign the result of a sum to the variable instead of a literal.

### Grammar
This is my attempt to specify the grammar that the parser is going to be creating nodes for.

|Name|Definition|
|-|-|
|`Operator`|Represents all the operations that can be done.|
|`Type`|Represents all the types available.|
|`Expr`|Represents all portions of code that resolve into a value.|
|`Stms`|Represents all posible sentences.|

$$
\begin{aligned}
    \boxed{Operator}
    &\Rightarrow
    \begin{cases}
        \text{ADD}
        \\
        \text{SUB}
        \\
        \text{MUL}
        \\
        \text{DIV}
    \end{cases}
    \\\\
    \boxed{Type}
    &\Rightarrow
    \begin{cases}
        \text{INT}
    \end{cases}
    \\\\
    \boxed{Expr}
    &\Rightarrow
    \begin{cases}
        \text{LITERAL}
        \\
        \text{IDENTIFIER}
        \\
        \boxed{Expr}\boxed{Operator} \boxed{Expr}
    \end{cases}
    \\\\
    \boxed{Stmt}
    &\Rightarrow
    \begin{cases}
        \text{EXIT }\boxed{Expr}
        \\
        \boxed{Type}\text{ IDENTIFIER EQ }\boxed{Expr}
        \\
        \text{IDENTIFIER EQ }\boxed{Expr}
    \end{cases}
\end{aligned}
$$

## Generator

`[To be refactored]`

# References
- This project is inspired by [Pixeled](https://www.youtube.com/playlist?list=PLUDlas_Zy_qC7c5tCgTMYq2idyyT241qs)
- [Syscalls table for reference.](https://chromium.googlesource.com/chromiumos/docs/+/master/constants/syscalls.md#x86-32_bit)
- Some learning resources:
    - [x86 Assembly with NASM](https://www.youtube.com/playlist?list=PL2EF13wm-hWCoj6tUBGUmrkJmH1972dBB)
    - [Lexical analysis](https://en.wikipedia.org/wiki/Lexical_analysis)
    - [Abstract syntax tree](https://en.wikipedia.org/wiki/Abstract_syntax_tree)
