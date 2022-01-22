build:
	go build -o ./bin/head ./src/main.go

clean:
	rm ./bin/head

run: build
	./bin/head -h
