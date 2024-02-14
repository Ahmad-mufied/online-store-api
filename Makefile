BINARY_NAME=online-store
DOCKER_IMAGE=ahmadryzen/online-store:latest
## build: Build binary
build:
	@echo "Building..."
	env CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o ${BINARY_NAME} .
	docker build -f Dockerfile -t ${DOCKER_IMAGE} .
	@echo "Built!"

## push: pushes the docker image to the registry
push: build
	@echo "Pushing..."
	docker push ${DOCKER_IMAGE}
	@echo "Pushed!"

## run: builds and runs the application
run: build
	@echo "Starting..."
	@docker-compose up -d
	@echo "Started!"

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm ${BINARY_NAME}
	@echo "Cleaned!"

## start: an alias to run
start: run
## down: stop and remove docker-compose services
stop:
	@echo "Stopping and removing services..."
	@docker-compose down
	@echo "Services stopped and removed!"

## restart: stops and starts the application
restart: stop start
