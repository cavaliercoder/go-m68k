all: check

testdata:
	cd testdata/; make

check:
	go test -v -cover ./...

.PHONY: all testdata check
