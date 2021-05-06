default:
	go fmt

run: default
	go run ./*.go

linux: default
	go build -o build/dark-fiber ./*.go