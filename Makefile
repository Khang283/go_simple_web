build :
	go build -o bin/crud cmd/main.go
run : build
	./bin/crud
