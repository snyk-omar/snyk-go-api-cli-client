build:
	go build -o bin/Client
	make test

run:
	go run main.go

test:
	./bin/Client

clean:
	rm bin/*