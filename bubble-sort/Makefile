all: build

#build: c golang
build:
	$(MAKE) -C c
	$(MAKE) -C golang

golang: golang/bubble
	$(MAKE) -C golang 

clean: c-clean golang-clean

c-clean:
	$(MAKE) -C c clean

golang-clean:
	$(MAKE) -C golang clean
