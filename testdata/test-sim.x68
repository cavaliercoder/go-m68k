* Test 68000 simulator program

WRSTRZ  EQU     $0E                     Trap # for write A1 string to screen

        ORG $1000

        LEA     MESSAGE(PC),A1          Point A0 to start of message
        MOVE.B  #WRSTRZ,D0              Set up trap to write to screen
        TRAP    #15                     Write char. to screen

FINISH  STOP    #$2700                  Halt
MESSAGE DC.B    'Hello world!',$0D,$0A,0

        END     $1000