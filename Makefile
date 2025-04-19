run-build:
	docker-compose up --build

run:
	docker-compose up

run-file:
	go run cmd/*.go  

test:
	go test ./tests/* -v

reset:
	docker-compose down --rmi all --volumes && docker-compose up --build
