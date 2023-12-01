.PHONY: up

up:
	docker-compose -f ./docker/docker-compose.yaml up -d
	echo "Waiting for FAKTORY to start..."
	sleep 5

down: 
	docker-compose -f ./docker/docker-compose.yaml down

run: up
	go run main.go

run-server-standalone: up
	go run main.go server

run-worker-standalone: up
	go run main.go worker