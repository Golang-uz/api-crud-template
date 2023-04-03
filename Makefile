include .env

run:
	go run cmd/main.go
get: 
	go mod tidy
login:
	docker exec -it ${DOCKER_CONTAINER_POSTGRES_NAME} psql ${POSTGRES_DB} ${POSTGRES_USER} 
stop: 
	docker stop ${DOCKER_CONTAINER_POSTGRES_NAME}
start: 
	docker start ${DOCKER_CONTAINER_POSTGRES_NAME}
swag:	
	swag init -g ./cmd/main.go -o ./docs
