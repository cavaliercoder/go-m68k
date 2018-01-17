* Test 68000 simulator program

WRCHAR  EQU     6                       Trap # for write D1.B char to screen

        ORG $1000

        LEA      MESSAGE(PC),A0         Point A0 to start of message
NEXT    MOVE.B  (A0)+,D1                Get character, increment pointer
        BEQ.S   FINISH                  Exit if end
        MOVE.B  #WRCHAR,D0              Set up trap to write to screen
        TRAP    #15                     Write char. to screen
        BRA.S   NEXT                    ..and loop back
FINISH  STOP    #$2700                  Halt
MESSAGE DC.B    'This is the message',$0D,$0A,0

        END     $1000