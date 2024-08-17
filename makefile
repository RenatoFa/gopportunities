APP_NAME=gopportunities

default: run

run:
		@go run main.go

build:
		@go build -o $(APP_NAME) main.go

docs:
		@swag init

clean:
		@rm -f $(APP_NAME)
		@rm -rf ./docs