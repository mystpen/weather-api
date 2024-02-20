PACKAGE=github.com/mystpen/weathet-api/cmd/app

test:
	go test ./...
	
run:
	go run ${PACKAGE}