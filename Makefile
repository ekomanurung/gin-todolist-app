BINARY=engine
test:
		go test -v -cover -covermode=atomic ./...

unittest:
		go test -short ./...

run:
		sh ./deployment.sh

stop:
		sh ./rollback.sh

.PHONY: unittest test run stop