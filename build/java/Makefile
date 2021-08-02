JAVA = java
RUNFLG = -jar
JAVAC = javac
JAR = jar
SOURCES = $(wildcard *.java)
CLASSES = $(SOURCES:.java=.class)
MANIFEST = manifest.txt
TARGET = wr.jar

all: $(TARGET)

$(TARGET): $(CLASSES)
	$(JAR) -cvmf $(MANIFEST) $@ $^

%.class: %.java
	$(JAVAC) $<

clear:
	rm *.class
	rm *.jar

run: $(TARGET)
	$(JAVA) $(RUNFLG) $^