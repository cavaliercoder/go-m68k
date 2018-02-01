SOURCES = $(wildcard *.go)

all: m68k

m68k: $(SOURCES)
	go build -o $@ $^

clean:
	rm -f m68k

.PHONY: all clean
