build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size 

run:
	./bin/hexlet-path-size $(ARGS)

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix
