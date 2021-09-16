# basicbots overview and guide

This is the overview manual for **basicbots** programming game; with examples that will hopefully start you on your way. 

--- 

## The simulation

### The Battlefield
The robots compete in a virtual battlefield that is 1000 units wide by 1000 units long. The upper left corner is 0,0 while the lower right corner is 1000,1000. There is a virtual wall that surrounds the battlefield. If a robot makes contact with this wall, it will take damage and its speed will be reduced to 0.

### Cardinal directions
|     |     |     |
|:----|:---:| ---:|
| 225 | 270 | 315 |
| 180 |  *  | 000 |
| 135 | 090 | 045 |
|     |     |     |


### The Robot's Hardware
The robot tank has three hardware systems.

- **A Scanner**. The scanner can scan in any direction instantly but has an angular resolution from +-2 degrees to +- 10 degrees. Example: Scan 90 degrees with a width of 2 would return results that are between 88 - 92 degrees. A width of 10 would return results that are from 80 to 100 degrees.

- **A Cannon**. The Cannon can fire a projectile in any direction but at a maximum range of 700 units. Any robot caught in the blast radius of 40 units. The farther from the center of the blast, the less damage a robot will take. Each robot can only have two projectiles in the air at any given time and it takes time to reload. 

- **A Drive system**. The drive system has two parameters. Angle and speed. The angle can be any angle from 0 to 360 degrees and the speed can be from 0 to 100%. The robots do require time to reach a given speed and require time to slow down. The robot can only negotiate turns when the speed is less than 50%. If the speed is above 50% the robot will slow down to 50% then turn and increase speed to the previous setting.

### The Robot's Status
The robot has access to:
- **The X location**. The current X position on the battlefield of the robot.
- **The Y location**. The current Y position on the battlefield of the robot.
- **The Speed**. The current speed of the robot. 
- **The Damage**. The current amount of damage inflicted upon the robot. The robot can withstand up to 100 points of damage. Once a robot has reached 100 points, it has died and is removed from the battlefield. The amount of damage does not affect its performance.

That's it. Three commands and four status sensors. With that, you can navigate the battlefield and annihilate the other robots!



# The *BASIC* of **basicbots**

Keywords

| Command |      |      |  |  |  | Notes |
|:--------|:-----|:-----|:-|:-|:-| :------|
| REM     | Single line | | | | | A remark |
| LET     | Variable| = Value | | | | Assign a value to a variable |
| GOTO    | Line number| | | |  | Transfer execution to the line number |
| GOSUB   | Line number| |  | | | Transfer execution to the subroutine at line number |
| RETURN  | | | | |  | Transfers execution to the line following the GOSUB command|
| IF | condition | THEN | statement | ELSE | statement | Excute the statement after THEN if condition is true if not then execute statement after ELSE
| FOR | assignment | TO | value | STEP | value | Set the loop paramters
| NEXT | variable | | | | | Return back to FOR changing the value of the assignment by the value in STEP |
| DIM | variable | [size] | | | | Create an array of the size specified. Note this array is 0 based unlike other BASICs.|
| DEF | FN | name | expresstion | | | Create a user defined function |
| FN | name | values | | | | execute defined funtion name |
| SWAP | variable1 | variable2| | | | Exchange the values variable1 and variable2 |
| DATA | values, | | | | | List of values to be read |
| READ | | | | | | Reads the data stored by DATA. There is no reset ability at this time.|
| AND | value1 | value2 | | | | Bit AND of the value1 and value2. See [truth table](#truth-table) |
| OR | value1 | value2  | | | | Bit OR of the value1 and value2. See [truth table](#truth-table) |
| XOR | value1 | value2 | | | | Bit XOR or the value1 and value2. See [truth table](#truth-table) |
| LOCX | | | | | | Returns the X Location of the robot.|
| LOCY | | | | | | Returns the Y Location of the robot.|
| DAMAGE | | | | | | Returns the current damage of the robot.|
| SPEED | | | | | | Returns the current speed of the robot.|
| SCAN | angle| width| | | | Scan for an enemy robot in the direction of angle with a width of +- width.
| CANNON | angle | range | | | | Fire the cannon at angle and range |
| DRIVE | angle | speed% | | | |  Engage the drive in the direction of angle with a speed of speed%








--- 

### Truth table

|     | VALUE | VALUE | RESULT |
|:----|:------|:------|:-------|
| AND | 0     | 0     | 0      |
| AND | 1     | 0     | 0      |
| AND | 0     | 1     | 0      |
| AND | 1     | 1     | 1      |

    48 AND 12 = 0
    0000000000110000 = 48
    0000000000001100 = 12
    0000000000000000 = 0

    48 AND 16 = 16
    0000000000110000 = 48
    0000000000010000 = 16
    0000000000010000 = 16

|     | VALUE | VALUE | RESULT |
|:----|:------|:------|:-------|
| OR  | 0     | 0     | 0      |
| OR  | 1     | 0     | 1      |
| OR  | 0     | 1     | 1      |
| OR  | 1     | 1     | 1      |

    48 OR 12 = 60
    0000000000110000 = 48
    0000000000001100 = 12
    0000000000111100 = 60

    48 OR 16 = 48
    0000000000110000 = 48
    0000000000010000 = 16
    0000000000110000 = 48

|     | VALUE | VALUE | RESULT |
|:----|:------|:------|:-------|
| XOR | 0     | 0     | 0      |
| XOR | 1     | 0     | 1      |
| XOR | 0     | 1     | 1      |
| XOR | 1     | 1     | 0      |

    48 XOR 12 = 60
    0000000000110000 = 48
    0000000000001100 = 12
    0000000000111100 = 60

    48 XOR 16 = 32
    0000000000110000 = 48
    0000000000010000 = 16
    0000000000100000 = 32

## Whats next?

