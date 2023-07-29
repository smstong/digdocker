PROG=minidocker
CFLAGS=-I

%.o: %.c
	gcc -c -o $@ $<

$(PROG): $(PROG).o netns.o
	gcc -o $(PROG) $(PROG).o netns.o

clean:
	rm -f *.o $(PROG)
