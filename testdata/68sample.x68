*       68Sample.x68
*       By Greg Colley
*       Program which adds 3 numbers together and subtracts another.
LOC1    EQU     $2000
LOC2    EQU     $2002
LOC3    EQU     $2004

RESULT  EQU     $2006   Answer Location

        ORG     $1000

*       Main Program
        MOVE.W  LOC1,D0
        ADD.W   LOC2,D0
        ADD.W   LOC3,D0
        SUB.W   #$1F0,D0
        MOVE.W  D0,RESULT
        STOP    #$2700

        END     $1000
