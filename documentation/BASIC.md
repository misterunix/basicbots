# basicbots *BASIC*

**basicbots** version of *BASIC* is based and adapted from [gobasic](https://github.com/skx/gobasic) by [James Kemp](https://github.com/skx). Mr. Kemp has gracely given me permission to adapt [gobasic](https://github.com/skx/gobasic) for my game.


## BASIC
---
**basicbot** uses a subset of the *BASIC* language. Most of the early BASICs functionality is implemented. The assumption is that the players are passingly familiar with *BASIC*. This document is not meant to teach the *BASIC* programming language. I encourage you to seek out online guides.

**NOT TRUE**

### Qurkes
- Variables are case sensitive. Primitives are not.
  - LOCALX is different than localx
- Line numbers are not strictly required but are strongly suggested.
  - I do not suggest doing leaving out the line numbers. Error handling is still awful and needs work. Doing this would just make it harder.
- Some of the custom statements seem not to play well with `IF` `THEN` & `ELSE`. I am not sure why but for now keep the `IF` states as simple as you can. If you get odd error messages. Rewrite the `IF` states around it.
- Error reporting can be a bit difficult to interpret. Advice, save work often and make copies as you go. **git** is a good way to handle this.
- I find it useful at the top of the program to initialize all the variables. This cuts down on some errors.

### Limitations
A previously stated this is a subset of the *BASIC* standard. 
- There is no editor.
- Errors in the program are only found during execution.
- Only a single statement is allowed on each line. The exception is for `REM`arks. `REM`arks can be placed after ':' at the end of a line.
- The only data types are floating-point and strings.
- Arrays can have a maximum of two dimensions.
- `LET` is not optional.
- Only a single statement is allowed between `THEN` and `ELSE` and a single statement from `ELSE` to a newline.

### Program format

Programs are text files that are specified on the command line. They are loaded and parsed when the robots are initialized. 

Each line contains a line number followed by a statement or command with an optional parameter(s). 

All line numbers are integers.

### basicbots robot statements

- `SCAN` direction, width
  - Scan for other robots. Returning the range if any are found.
  - "direction" is the direction to scan
  - "width" is how wide the scan is. Width is specified as +- width. For example, SCAN 90, 10 would scan @ 90 degrees with a width that spans from 80 to 100 degrees. 
- `CANNON` direction, range
  - Fire the cannon in the direction "direction" with it exploding at the distance specified by "range"
- `SPEED`
  - Returns the current speed of the robot.
- `LOCX`
  - Returns the current X location of the robot on the battlefield.
- `LOCY` 
  - Returns the current Y location of the robot on the battlefield.
- `DRIVE` direction, speed
  - Starts the drive system in the direction of "direction" with the speed specified. Turns can only be made when the speed is 50 or below. If the speed is above 50, the robot will slow down and then turn, returning to the former speed.
- `DAMAGE`
  - Returns the amount of damage the robot has sustained.

### *BASIC* primitives

- Comments 
  - `REM` Single-line comment only.

- Variable assignment
  - `LET` 
  - `DIM`

- Loops
  - `FOR` `NEXT` `STEP` & `TO`

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
  - ABS
  - ACS
  - ATN
  - ATN2
  - BIN
  - COS
  - EXP
  - INT
  - LN
  - LOG
  - PI
  - RND
  - SGN
  - SIN
  - SQR
  - TAN
  - VAL

- `LET` variable = expression
  - Assigns a value, or the result of expression, to a variable. Variables are case sensitive and can be floating-point or strings.

- `IF` condition `THEN` statement `ELSE` statement
  - If the condition is true, then the statement following `THEN` is executed. If it is false then the statement after `ELSE` is executed. 

- `GOTO` line number
  - Transfer execution to another part of the program specified by the line number.

- `GOSUB` line number
  - Transfers execution to a subroutine specified by the line number. The position of the `GOSUB` is remembered so that a `RETURN` can bring program execution back to the statement following the `GOSUB`.

- `RETURN`
  - Returns from a subroutine. Program execution returns to the statement following the `GOSUB`,  which called the present subroutine.

- `END`
  - Terminates program execution.

~~PRINT output-list~~

~~Produces output to the console. Output-list is a list of items separated by commas. Each item can be either a string literal enclosed in double quotation marks or a numeric expression. An end of line sequence is output after all the values so that the next PRINT statement will put its output on a new line.~~

~~INPUT variable-list~~

~~Asks for input from the console. Variable-list is a list of variable names. For each variable given, a question mark is output, and the value typed by the user is stored in that variable. Tinybasic allows multiple values to be typed by the user on one line, each separated by any non-numeric character.~~

- `REM` comment-text
  - Provides space for free-format comment text in the program. Comments have no effect on the execution of a program and exist only to provide human-readable information to the programmer. 

### EXPRESSIONS

Expressions in **basicbots** *BASIC* are purely arithmetic expressions. The four basic arithmetic operators are supported: multiplication (*), division (/), addition (+), and subtraction (-). Unary operators for positive (+) and negative (-) are supported, as are parentheses for affecting the order of operations.

Standard operator precedence evaluates parentheses first, then unary signs, then multiplication and division, with addition and subtraction last.

### CONDITIONS

The relational operators are =, >, <, <> or ><, >=, and <=. They are not supported within arithmetic expressions, but can only be used as conditions in IF statements in the form: expression relational-operator expression

- `IF a < b THEN ..`
- `IF a > b THEN ..`
- `IF a <= b THEN ..`
- `IF a >= b THEN ..`
- `IF a = b THEN ..`
- `IF a <> b THEN ..`
- `IF a THEN ..`
  - This passes if `a` is a number which is not zero.
  - This passes if `a` is a string which is non-empty.



---
So much more to do.
---
---

## Interesting articles and Guides on *BASIC* 

[Fifty Years of *BASIC*, the Programming Language That Made Computers Personal](https://web.archive.org/web/20150708131957/http://time.com/hive.org/web/20150708131957/http://time.com/69316/basic/)

[*BASIC* Programming](https://en.wikibooks.org/wiki/BASIC_Programming)

[Begginers Guide to *BASIC*](http://www.hoist-point.com/applesoft_basic_tutorial.htm)

