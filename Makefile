all:	build example

build:
	go build ./...

example:	examples/example.go
	go build -o example ./examples

clean:
	$(RM) example