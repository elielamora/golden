test:
	go test -v ./...

golden:
	UPDATE_GOLDEN=true go test -v ./...