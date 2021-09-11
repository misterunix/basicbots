# basicbots Manual

This is the manual for **basicbots** programming game.

---

## What is baiscbots? 
It is a programming game where you program your robot tanks to compete agaist other robot tanks.

The idea for this isn't new. My first introduction to programming combat games was [Tom Poindex's](https://github.com/tpoindex) original [crobots](http://tpoindex.github.io/crobots/) from 1985. Which is, imho, still a whole lot of fun to play. 

**basicbot's** programming language is a subset of *BASIC*. Why *BASIC*? Well, because why not? *BASIC* is not a completely dead language, but for those of us older enough this should take you back in time to the versions around the late 1970's and early 80's. If not, that is okay as well. You can get a sample of what it was like to live in the 8bit world when computers were simpler.

**basicbots** version of *BASIC* is based and adapted from [gobasic](https://github.com/skx/gobasic) by [James Kemp](https://github.com/skx). Mr. Kemp has gracely given me permission to adapt [gobasic](https://github.com/skx/gobasic) for my game.

## The simulation
---

### Cardinal directions
|     |     |     |
|:----|:---:| ---:|
| 225 | 270 | 315 |
| 180 |  *  | 000 |
| 135 | 090 | 045 |
|     |     |     |

### The Battlefield
The robots compete in a virtual battlefield that is 1000 units wide by 1000 units long. The upper left corner is 0,0 while the lower right corner is 1000,1000. There is a virtual wall that surrounds the battlefield. If a robot makes contact with this wall, it will take damage and its speed will be reduced to 0.

### The Robot's Hardware
The robot tank has three hardware systems.

- **A Scanner**. The scanner can scan in any direction instantly but has a angular resolution from +-2 degrees to +- 10 degrees. Example: Scan 90 degrees with a width of 2 would return results that are between 88 - 92 degrees. A width of 10 would return results that are from 80 to 100 degrees.

- **A Cannon**. The Cannon can fire a projectile in any direction but at a maximum range of 700 units. Any robot caught in the blast radius of 40 units. The farther from the center of the blast, the less damage a robot will take. Each robot can only have two projectiles in the air at any given time and it takes time to reload. 

- **A Drive system**. The drive system has two parameters. Angle and speed. The angle can be any angle from 0 to 360 degrees and the speed can be from 0 to 100%. The robots does require time to reach a given speed and requires time to slow down. The robot can only nogotate turns when the speed is less than 50%. If the speed is above 50% the robot will slow down to 50% then turn and increase speed to the previous setting.

### The Robot's Status
The robot has access to:
- **The X location**. The current X position on the battlefield of the robot.
- **The Y location**. The current Y position on the battlefield of the robot.
- **The Speed**. The current speed of the robot. 
- **The Damage**. The current amount of damage inflicted upon the robot. The robot can withstand up to 100 points of damage. Once a robot has reached 100 points, it has died and is removed from the battlefield. The amount of damage does not affects its performance.

Those are the seven commands and sensors your program has access to.

## BASIC
---
**basicbot** uses a subset of the *BASIC* language. Most of the early BASICs funcionality is implimented. The assumtion is that the players are passingly familuar with *BASIC*. This document is not meant to teach the *BASIC* programming language. I encurenge you to seek out online guides.

### Qurkes
- Variables are case senitive. Primitives are not.
- Line numbers are not strictly required but are strongly suggested.

### Limitations
A previously stated, this is a subset of the *BASIC* standard. 
- There is no editor.
- Errors in the program are only found during execicution.
- Only a single statement is allowed on each line. The excption is for `REM`arks.
- The only data types are floating-point and strings.
- Arrays can have a maximum of two dimentions.
- `LET` is not optional.
- Only a single statement is allowed between `THEN` and `ELSE` and a single statement from `ELSE` to a newline.

### Program format

Programs are text files that are specified on the command line. They are loaded and parsed when the robots are initlized. 

Each line contains a line number followed by a statment or command with an optional parameter(s). 

All line numbers are integers.

### COMMANDS

- `LET` variable = expression
  - Assigns a value, or the result of expression, to a variable. Variables are case sensitive and can be floating-point or strings.

- `IF` condition `THEN` statement `ELSE` starement
  - If condition is true, then the statement following `THEN` is executed. If it is false then the statement after `ELSE` is executed. 

- `GOTO` line number
  - Transfer execution to another part of the  program specified by the line number.

- `GOSUB` line number
  - Transfers execution to a subroutine specified by the line number. The position of the `GOSUB` is remembered so that a `RETURN` can bring program execution back to the statement following the `GOSUB`.

- `RETURN`
  - Returns from a subroutine. Program execution returns to the statement following the `GOSUB` which called the present subroutine.

- `END`
  - Terminates program execution.

~~PRINT output-list~~

~~Produces output to the console. Output-list is a list of items separated by commas. Each item can be either a string literal enclosed in double quotation marks, or a numeric expression. An end of line sequence is output after all the values, so that the next PRINT statement will put its output on a new line.~~

~~INPUT variable-list~~

~~Asks for input from the console. Variable-list is a list of variable names. For each variable given, a question mark is output and the value typed by the user is stored in that variable. Tinybasic allows multiple values to be typed by the user on one line, each separated by any non-numeric character.~~

- `REM` comment-text
  - Provides space for free-format comment text in the program. Comments have no effect on the execution of a program, and exist only to provide human-readable information to the programmer. 

### EXPRESSIONS

Expressions in **basicbots** *BASIC* are purely arithmetic expressions. The four basic arithmetic operators are supported: multiplication (*), division (/), addition (+) and subtraction (-). Unary operators for positive (+) and negative (-) are supported, as are parentheses for affecting the order of operations.

Standard operator precedence evaluates parentheses first, then unary signs, then multiplication and division, with addition and subtraction last.

CONDITIONS

The relational operators are =, >, <, <> or ><, >=, and <=. They are not supported within arithmetic expressions, but can only be used as conditions in IF statements in the form: expression relational-operator expression

COMPILATION

Tinybasic is capable of compiling programs into executables with the help of a C compiler. To use this facility, the TBEXE environment variable must be set before invoking tinybasic. The variable should contain the command that compiles a C program into an executable, and may contain the following tokens:

$(SOURCE): the C source filename is substituted here.
$(TARGET): a target filename is substituted here.

The C source filename will be the same as the BASIC filename but with the extension .c added. The target filename is the BASIC source filename with the .bas extension removed; if the BASIC source filename has no extension, then .out is added to prevent the source being overwritten by the executable. If your operating system requires an extension like .exe for its executables, then you need to add it explicitly (i.e. $(TARGET).exe) - unless the compiler adds that itself. As an example, the file test.bas could be compiled on a Unix system with the following commands:

$ TBEXE=’gcc -o $(TARGET) $(SOURCE)’
$ tinybasic -Oexe test.bas

This would produce the executable file test, and as a side effect, the C source file test.bas.c.

ERROR MESSAGES

Program error messages can be in one of two forms:

Parse error: description, line line-number, label line-label
Run-time error: description, label line-label

Parse errors are those that are detected before the program starts. Run-time errors are those that cannot be detected until the program is running. If a parse error is detected on a line without a label, then the label section is omitted from the error message. The error messages and their meanings are as follows.

Invalid line number
One of the following has occurred: (i) a line label is missing when line numbers are mandatory; (ii) a line label is lower than the previous one when line numbers are mandatory or implied.

Unrecognised command
The command keyword is not recognised. Note that REM will not be recognised when comments are disabled, and will produce this error.

Invalid variable
In a LET or INPUT statement, something other than a letter from A to Z was supplied when a variable name was expected.

Invalid assignment
The = sign was missing from a LET statement.

Invalid expression
An expression in this line is invalid. It is possibly lacking an operator, variable or value where one is expected.

Missing )
An expression contains a left parenthesis and no corresponding right parenthesis.

Invalid PRINT output
Something is wrong with the output list in a PRINT statement. It could be: (i) completely missing, (ii) missing a separator between two items, or (iii) missing an item between two separators or at the start or end of the list.

Invalid operator
An unrecognised operator was encountered in an expression or a condition.

THEN expected
The mandatory THEN keyword is missing from its expected place in an IF statement.

Unexpected parameter
A parameter was given to a command that should not have one, such as END or RETURN.

RETURN without GOSUB
A RETURN was encountered without having executed a GOSUB. This commonly occurs when a programmer forgets to put an END or a GOTO before a subroutine, and allows execution to blunder into it.

Divide by zero
The divisor in an expression was 0. If dividing by a variable or an expression, it is advisable to check beforehand that it cannot be zero. An intentional division by zero is not the most graceful way to stop a program.

Overflow
When given as a parse error, there is a value in the program that is outside the range of -32768 to 32767. When given as a runtime error, an expression in the program or an input from the user has produced a result outside this range.

VERSION INFORMATION

This manual page documents tinybasic, version 1.0.

AUTHORS

Tiny BASIC was originally designed by Dennis Allison. This implementation was written by Damian Gareth Walker.

EXAMPLE

This program prints out all of the numbers in the Fibonnaci series between 0 and 1000.

    LET A=0
    LET B=1
    PRINT A
100 PRINT B
    LET B=A+B
    LET A=B-A
    IF B<=1000 THEN GOTO 100
    END











































































































# basicbots

A programming game where you program your robot tanks to compete agaist other robot tanks.

The idea for this isn't new. My first introduction to programming combat games was [Tom Poindex's](https://github.com/tpoindex) original [crobots](http://tpoindex.github.io/crobots/) from 1985. Which is, imho, still a whole lot of fun to play. 

**basicbot's** programming language is a subset of *BASIC*. Why *BASIC*? Well, because why not? *BASIC* is not a completely dead language, but this should take you back in time to the versions around the late 1970's and early 80's. Just right in that sweet spot when I was learning to program on my TRS-80 model 1.

basicbot's BASIC is based on [gobasic](https://github.com/skx/gobasic) by [Seve Kemp](https://github.com/skx)

---
---

---
---

## BASIC

---
basicbot uses a subset of the *BASIC* language. Most of the early BASICs funcionality is implimented. The assumtion is that are passingly familuar with *BASIC*. The document is not meant to teach the *BASIC* programming language. I encurenge you to seek out online guides.

### Qurkes
- Variables are case senitive. Primitives are non.
- Line numbers are not strictly required but are strongly suggested.

### Limitations
A previously stated, this is a subset of the *BASIC* standard. 
- Only a single statement is allowed on each line. The excption is for REMarks.
- The only data types are floating-point and strings.
- Arrays can have a maximum of two dimentions.
- `LET` is not optional.
- Only a single statement is allowed between `THEN` and `ELSE` and a single statement from `ELSE` to a newline.

---
# Intoduction to programming in *BASIC*





### Comparison
- `IF a < b THEN ..`
- `IF a > b THEN ..`
- `IF a <= b THEN ..`
- `IF a >= b THEN ..`
- `IF a = b THEN ..`
- `IF a <> b THEN ..`
- `IF a THEN ..`
  - This passes if `a` is a number which is not zero.
  - This passes if `a` is a string which is non-empty.

### BASIC Primitives

- Comments 
  - `REM` Single-line comment only.

- Variable assignment
  - `LET`
  - `DIM`

- Loops
  - `FOR` & `NEXT`

- Conditionals
  - `IF` / `THEN` / `ELSE`

- Communication
  - `PRINT` & `INPUT`

- Program termination
  - `END`

- Branching
  - `GOTO` & `GOSUB` / `RETURN`

- Stored data
  - `READ` & `DATA`

- Misc
  - `SWAP`
  - `DEF FN` & `FN`

- Strings
  - `LEN`
  - `LEFT$` & `RIGHT$`
  - `CHR$`
  - `CODE`

- Math
  -``


### Extended Primitives

# 05 PRINT "Index"

* [10 PRINT "GOBASIC!"](#10-print-gobasic)
* [20 PRINT "Limitations"](#20-print-limitations)
  * [Arrays](#arrays)
  * [Line Numbers](#line-numbers)
  * [IF Statement](#if-statement)
  * [DATA / READ Statements](#data--read-statements)
  * [Builtin Functions](#builtin-functions)
  * [Types](#types)
* [30 PRINT "Installation"](#30-print-installation)
  * [Build without Go Modules (Go before 1.11)](#build-without-go-modules-go-before-111)
  * [Build with Go Modules (Go 1.11 or higher)](#build-with-go-modules-go-111-or-higher)
* [40 PRINT "Usage"](#40-print-usage)
* [50 PRINT "Implementation"](#50-print-implementation)
* [60 PRINT "Sample Code"](#60-print-sample-code)
* [70 PRINT "Embedding"](#70-print-embedding)
* [80 PRINT "Visual BASIC!"](#80-print-visual-basic)
* [90 PRINT "Bugs?"](#90-print-bugs)
* [100 PRINT "Project Goals / Links"](#100-print-project-goals--links)

<br />
<br />
<br />
<br />

# 10 PRINT "GOBASIC!"

This repository contains a naive implementation of BASIC, written in Golang.

> If you'd prefer to see a more "real" interpreted language, implemented in Go, you might prefer [monkey](https://github.com/skx/monkey/).

The implementation is simple for two main reasons:

* There is no UI, which means any and all graphics-primitives are ruled out.
  * However the embedded sample, described later in this file, demonstrates using BASIC to create a PNG image.
  * There is also a HTTP-based BASIC server, also described later, which allows you to create images "interactively".
* I didn't implement the full BASIC set of primitives.
  * Although most of the commands available to the ZX Spectrum are implemented. I only excluded things relating to tape, PEEK, POKE, etc.
  * If you want to add new BASIC keywords this is easy, and the samples mentioned above do that.

The following obvious primitives work as you'd expect:

* `DIM`
  * Create an array.  Note that only one and two-dimensional arrays are supported.
  * See [examples/95-arrays.bas](examples/95-arrays) and [examples/40-array-sort.bas](examples/40-array-sort.bas) for quick samples.
* `END`
  * Exit the program.
* `GOTO`
  * Jump to the given line.
* `GOSUB` / `RETURN`
  * Used to call the subroutines at the specified line.
* `IF` / `THEN` / `ELSE`
  * Conditional execution.
* `INPUT`
  * Allow reading a string, or number (see later note about types).
* `LET`
  * Assign a value to a variable, creating it if necessary.
* `FOR` & `NEXT`
  * Looping constructs.
* `PRINT`
  * Print a string, an integer, or variable.
  * Multiple arguments may be separated by commas.
* `REM`
  * A single-line comment (BASIC has no notion of multi-line comments).
* `READ` & `DATA`
  * Allow reading from stored data within the program.
  * See [examples/35-read-data.bas](examples/35-read-data.bas) for a demonstration, along with [examples/100-array-sort.bas](examples/100-array-sort.bas).
* `SWAP`
  * Allow swapping the contents of two variables.
  * Useful for sorting arrays, as shown in [examples/100-array-sort.bas](examples/100-array-sort.bas).
* `DEF FN` & `FN`
  * Allow user-defined functions to be defined or invoked.
  * See [examples/25-def-fn.bas](examples/25-def-fn.bas) for an example.

Most of the maths-related primitives I'm familiar with are also present, for example SIN, COS, PI, ABS, along with the similar string-related primitives:

* `LEN "STEVE"`
  * Returns the length of a string "STEVE" (5).
* `LEFT$ "STEVE", 2`
  * Returns the left-most 2 characters of "STEVE" ("ST").
* `RIGHT$ "STEVE", 2`
  * Returns the right-most 2 characters of "STEVE" ("VE").
* `CHR$ 42`
  * Converts the integer 42 to a character (`*`).  (i.e. ASCII value.)
* `CODE " "`
  * Converts the given character to the integer value (32).

<br />
<br />
<br />
<br />

## 20 PRINT "Limitations"

This project was started as [a weekend-project](https://blog.steve.fi/so_i_wrote_a_basic_basic.html), although it has subsequently been improved and extended.

The code has near-total test-coverage, and has been hammered with multiple days of fuzz-testing (i.e. Feeding random programs into the interpreter to see if it will die - see [FUZZING.md](FUZZING.md) for more details on that.)

That said there are some (obvious) limitations:

* Only a single statement is allowed upon each line.
* Only a subset of the language is implemented.
  * If there are specific primitives you miss, then please [report a bug](https://github.com/skx/gobasic/issues/).
    * The project is open to suggestions, but do bear in mind the [project goals]((#100-print-project-goals--links)) listed later on.
* When it comes to types only floating-point and string values are permitted.
  * There is support for arrays but only one or two dimensional ones.

### Arrays

Arrays are used just like normal variables, but they need to be declared using the `DIM` statement.   Individual elements are accessed using the offsets in brackets after the variable name:

    10 DIM a(10,10)
    20 LET a[1,1]=10
    30 PRINT a[1,1], "\n"

Arrays are indexed from 0-N, so with an array size of ten you can access eleven
elements:

     10 DIM a(10)
     20 a[0] = 0
     30 a[1] = 1
     40 ..
     90 a[9] = 9
    100 a[10] = 10

ZX Spectrum BASIC indexed arrays from 1, denying the ability to use the zeroth element, which I've long considered a mistake.


### Line Numbers

Line numbers are _mostly_ optional, for example the following program is valid and correct:

     10 READ a
     20 IF a = 999 THEN GOTO 100
     30 PRINT a, "\n"
     40 GOTO 10
    100 END
        DATA 1, 2, 3, 4, 999

The main reason you need line-numbers is for the `GOTO` and `GOSUB` functions,
if you prefer to avoid them then you're welcome to do so.

### `IF` Statement

The handling of the IF statement is perhaps a little unusual, since I'm
used to the BASIC provided by the ZX Spectrum which had no ELSE clause.
The general form of the IF statement I've implemented is:

    IF $CONDITIONAL THEN $STATEMENT1 [ELSE $STATEMENT2]

Only a single statement is permitted between "THEN" and "ELSE", and again between "ELSE" and NEWLINE.  These are valid IF statements:

    IF 1 > 0 THEN PRINT "OK"
    IF 1 > 3 THEN PRINT "SOMETHING IS BROKEN": ELSE PRINT "Weird!"

In that second example you'll see that "`:`" was used to terminate the `PRINT` statement, which otherwise would have tried to consume all input until it hit a newline.

The set of comparison functions _probably_ includes everything you need:

* `IF a < b THEN ..`
* `IF a > b THEN ..`
* `IF a <= b THEN ..`
* `IF a >= b THEN ..`
* `IF a = b THEN ..`
* `IF a <> b THEN ..`
* `IF a THEN ..`
  * This passes if `a` is a number which is not zero.
  * This passes if `a` is a string which is non-empty.

You can see several examples of the IF statement in use in the example [examples/70-if.bas](examples/70-if.bas).



### `DATA` / `READ` Statements

The `READ` statement allows you to read the next value from the data stored
in the program, via `DATA`.  There is no support for the `RESTORE` function,
so once your data is read it cannot be re-read.


### Builtin Functions

You'll also notice that the primitives which are present all suffer from the flaw that they don't allow brackets around their arguments.  So this is valid:

    10 PRINT RND 100

But this is not:

    10 PRINT RND(100)

This particular problem could be fixed, but I've not considered it significant.


### Types

There are no type restrictions on variable names vs. their contents, so these statements are each valid:

* `LET a = "steve"`
* `LET a = 3.2`
* `LET a% = ""`
* `LET a$ = "steve"`
* `LET a$ = 17 + 3`
* `LET a% = "string"`

The __sole__ exception relates to the `INPUT` statement.  The `INPUT` statement prompts a user for input, and returns it as a value - it doesn't know whether to return a "string" or a "number".  So it returns a string if it sees a `$` in the variable name.

This means this reads a string:

    10 INPUT "Enter a string", a$

But this prompts for a number:

    10 INPUT "Enter a number", a

This seemed better than trying to return a string, unless the input looked like a number (i.e. the input matched `/^([0-9\.]+)$/` we could store a number, otherwise a string).


<br />
<br />
<br />
<br />


## 30 PRINT "Installation"

### Build without Go Modules (Go before 1.11)

Providing you have a working [go-installation](https://golang.org/) you should be able to install this software by running:

    go get -u github.com/skx/gobasic

**NOTE** This will only install the command-line driver, rather than the HTTP-server, or the embedded example code.

### Build with Go Modules (Go 1.11 or higher)

    git clone https://github.com/skx/gobasic ;# make sure to clone outside of GOPATH
    cd gobasic
    go install

If you don't have a golang environment setup you should be able to download various binaries from the github release page:

* [Binary Releases](https://github.com/skx/gobasic/releases)


<br />
<br />
<br />
<br />


## 40 PRINT "Usage"

gobasic is very simple, and just requires the name of a BASIC-program to
execute.  Write your input in a file and invoke `gobasic` with the path.

For example the following program was useful to test my implementation of the `GOTO` primitive:

     10 GOTO 80
     20 GOTO 70
     30 GOTO 60
     40 PRINT "Hello, world!\n"
     50 END
     60 GOTO 40
     70 GOTO 30
     80 GOTO 20

Execute it like this:

    $ gobasic examples/10-goto.bas

**NOTE**: I feel nostalgic seeing keywords in upper-case, but `PRINT` and `print` are treated identically.


<br />
<br />
<br />
<br />


## 50 PRINT "Implementation"

A traditional interpreter for a scripting/toy language would have a series of
well-defined steps:

* Split the input into a series of tokens ("lexing").
* Parse those tokens and build an abstract syntax tree (AST).
* Walk that tree, evaluating as you go.

As is common with early 8-bit home-computers this implementation is a little more BASIC:

* We parse the input into a series of tokens, defined in [token/token.go](token/token.go)
  * The parsing happens in [tokenizer/tokenizer.go](tokenizer/tokenizer.go)
* We then __directly__ execute those tokens.
  * The execution happens in [eval/eval.go](eval/eval.go) with a couple of small helpers:
    * [eval/for_loop.go](eval/for_loop.go) holds a simple data-structure for handling `FOR`/`NEXT` loops.
    * [eval/stack.go](eval/stack.go) holds a call-stack to handle `GOSUB`/`RETURN`
    * [eval/vars.go](eval/vars.go) holds all our variable references.
    * We have a facility to allow golang code to be made available to BASIC programs, and we use that facility to implement a bunch of our functions as "builtins".
      * Our builtin-functions are implemented beneath [builtin/](builtin/).
* Because we support both strings and ints/floats in our BASIC scripts we use a wrapper to hold them on the golang-side.  This can be found in [object/object.go](object/object.go).

As there is no AST step errors cannot be detected prior to the execution of programs - because we only hit them after we've started running.


<br />
<br />
<br />
<br />


## 60 PRINT "Sample Code"

There are a small number of sample-programs located beneath [examples/](examples/).   These were written in an adhoc fashion to test various parts of the implementation.

Perhaps the best demonstration of the code are the following two samples:

* [examples/90-stars.bas](examples/90-stars.bas)
  * Prompt the user for their name and the number of stars to print.
  * Then print them.  Riveting!  Fascinating!  A program for the whole family!
* [examples/55-game.bas](examples/55-game.bas)
  * A classic game where you guess the random number the computer has thought of.

<br />
<br />
<br />
<br />


## 70 PRINT "Embedding"

The interpreter is designed to be easy to embed into your application(s)
if you're crazy enough to want to do that!

You can see an example in the file [embed/main.go](embed/main.go).

The example defines several new functions which can be called by BASIC:

* `PEEK`
* `POKE`
* `PLOT`
* `SAVE`
* `CIRCLE`

When the script runs it does some BASIC variable manipulation and it also
creates a PNG file - the `PLOT` function allows your script to set a pixel
and the `CIRCLE` primitive draws an outline of a circle.  Finally the
`SAVE` function writes out the result.

Extending this example to draw filled circles, boxes, etc, is left as an
exercise ;)

Hopefully this example shows that making your own functions available to
BASIC scripts is pretty simple.  (This is how SIN, COS, etc are implemented
in the standalone interpreter.)


<br />
<br />
<br />
<br />


## 80 PRINT "Visual BASIC!"

Building upon the code in the embedded-example I've also implemented a simple
HTTP-server which will accept BASIC code, and render images!

To run this:

    cd goserver ; go build . ; ./goserver

Once running open your browser at the URL:

* [http://localhost:8080](http://localhost:8080)

The view will have an area of entering code, and once you run it the result will
be shown in the bottom of the screen.  Something like this:

![alt text](https://github.com/skx/gobasic/raw/master/goserver/screenshot.png "Sample view")

There are several included examples which you can load/launch by clicking upon them.


<br />
<br />
<br />
<br />


## 90 PRINT "Bugs?"

It is probable that bugs exist in this interpreter, but I've tried to do
as much testing as I can.  If you spot anything that
seems wrong please do [report an issue](https://github.com/skx/gobasic/issues).

* If the interpreter segfaults that is a bug.
  * Even if the program is invalid, bogus, or malformed the interpreter should cope with it.
* If a valid program produces the wrong output then that is also a bug.

The project contain a number of test-cases, which you can execute like so:

    $ go test ./...

Finally __if our test-coverage drops beneath 95%__ that is __a bug__.  The
test coverage of most of our packages is 100%, unfortunately the main `eval/`
package is not yet completely covered.

You can see the __global__ coverage via:

    $ ./test-coverage
    97.9%

In addition to the test-cases which have been manually written the interpreter
has also been fuzz-tested, which has resulted in some significant improvements.

See [FUZZING.md](FUZZING.md) for details of how to run the fuzz-tests.


<br />
<br />
<br />
<br />


## 100 PRINT "Project Goals / Links"

It is never the intention of this project to support _all_ things that are
possible in the various dialects of BASIC.

There are facilities which will make porting programs useful, such as
the ability to use `WHILE`/`END` loops, functions with named-parameters,
and primitives such as SLEEP, BEEP, & etc.

Above all else this project is supposed to be fun, for me.  Which means
if there are two ways of implementing something I'll pick the way I remember
back when I was 12 and computers were .. fun!

If there are feature-requests which seem less fun, and less immediately
useful to me - with my biased memories of coding on a ZX Spectrum - I will
tag them "wontfix".  If you contribute a pull-request to support them I will
accept them, but I'm probably not likely to work upon them directly.

That said there are cases where I can be persuaded, and there are a lot
of other BASIC intepreters out there, so I won't feel bad if this particular
project doesn't suit your needs.

One project, slightly related to this, which might be worth checking up
on is this one:

* https://github.com/udhos/basgo


## Github Setup

This repository is configured to run tests upon every commit, and when
pull-requests are created/updated.  The testing is carried out via
[.github/run-tests.sh](.github/run-tests.sh) which is used by the
[github-action-tester](https://github.com/skx/github-action-tester) action.

Releases are automated in a similar fashion via [.github/build](.github/build),
and the [github-action-publish-binaries](https://github.com/skx/github-action-publish-binaries) action.


Steve
--

## Intersting articales and Guides on *BASIC* 

[Fifty Years of *BASIC*, the Programming Language That Made Computers Personal](https://web.archive.org/web/20150708131957/http://time.com/hive.org/web/20150708131957/http://time.com/69316/basic/)

[*BASIC* Programming](https://en.wikibooks.org/wiki/BASIC_Programming)

[Begginers Guide to *BASIC*](http://www.hoist-point.com/applesoft_basic_tutorial.htm)