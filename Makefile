TITLE?="znschaffer.com"
URL?="https://znschaffer.com"
all:
	tree . -T $(TITLE) -H $(URL) -n -I Makefile -I index.html > index.html

clean:
	rm -rf index.html
