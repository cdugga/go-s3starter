.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./s3starter/s3starter
	
build:
	GOOS=linux GOARCH=amd64 go build -o s3starter/s3starter ./s3starter