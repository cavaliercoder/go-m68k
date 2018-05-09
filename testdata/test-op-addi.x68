		* tests for addi
		org $1000

		* add byte
		addi.b #$FF,D0					; add 0xFF to D0
		addi.b #$FF,D0				; overflow D0 to 0x01

		* add word
		addi.w #$FFFF,D1			; add 0xFFFF to D1
		addi.w #$FFFF,D1			; overflow D1 to 0x0001

		* add long
		addi.l   #$FFFFFFFF,D2	; add 0xFFFFFFFF to D2
		addi.l   #$FFFFFFFF,D2	; overflow D2 to $0x00000001

		* ensure no overflow
		addi.b	#$FF,D3					; add 0xFF to D3
		addi.w	#$FF00,D3				; add 0xFF00 to D3
		addi.l	#$FFFF0000,D3		; add 0xFFFF0000

		stop #$2700
