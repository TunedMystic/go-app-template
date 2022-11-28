run: build
	./bin/app

clean:
	rm -rf ./bin/app

build: clean
	go build -o ./bin/app ./app

build-small: clean
	go build -ldflags="-s -w" -o ./bin/app ./app

test:
	go test -v ./app
