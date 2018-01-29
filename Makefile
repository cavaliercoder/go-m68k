all: check m68k

m68k:
	cd cmd/m68k; make

m68kgen:
	cd cmd/m68kgen; make

opcodes.go: m68kgen
	go generate -x

testdata:
	cd testdata/; make

check: opcodes.go testdata
	go test -cover ./...

clean:
	cd cmd/m68k; make clean
	cd cmd/m68kgen/; make clean
	cd testdata/; make clean
	rm -f opcodes.go

.PHONY: all m68kgen testdata check clean
