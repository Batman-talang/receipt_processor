build:
	go build -o bin/ags cmd/server/main.go
	cp .env bin/

run: build
	./bin/ags

build-and-run:
	go build -o bin/ags cmd/server/main.go
	cp .env bin/
	./bin/ags

test:
	go test ./...


up:
	docker compose up --build

