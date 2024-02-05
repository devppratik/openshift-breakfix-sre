install:
	go build -o bin/main main.go

run:
	go run main.go

build:
	go build .

dev:
	go build && ./openshift-breakfix