# **basicbots** code library

A curated list of subroutines to help people get started.


## Get the angle to a target
- Variables passed
  - DX = target X
  - DY = target Y
- Variables returned
  - ANTT = Angle to target
```basic
9000 LET LLX = LOCX
9005 LET LLY = LOCY
9010 LET ANTT = ATN2 DY - LLY, DX - LLX
9015 RETURN
```

## Get the distance to target
- Variables passed
  - DX = target X
  - DY = target Y
- Returned variables
  - DTT = Distance to target
```basic  
9050 LET LX = LOCX
9055 LET LY = LOCY
9060 LET TX = DX - LX
9065 LET TY = DY - LY
9070 LET DTT = SQR (TX*TX) + (TY*TY)
9075 RETURN
```

## Pick random location in the battlefield. Staying safely away from the walls.
- Variables passed
  - None
- Returned variables
  - DX = Destination X
  - DY = Destination Y
```basic
9100 LET DX = (RND 700) + 100
9105 LET DY = (RND 700) + 100
9110 RETURN
```

## Pick a random location on the battlefield and move towards it at full speed, stopping when close.

```basic
9200 GOSUB 9100
9210 GOSUB 9000
9220 DRIVE ANTT, 100
9230 GOSUB 9050
9240 IF DTT <= 50 THEN GOTO 9260
9250 GOTO 9230
9260 DRIVE ANTT,0
9270 RETURN
```

## Pick a random location on the battlefield and move towards it at full speed. Scan and shoot while moving.
- Variables passed
  - SCA = Starting scan angle
  - SCW = Scanning width
- Returned variables
  - None

```basic
9300 GOSUB 9100
9310 GOSUB 9000
9320 DRIVE ANTT, 100
9340 GOSUB 9900
9350 GOSUB 9050
9360 IF DTT <= 50 THEN GOTO 9380
9370 GOTO 9340
9380 DRIVE ANTT, 0
9390 RETURN 
```

## Simple scan and shoot. Keep scan angle between calls.
- Variables passed
  - SCA = Starting scan angle
  - SCW = Scanning width
- Returned variables
  - SCA = New starting scan angle

```basic
9900 LET SDTT = SCAN SCA,SCW
9905 IF SDTT > 0 AND SDTT < 700 THEN GOTO 9920
9910 LET SCA = SCA + SCW
9915 RETURN
9920 CANNON SCA,SDTT
9930 GOTO 9900
```

## Try to lead the shot. Call after first scan.
- Variables passed
  - SCA = Scan angle
  - SCW = Scan width
- Returned variables
  - none

```basic
9970 LET DTTN = SCAN SCA,SCW
9975 LET DIF = DTTN + (DTTN - SDTT)
9980 CANNON RSA, DIF
9985 RETURN
```
