RUST = rustc
CFLAGS = --edition 2018 -O
SOURCES = $(wildcard *.rs)
TARGET = wr.exe

all: $(TARGET)

$(TARGET): $(SOURCES)
	$(RUST) $(CFLAGS) -o $@ $^
	rm *.pdb

clear:
	rm *.exe