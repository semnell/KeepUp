.PHONY: up

up:
	docker-compose -f ./docker/docker-compose.yaml up -d

down: 
	docker-compose -f ./docker/docker-compose.yaml down