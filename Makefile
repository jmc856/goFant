install:
	- go get github.com/stretchr/testify/assert
	- go get github.com/lib/pq"

migrations:
	go run models/models.go
test:
	go test ./...
run:
	go run main.go
