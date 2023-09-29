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
$$
\begin{aligned}
    [\text{Program}]&\to            % [Program]
    [\text{Statement}]^*
    \\
    [\text{Statement}]&\to          % [Statement]
    \begin{cases}
        \text{int}\space\it{identifier}\space\text{=}\space[\text{Expresion}]
        \\
        \it{identifier}\space\text{=}\space[\text{Expresion}]
        \\
        [\text{Scope}]
        \\
        \text{exit}\space[\text{Expresion}]
    \end{cases}
    \\
    [\text{Scope}]&\to              % [Scope]
    \begin{cases}
        \{[\text{Statement}]^*\}
    \end{cases}
    \\
    [\text{Expresion}]&\to          % [Expresion]
    \begin{cases}
        [\text{Term}]
        \\
        [\text{Operation}]
    \end{cases}
    \\
    [\text{Term}]&\to               % [Term]
    \begin{cases}
        \it{literal}
        \\
        \it{identifier}
    \end{cases}
    \\
    [\text{Operation}]&\to          % [Operation]
    \begin{cases}
        [\text{Expresion}]\space\text{+}\space[\text{Expresion}]
        \\
        [\text{Expresion}]\space\text{-}\space[\text{Expresion}]
        \\
        [\text{Expresion}]\space\text{*}\space[\text{Expresion}]
        \\
        [\text{Expresion}]\space\text{/}\space[\text{Expresion}]
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
