all: check

testdata:
	cd testdata/; make

check: testdata
	go test -v -cover ./...

.PHONY: all testdata check
