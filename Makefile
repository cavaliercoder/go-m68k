all: check m68k

m68k:
	cd cmd/m68k; make

testdata:
	cd testdata/; make

check: testdata
	go test -cover ./...

clean:
	cd cmd/m68k; make clean
	cd testdata/; make clean

.PHONY: all testdata check clean
