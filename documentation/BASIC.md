# **basicbots** *BASIC*

**basicbots** version of *BASIC* is based and adapted from [gobasic](https://github.com/skx/gobasic) by [James Kemp](https://github.com/skx). Mr. Kemp has gracely given me permission to adapt [gobasic](https://github.com/skx/gobasic) for my game.


## What is *BASIC*
*BASIC* stands for Beginners' All-purpose Symbolic Instruction Code

*BASIC* was developed at Dartmouth College in 1964 by John G. Kemeny and Thomas E. Kurtz. It was designed with the philosophy of ease of use. It allowed students in fields other than science and mathematics to use computers. It has a long lasting leagcy spanning mainframes to personal computers and now the IoT. 

---

## **basicbots** version

**basicbot** uses a subset of the *BASIC* language. Most of the early functionality is present with addition of commands for controlling the robots. 

I encourage you to seek out online guides.

### Qurkes
- Variables are case sensitive. Primitives are not.
  - LOCALX is different than localx
- Parens around statemenst are not allowed. Example, ```LET A = SQR(25)``` is an error. Where as ```LET A = SQL 25``` is correct.
- Line numbers are not strictly required but are strongly suggested.
  - I do not suggest leaving out the line numbers. Error handling is still awful and needs work. Doing this would just make it harder to debug your program.
- Some of the custom statements seem not to play well with `IF` `THEN` & `ELSE`. I am not sure why but for now keep the `IF` states as simple as you can. If you get odd error messages. Rewrite the `IF` states around it.
- Error reporting can be a bit difficult to interpret. 
  - Lot at the line after the reported error for possible problems.
  - Advice, save work often and make copies as you go. **git** is a good way to handle this.
- I find it useful, at the top of the program to initialize all the variables. This cuts down on some errors.

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
  - `FOR` / `NEXT` / `STEP` / `TO`

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
  - Instant death to a robot.

~~PRINT output-list~~

~~Produces output to the console. Output-list is a list of items separated by commas. Each item can be either a string literal enclosed in double quotation marks or a numeric expression. An end of line sequence is output after all the values so that the next PRINT statement will put its output on a new line.~~

~~INPUT variable-list~~

~~Asks for input from the console. Variable-list is a list of variable names. For each variable given, a question mark is output, and the value typed by the user is stored in that variable. Tinybasic allows multiple values to be typed by the user on one line, each separated by any non-numeric character.~~

- `REM` comment-text
  - Provides space for free-format comment text in the program. Comments have no effect on the execution of a program and exist only to provide human-readable information to the programmer. 

### EXPRESSIONS

Expressions in **basicbots** *BASIC* are purely arithmetic expressions. The four basic arithmetic operators are supported: multiplication (*), division (/), addition (+), subtraction (-) and (%) modulus. Unary operators for positive (+) and negative (-) are supported, as are parentheses for affecting the order of operations.

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

That's a lot to take in. In the next section we will break it down by using real examples. 

---

## Intorduction to *BASIC* programming.



### Assignment
`LET` assignes a number or result to a variable. `LET` will be used often in most programs.

- `10 LET A = 0`
  - Would assign the number 0 to the variable named A. 
- `10 LET A = A + 1`
  - Would increment the variable by adding 1 to A and storing the results back in A. 
- `10 LET DISTANCE = SCAN A, 5`
  - The results of a `SCAN` at angle `A` with a width of 5 would be stored in the variable `DISTANCE'

### Conditionals
- `>` more than. 
  - `IF A > B THEN LET A = A + 10` 
    - if `A` is more than `B` then add `10` to the variable `A`
- `<` lets than
  - `IF A < B THEN LET A = A - 10` 
    - if `A` is less than `B` then subtract `10` from the variable `A`
- `=` equals
  - `IF A = B THEN GOTO 100 ELSE GOTO 200`
    - if `A` is equal to `B` then jump to location `100` else jump to location `200`
- `>=` more than or equals 
  - `IF A >= B THEN LET A = A + 10` 
    - if `A` is more than or equals `B` then add `10` to the variable `A`
- `<=` lets than or equals
  - `IF A <= B THEN LET A = A - 10` 
    - if `A` is less than or equals `B` then subtract `10` from the variable `A`
- `<>` does not equal
  - `IF A <> B THEN GOSUB 200`
    - If `A` does not equal B then `GOSUB` to location 200
- no conditional
  - `IF A THEN LET B = A`
    - If `A` is not 0 then assign the value in `A` to the variable `B`
    - If `A` is any number expect 0 then ...
  
## Loops

Loops allow an effecient way of doing a section of code over a set number of times.

	5 LET B = 0
	10 FOR A = 1 TO 10
	20 LET B = B + 1
	25 PRINT B,
	30 NEXT A
	35 PRINT "\n"
	40 END


Will output 

	1 2 3 4 5 6 7 8 9 10

You can control the step that each iteration does,

	10 FOR A = 1 TO 10 STEP 2
	20 PRINT A,
	30 NEXT A
	35 PRINT "\n"
	40 END

Will output 
    
	1 3 5 7 9

Not wha you expected? We started the loop at 1, so 1 + 2 = 3. You can change this by starting your loop at 0.

	10 FOR A = 0 TO 10 STEP 2
	20 PRINT A,
	30 NEXT A
	35 PRINT "\n"
	40 END


will output

	0 2 4 6 8 10 

Loops can also step by a fraction.

	10 FOR A = 0 TO 1 STEP 0.2
	20 PRINT A,
	30 NEXT A
	35 PRINT "\n"
	40 END

will output

	0 0.200000 0.400000 0.600000 0.800000 1









  ## Lets get started with *basicbots* programming.

  Actual working code! Featuring a subroutine. 

```
1 REM scan and shoot
10 LET DISTANCE = 0
11 LET ANGLE = 0
12 LET WIDTH = 5

100 GOSUB 1000
110 IF DISTANCE < 700 THEN CANNON ANGLE,DISTANCE
120 GOTO 100 

1000 REM scan for enemy robot.
1010 LET DISTANCE = SCAN ANGLE, WIDTH
1020 IF DISTANCE > 0 THEN GOTO 1100
1030 LET ANGLE = ANGLE + WIDTH
1040 IF ANGLE >= 360 THEN LET ANGLE = ANGLE - 360
1050 GOTO 1010
1100 RETURN

```



The flow is
- initialize all variables
- jump to subroutine 
  - scan for enemy robot 
  - if found return frome subroutine
    - if not found add the width size to the angle
  - jump back to scan command
  - return from subroutine
- if the distance to the enemy is less than 700 then shot cannon at the angle of the last scan with the distance returned
- do it all again





---
So much more to do.

---
---

## Interesting articles and Guides on *BASIC* 

[Fifty Years of *BASIC*, the Programming Language That Made Computers Personal](https://web.archive.org/web/20150708131957/http://time.com/hive.org/web/20150708131957/http://time.com/69316/basic/)

[*BASIC* Programming](https://en.wikibooks.org/wiki/BASIC_Programming)

[Begginers Guide to *BASIC*](http://www.hoist-point.com/applesoft_basic_tutorial.htm)

