SOURCES = $(wildcard *.go)

all: m68kgen

m68kgen: $(SOURCES)
	go build -o $@ $^

clean:
	rm -f m68kgen

.PHONY: all clean
