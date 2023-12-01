.PHONY: up

faktory-up:
	docker-compose up faktory -d
	echo "Waiting for FAKTORY to start..."
	sleep 5

faktory-down:
	docker-compose stop faktory

run: faktory-up
	go run main.go

run-server-standalone: faktory-up
	go run main.go server

run-worker-standalone: faktory-up
	go run main.go worker

docker-build:
	docker build . -f docker/Dockerfile --target dev

up: 
	docker-compose up -d

down: 
	docker-compose down