		* tests for ori
		org $1000

    * set switch on all CCR flags
    ori.b #$FF,CCR

    * set D1 to 0xFFFFFFFF, one bit at a time
    ori.b #$01,D1
    ori.b #$02,D1
    ori.b #$04,D1
    ori.b #$08,D1
    ori.b #$10,D1
    ori.b #$20,D1
    ori.b #$40,D1
    ori.b #$80,D1
    ori.w #$0100,D1
    ori.w #$0200,D1
    ori.w #$0400,D1
    ori.w #$0800,D1
    ori.w #$1000,D1
    ori.w #$2000,D1
    ori.w #$4000,D1
    ori.w #$8000,D1
    ori.l #$010000,D1
    ori.l #$020000,D1
    ori.l #$040000,D1
    ori.l #$080000,D1
    ori.l #$100000,D1
    ori.l #$200000,D1
    ori.l #$400000,D1
    ori.l #$800000,D1
    ori.l #$01000000,D1
    ori.l #$02000000,D1
    ori.l #$04000000,D1
    ori.l #$08000000,D1
    ori.l #$10000000,D1
    ori.l #$20000000,D1
    ori.l #$40000000,D1
    ori.l #$80000000,D1
    ori.l #$0,D1 * try break it

    * Set D2 to 0x33
    move.b #$22,D2
    ori.b #$11,D2

    * Set D3 to 0x3333
    move.w #$2222,D3
    ori.w #$1111,D3

    * Set D3 to 0x33333333
    move.l #$22222222,D4
    ori.l #$11111111,D4