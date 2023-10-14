TITLE?="znschaffer.com"
URL?="https://znschaffer.com"
IGNORES=-I Makefile -I index.html -I CNAME -I fonts -I css.css
all:
	tree . -T $(TITLE) -H $(URL) -n $(IGNORES) > index.html

clean:
	rm -rf index.html
