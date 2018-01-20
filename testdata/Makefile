################################################################################
# This Makefile uses the 68000 assembler from BSVC to compile object and list  #
# files from 68000 assembler files.                                            #
#                                                                              #
# To obtain a copy of the 68kasm compiler, see: https://github.com/BSVC/bsvc   #
################################################################################
SOURCES = $(wildcard *.x68)
PROGRAMS = $(patsubst %.x68, %.h68, $(SOURCES))
LISTFILES = $(patsubst %.x68, %.lis, $(SOURCES))

all: $(PROGRAMS) $(LISTFILES)

%.h68: %.x68
	68kasm $^

%.lis: %.x68
	68kasm -l -n $^

clean:
	rm -vf $(PROGRAMS) $(LISTFILES)

# more samples: https://www.xrmx.com/solutions/software/68k-fe/samples/index.html
