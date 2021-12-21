JAVA = java
RUNFLG = -jar
JAVAC = javac
JAR = jar
SOURCES = $(wildcard *.java)
CLASSES = $(SOURCES:.java=.class)
MANIFEST = manifest.txt
TARGET = wr.jar
ENCODING = -encoding utf-8 

all: $(TARGET)
	rm *.class

$(TARGET): $(CLASSES)
	$(JAR) -cvmf $(MANIFEST) $@ $^

%.class: %.java
	$(JAVAC) $(ENCODING) $<

clear:
	rm *.class
	rm *.jar

run: $(TARGET)
	$(JAVA) $(RUNFLG) $^