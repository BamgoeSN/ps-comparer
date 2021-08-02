CC = g++
CFLAGS = -O2 -Wall -lm -static -std=gnu++17
TARGET = wr.exe
SOURCES = $(wildcard *.cpp)

all: $(TARGET)

$(TARGET): $(SOURCES)
	$(CC) $^ -o $@ $(CFLAGS)

clear:
	rm *.exe