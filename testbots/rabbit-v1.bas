1 REM rabbit-v1 moves around battle field randomly. Used as moving target for testing.

5 LET HIGHSPEED = 100
6 LET SLOWSPEED = 25

10 REM PICK A NEW LOCATION
11 LET dx = (RND 800 ) + 100 : REM destination X
15 LET dy = (RND 800 ) + 100 : REM destination Y
16 REM PRINT "NEW LOCATION ",dx,dy,"\n"

20 LET x = LOCX : REM current location X
25 LET y = LOCY : REM current location Y
26 PRINT "RABBIT:",x,y,"\n"
30 LET a = ATN2 dy-y,dx-x : REM get angle to target
31 REM PRINT x,y,dx,dy,a,"\n"
40 DRIVE a,HIGHSPEED
50 GOSUB 2010 
60 if d > 100 THEN GOTO 50 
70 REM slow down and creep to destination
80 DRIVE a,SLOWSPEED
90 GOSUB 2010
100 if d > 10 THEN GOTO 90
110 REM Arrived. Shut drive down.
110 DRIVE a,0
120 GOTO 10 : REM pick a new destination

1000 LET XX = LOCX
1002 LET YY = LOCY
1010 IF XX <= dx THEN LET XR = XX/dx*100
1015 IF YY <= dy THEN LET YR = YY/dx*100
1020 IF XX > dx THEN LET XR = dx/XX*100
1025 IF YY > dy THEN LET YR = dx/YY*100
1027 REM print XR,YR,"\n"
1050 IF XR > 90 AND YR > 90 THEN RETURN
1060 GOTO 1000

2000 REM Distance routine
2001 REM Destination is dx,dy
2010 LET cx = LOCX
2012 LET cy = LOCY
2014 LET tx = cx-dx
2016 LET ty = cy-dy
2020 LET d = SQR (tx*tx)+(ty*ty)
2030 RETURN
