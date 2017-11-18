run:
	./run.sh

help:
	go run main.go -h

test:
	go test -v

build: main.go
	go build -o bin/vlookup

cross-compile:
	gox --output bin/{{.Dir}}_{{.OS}}_{{.Arch}}

.PHONY: clean
clean:
	rm -r bin/*
