run:
	make build && make ui

build:
	go build -o bin/main pkg/app/* 

ui:
	./bin/main

clean:
	rm bin/main