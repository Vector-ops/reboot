APP_NAME = reboot


build: 
	@go build -o bin/$(APP_NAME) ./main.go

run: build
	@./bin/$(APP_NAME)

clean:
	@rm -rf ./bin

test:
	@go test -v ./...