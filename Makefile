all: check

m68kgen:
	cd cmd/m68kgen; make

opcodes.go: m68kgen
	go generate -x

testdata:
	cd testdata/; make

check: opcodes.go testdata
	go test -v -cover ./...

clean:
	cd cmd/m68kgen/; make clean
	cd testdata/; make clean
	rm -f opcodes.go

.PHONY: all m68kgen testdata check clean
