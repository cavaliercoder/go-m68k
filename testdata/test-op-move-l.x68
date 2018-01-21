		* tests for move.l
		org $1000

		*
		* source: data register
		*
		move.l #$44304430,D0
		move.l D0,D1							* to data register
		movea.l D0,A1							* to address register

		movea.l #$2000,A0
		move.l D0,(A0)						* to indirect address

		move.l #$44314431,D0
		movea.l #$2004,A0
		move.l D0,(A0)+						* to indirect with post-increment

		move.l #$44324432,D0
		move.l #$FFFFFFFF,(A0)+
		move.l D0,-(A0)						* to indirect with pre-decrement

		move.l #$44334433,D0
		move.l D0,$200C						* to absolute short
		move.l D0,$12000					* immediate to absolute long

		*
		* source: address register
		*
		move.l #$41304130,A0
		move.l A0,A2							* to address register
		move.l A0,D2							* to data register

		movea.l #$3000,A6
		move.l A0,(A6)						* to indirect address

		movea.l #$41314131,A0
		movea.l #$3004,A6
		move.l A0,(A6)+						* to indirect with post-increment

		movea.l #$41324132,A0
		move.l #$FFFFFFFF,(A6)+
		move.l A0,-(A6)						* to indirect with pre-decrement

		movea.l #$41334133,A0
		move.l A0,$300C						* to absolute short
		move.l A0,$13000					* immediate to absolute long


		*
		* source: immediate long
		*
		movea.l #$49304930,A3			* to address register
		move.l #$49304930,D3			* to data register

		movea.l #$F000,A0
		move.l #$49304930,(A0)		* to indirect address

		movea.l #$F004,A0
		move.l #$49314931,(A0)+		* to indirect with post-increment
		move.l #$FFFFFFFF,(A0)+
		move.l #$49324932,-(A0)		* to indirect with pre-decrement

		move.l #$49334933,$F00C		* immediate to absolute short
		move.l #$49334933,$1F000	* immediate to absolute long
