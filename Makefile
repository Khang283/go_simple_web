build :
	go build -o bin/crud cmd/main.go
run : build
	./bin/crud
go_test : 
	go test -v ./test
