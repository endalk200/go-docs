> Learn go through concept deep dive, code snippets with comments, excercises and mini projects

# Learn Go

# Table of contents

- [Learn Go](#learn-go)
  - [GoLang History](#golang-history)
  - [Why GoLang?](#why-golang-)
  - [Advantages of GoLang](#advantages-of-golang)
  - [Components of the Golang](#components-of-the-golang)
    - [Compiler](#compiler)
    - [Standard Library](#standard-library)
    - [Runtime](#runtime)
    - [Tools](#tools)
    - [Third-party packages](#third-party-packages)
  - [Go compiler deep dive](#go-compiler-deep-dive)
    - [Phase 1 - Lexing, Parsing and AST generation](#phase-1---lexing--parsing-and-ast-generation)
    - [Phase 2 - Type checking](#phase-2---type-checking)
    - [Phase 3 - SSA generation](#phase-3---ssa-generation)
    - [Phase 4 - Machine code generation](#phase-4---machine-code-generation)
  - [Continue learning](#continue-learning)
  - [Syntax Basics](/docs/data-types.md#data-types-in-go)
    - [Arrays](/docs/data-types.md#arrays)
    - [Slices](/docs/data-types.md#slices)

Google's GoLang, shortened as Go is an open-source programming language that was created in 2009. It is intended for efficient, concurrent, and scalable system programming with an emphasis on simplicity and usability. Go is a popular choice for building networked and distributed systems, web applications, and other types of software because it has strong typing, garbage collection, and syntax similar to C.

## GoLang History

Google first unveiled Go programming language in 2009. It was introduced as a language that could address some of the challenges associated with large-scale distributed systems while remaining easy to use and efficiently driving the development of Go.

The language was developed by three Google engineers: Robert Griesemer, Rob Pike, and Ken Thompson. Go's first public release was in March 2012, and it has since grown in popularity among developers due to its simplicity, concurrency, and performance. Today, companies like Google, Uber, Dropbox, and many others use Go to build high-performance, scalable, and reliable software systems, making it one of the fastest-growing programming languages.

## Why GoLang?

There are several reasons why GoLang should be used to build a program. First and foremost, Go is intended to be fast and efficient and unlike an interpreted language, it has a compiler that generates native machine code and a garbage collector that manages memory automatically, eliminating the need for manual memory management.

Second, Go has built-in concurrency support, allowing developers to easily write programs that can run multiple tasks or processes at the same time. As a result, Go is an excellent choice for developing high-performance, parallel systems, for example, web servers or networked services.

Furthermore, Go has a simple and concise syntax, making it simple to learn and read. It has an expanding library of open-source packages and tools that can assist developers in developing a wide range of applications.

Finally, Go is supported by a large and active developer community that contributes to the language and helps ensure its ongoing software development and improvement. As a result, Go is an excellent choice for developing efficient, scalable, and dependable software systems.

## Advantages of GoLang

There are several advantages of using GoLang:

- Go's efficient performance is due to its small memory footprint, quick compilation times, and ability to generate native machine code. As a result, it's ideal for creating high-performance applications like web servers or networked services that can handle a large volume of requests without slowing down.
- Concurrency is built into Go through Goroutines and channels, allowing developers to write programs that can run multiple tasks or processes simultaneously. This makes creating efficient and scalable systems simple without dealing with the complexities of manually managing threads or processes.
- Go has a simple and concise syntax that is simple to learn and read. The syntax is similar to C but has certain features like garbage collection and strong typing. This makes it an excellent choice for developers who are new to programming or want to write clean, maintainable code.
- Its strong and static type system aids in the prevention of common programming errors like null pointer dereferences and buffer overflows. This makes it an excellent choice for developing dependable and secure software, especially in industries such as finance and healthcare, where software errors can have serious consequences.
- Go has a large and active developer community that contributes to the language and provides a growing Go standard library of open-source packages and tools to help developers build a wide range of applications. Everything from web frameworks and database drivers to machine learning libraries and blockchain tools is included.

## Components of the Golang

The components of GoLang are:

### Compiler

GoLang's compiler compiles the Go source code into executable files that can be run on a specific platform or architecture. The Go compiler is fast and efficient, producing native machine code for the target platform.

### Standard Library

GoLang includes a comprehensive standard library that includes packages for networking, I/O, text processing, cryptography, and much more. The standard library includes various useful tools and utilities developers can use to create their applications.

### Runtime

Go programming language's runtime includes a garbage collector and a scheduler for managing Goroutines (concurrent functions) and channels (interprocess communication). The Go runtime is in charge of memory management and ensuring that Goroutines are executed in an efficient and predictable manner.

```go
func main() {
    ch := make(chan int)

    go square(2, ch)
    go square(3, ch)
    go square(4, ch)

    fmt.Println("Hello World!")
    fmt.Println(<-ch)
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}

func square(x int, ch chan int) {
    ch <- x * x
}
```

In this example, we create a channel

```
ch
```

and call the

```
square()
```

function in parallel using Goroutines. Each

```
square()
```

function call computes the square of a number and returns the result to the channel

```
ch
```

. The main function receives the channel results and prints them to the console.

Learn to simplify your programs and make your code modular using package main and import:

```go
package main

import "fmt"

func main() {
    x := 3

    y := square(x)

    fmt.Println("Hello World! The square of %d is %d\n", x, y)
}

func square(n int) int {
    return n * n
}
```

### Tools

GoLang programming language includes command-line tools for building, testing, and debugging Go programs. These tools include the Go command for managing Go packages and modules and the Go debugger for debugging Go programs.

### Third-party packages

Go has a large and growing ecosystem of third-party packages and libraries that can be used to extend the language's capabilities. These packages cover web development, machine learning, databases, and other topics.

These components work together to provide a powerful and flexible programming environment for building efficient, scalable, and reliable software systems.

## Go compiler deep dive

The Go compiler is a tool that compiles Go source code into executable files that can be run on a specific platform or architecture. The Go compiler is fast and efficient, producing native machine code for the target platform.

The compiler may be logically split in four phases, which we will briefly describe alongside the list of packages that contain their code.

You may sometimes hear the terms "front-end" and "back-end" when referring to the compiler. Roughly speaking, these translate to the first two and last two phases we are going to list here. A third term, "middle-end", often refers to much of the work that happens in the second phase.

Note that the `go/_` family of packages, such as go/parser and `go/types`, are mostly unused by the compiler. Since the compiler was initially written in C, the `go/_` packages were developed to enable writing tools working with Go code, such as gofmt and vet. However, over time the compiler's internal APIs have slowly evolved to be more familiar to users of the go/\* packages.

It should be clarified that the name "gc" stands for "Go compiler", and has little to do with uppercase "GC", which stands for garbage collection.

The four phases are:

- Lexing, Parsing and AST generation
- Type checking
- SSA generation
- Machine code generation

![Compilation phases](/docs/assets/compilation-phases.svg)

### Phase 1 - Lexing, Parsing and AST generation

See diagram below for a visual representation of this phase:

![Lexing, Parsing and AST generation](/docs/assets/lexing-parsing-ast-generation.svg)

**Lexer**

The first step of every compiler is to break up the raw source code text into tokens, which is done by the scanner (also known as lexer). Tokens can be keywords, strings, variable names, function names, etc. Every valid program “word” is represented by a token. In concrete terms for Go, this might mean we have a token “package”, “main”, “func” and so forth. Each token is represented by its position, type, and raw text in Go.

To see list of known tokens refer to this code from the go codebase [here](https://github.com/golang/go/blob/master/src/cmd/compile/internal/syntax/tokens.go)

Say that lexer’s is at character `"`, this could only mean that the current token is a string so the lexer would run the stdString method to verify that the current token is in-fact a string (see inline comments for details):

```go
func (s *scanner) stdString() {
	ok := true // marks if the string is lexically valid.
	s.nextch() // moves the underlying source pointer to the start of the string

	for {
		if s.ch == '"' {
			s.nextch() // moves the underlying source pointer to the character after
      '"'
			break
		}

        // ...

		if s.ch == '\n' { // an example of an error: standard strings shouldn't contain new line characters
			s.errorf("newline in string")
			ok = false // the literal token will be marked as invalid because of the
      ok value
			break
		}
		if s.ch < 0 { // we never closed the the opening '"' and reached EOF
			s.errorAtf(0, "string not terminated")
			ok = false
			break
		}
		s.nextch() // everything else goes
	}

    // sets the internal scanner state to the current lexer with its value.
	s.setLit(StringLit, ok)
}
```

Go even allows us to execute the scanner ourselves in a Go program by using the go/scanner and go/token packages. That means we can inspect what our program looks like to the Go compiler after it has been scanned. To do so, we are going to create a simple program that prints all tokens of a Hello World program. The program will look like this:

```go
package main

import (
  "fmt"
  "go/scanner"
  "go/token"
)

func main() {
  src := []byte(`package main
import "fmt"
func main() {
  fmt.Println("Hello, world!")
}
`)

  var s scanner.Scanner

  fset := token.NewFileSet()
  file := fset.AddFile("", fset.Base(), len(src))
  s.Init(file, src, nil, 0)

  for {
     pos, tok, lit := s.Scan()
     fmt.Printf("%-6s%-8s%q\n", fset.Position(pos), tok, lit)

     if tok == token.EOF {
        break
     }
  }
}
```

We will create our source code string and initialize the scanner.Scanner struct which will scan our source code. We call Scan() as many times as we can and print the token’s position, type, and literal string until we reach the End of File (EOF) marker. When we run the program, it will print the following:

```
1:1   package "package"
1:9   IDENT   "main"
1:13  ;       "\n"
2:1   import  "import"
2:8   STRING  "\"fmt\""
2:13  ;       "\n"
3:1   func    "func"
3:6   IDENT   "main"
3:10  (       ""
3:11  )       ""
3:13  {       ""
4:3   IDENT   "fmt"
4:6   .       ""
4:7   IDENT   "Println"
4:14  (       ""
4:15  STRING  "\"Hello, world!\""
4:30  )       ""
4:31  ;       "\n"
5:1   }       ""
5:2   ;       "\n"
5:3   EOF     ""
```

Here we can see what the Go parser uses when it compiles a program. What we can also see is that the scanner adds semicolons where those would usually be placed in other programming languages such as C. This explains why Go does not need semicolons: they are placed intelligently by the scanner.

**Parsing - AST Generation**

The parser is (in a very simple terms) a software component that given a source file constructs an AST (Abstract Syntax Tree) from it, which itself is composed of nodes that have a semantic relationship between them according to the language definition.

The Go compiler defines its list of [nodes](https://github.com/golang/go/blob/master/src/cmd/compile/internal/syntax/nodes.go)

Every node should adhere to the [Node](https://github.com/golang/go/blob/345184496ce358e663b0150f679d5e5cf1337b41/src/cmd/compile/internal/syntax/nodes.go#L10) interface, which stats that every Node should be able to give it’s position in the source, and to accomplish this, the compiler authors performed a very neat trick of creating a dummy struct that implements the Node interface and added it as an embedded field to all of the other node definitions, this way, they didn’t have to implement the same thing over and over again for every node (neat).

This trick is used in the same way to mark nodes that are Declarations and those that are Expressions. The linked Go file has all of the node definitions, please take a look (it is very well documented) before continuing.

The [parser](https://github.com/golang/go/blob/345184496ce358e663b0150f679d5e5cf1337b41/src/cmd/compile/internal/syntax/parser.go#L17) uses the lexer to produce tokens from the given source file, the parsing loop is fairly simple as shown below:

```go
for p.tok != _EOF {
  switch p.tok {
  case _Const:
    p.next()
    f.DeclList = p.appendGroup(f.DeclList, p.constDecl)

  case _Type:
    p.next()
    f.DeclList = p.appendGroup(f.DeclList, p.typeDecl)

  case _Var:
    p.next()
    f.DeclList = p.appendGroup(f.DeclList, p.varDecl)

  case _Func:
    p.next()
    if d := p.funcDeclOrNil(); d != nil {
      f.DeclList = append(f.DeclList, d)
    }
    // ...
  }
}
 // ...
return f
```

Parsing a file is as simple as taking the first token of each new top level statement and deciding how to parse that and add it to the children of the current root node (i.e the file).

The design of the parser itself is based of having for each non-terminal rule a corresponding function that handles parsing that rule specifically, as an example if we take the [ConstDecl](https://github.com/golang/go/blob/5a6a830c1ceafd551937876f11590fd60aea1799/src/cmd/compile/internal/syntax/nodes.go#L66) declaration node, it has a corresponding parsing function in the parser [constDecl](https://github.com/golang/go/blob/5a6a830c1ceafd551937876f11590fd60aea1799/src/cmd/compile/internal/syntax/parser.go#L553), you can expect similar things for other declarations/expressions as well.

For a better understanding of parsing top level declarations lets walk through an example, the var declaration.

The Go language allows to declare top level variables using the var keyword, if you are reading this and never seen Golang code before, it looks something like this:

```go
var a = 10              // simple declaration
var b int               // simple declaration with a type
var c, d int            // two variable declaration
var e, f int = 10, 12   // two variable declarations with value and types
```

To parse this declaration, the parser uses the [varDecl](https://github.com/golang/go/blob/5a6a830c1ceafd551937876f11590fd60aea1799/src/cmd/compile/internal/syntax/parser.go#L703) function:

```go
func (p *parser) varDecl(group *Group) Decl {
  // ...

	d := new(VarDecl)
	d.pos = p.pos()
	d.Group = group
	d.Pragma = p.takePragma()

	d.NameList = p.nameList(p.name())
	if p.gotAssign() {
		d.Values = p.exprList()
	} else {
		d.Type = p.type_()
		if p.gotAssign() {
			d.Values = p.exprList()
		}
	}
	return d
}
```

So we start by creating a VarDecl node since this is what we expect to return after this function invocation, we initialize it with some metadata from the parser such as the current position, the [Group](https://github.com/golang/go/blob/69756b38f25bf72f1040dd7fd243febba89017e6/src/cmd/compile/internal/syntax/nodes.go#L118) etc …

We start by parsing a nameList which takes care of the comma-separated names on the left side of the declaration and from there it’s one of two things: we have a `=` character which means its an initialization with values and we set the values as parsed expression list, or we have a type declaration before and we parse that before trying to parse the actual values of the variables.

_Note about AST nodes_ -
The AST nodes defined by the compiler are not the same ones defined in the `ast` module that is used by go tools like `gofmt`.

### Phase 2 - Type checking

- `cmd/compile/internal/types2` (type checking)

The types2 package is a port of `go/types` to use the syntax package's
AST instead of `go/ast`.

**IR construction ("noding")**

- `cmd/compile/internal/types` (compiler types)
- `cmd/compile/internal/ir` (compiler AST)
- `cmd/compile/internal/typecheck` (AST transformations)
- `cmd/compile/internal/noder` (create compiler AST)

The compiler middle end uses its own AST definition and representation of Go
types carried over from when it was written in C. All of its code is written in
terms of these, so the next step after type checking is to convert the syntax
and types2 representations to ir and types. This process is referred to as
"noding."

There are currently two noding implementations:

1. irgen (aka "-G=3" or sometimes "noder2") is the implementation used starting
   with Go 1.18, and

2. Unified IR is another, in-development implementation (enabled with
   `GOEXPERIMENT=unified`), which also implements import/export and inlining.

Up through Go 1.18, there was a third noding implementation (just
"noder" or "-G=0"), which directly converted the pre-type-checked
syntax representation into IR and then invoked package typecheck's
type checker. This implementation was removed after Go 1.18, so now
package typecheck is only used for IR transformations.

**Middle end**

- `cmd/compile/internal/deadcode` (dead code elimination)
- `cmd/compile/internal/inline` (function call inlining)
- `cmd/compile/internal/devirtualize` (devirtualization of known interface method calls)
- `cmd/compile/internal/escape` (escape analysis)

Several optimization passes are performed on the IR representation:
dead code elimination, (early) devirtualization, function call
inlining, and escape analysis.

**Walk**

- `cmd/compile/internal/walk` (order of evaluation, desugaring)

The final pass over the IR representation is "walk," which serves two purposes:

1. It decomposes complex statements into individual, simpler statements,
   introducing temporary variables and respecting order of evaluation. This step
   is also referred to as "order."

2. It desugars higher-level Go constructs into more primitive ones. For example,
   `switch` statements are turned into binary search or jump tables, and
   operations on maps and channels are replaced with runtime calls.

### Phase 3 - SSA generation

- `cmd/compile/internal/ssa` (SSA passes and rules)
- `cmd/compile/internal/ssagen` (converting IR to SSA)

In this phase, IR is converted into Static Single Assignment (SSA) form, a
lower-level intermediate representation with specific properties that make it
easier to implement optimizations and to eventually generate machine code from
it.

During this conversion, function intrinsics are applied. These are special
functions that the compiler has been taught to replace with heavily optimized
code on a case-by-case basis.

Certain nodes are also lowered into simpler components during the AST to SSA
conversion, so that the rest of the compiler can work with them. For instance,
the copy builtin is replaced by memory moves, and range loops are rewritten into
for loops. Some of these currently happen before the conversion to SSA due to
historical reasons, but the long-term plan is to move all of them here.

Then, a series of machine-independent passes and rules are applied. These do not
concern any single computer architecture, and thus run on all `GOARCH` variants.
These passes include dead code elimination, removal of
unneeded nil checks, and removal of unused branches. The generic rewrite rules
mainly concern expressions, such as replacing some expressions with constant
values, and optimizing multiplications and float operations.

### Phase 4 - Machine code generation

- `cmd/compile/internal/ssa` (SSA lowering and arch-specific passes)
- `cmd/internal/obj` (machine code generation)

The machine-dependent phase of the compiler begins with the "lower" pass, which
rewrites generic values into their machine-specific variants. For example, on
amd64 memory operands are possible, so many load-store operations may be combined.

Note that the lower pass runs all machine-specific rewrite rules, and thus it
currently applies lots of optimizations too.

Once the SSA has been "lowered" and is more specific to the target architecture,
the final code optimization passes are run. This includes yet another dead code
elimination pass, moving values closer to their uses, the removal of local
variables that are never read from, and register allocation.

Other important pieces of work done as part of this step include stack frame
layout, which assigns stack offsets to local variables, and pointer liveness
analysis, which computes which on-stack pointers are live at each GC safe point.

At the end of the SSA generation phase, Go functions have been transformed into
a series of obj.Prog instructions. These are passed to the assembler
(`cmd/internal/obj`), which turns them into machine code and writes out the
final object file. The object file will also contain reflect data, export data,
and debugging information.

## Continue learning

Now that you have a basic understanding of go on a high level and the components of go, let's dive into the language itself. We will start with the basics and then move on to more advanced topics. Use the table of contents below to dig deep into different concepts in go.
