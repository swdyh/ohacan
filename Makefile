fmt:
	gofmt -d **/*.go
lint:
	golint
clean:
	go clean -i -n github.com/swdyh/ohacan/ohacan
install:
	go install github.com/swdyh/ohacan/ohacan
test:
	go test
