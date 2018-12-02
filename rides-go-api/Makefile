
all: clean

clean:
	find . -name 'clipper' -type f -exec rm -f {} \;
	go clean

run:
	go run src/app/$(app).go

main:
	go run src/app/main.go

format:
	go fmt clipper

install:
	go install clipper

go-get:
	rm -rf src/github.com
	go get -v github.com/BurntSushi/toml
	go get -v github.com/gorilla/mux
	go get -v gopkg.in/mgo.v2
	go get -v gopkg.in/mgo.v2/bson

build:
	go build clipper

start:
	./clipper

docker-build:
	docker build -t clipper .
	docker images

network-create:
	docker network create --driver bridge clipper

network-inspect:
	docker network inspect clipper

rabbitmq-run:
	docker run --name rabbitmq --network clipper \
			   -p 8080:15672 -p 4369:4369 -p 5672:5672 \
			   -d rabbitmq:3-management
mongodb-run:
	docker run --name mongodb --network clipper -p 27017:27017 -d mongo:3.7

docker-run:
	docker run --network clipper --name clipper -p 3000:3000 -td --env MONGO_SERVER=mongodb --env MONGO_DB=ridesDB  clipper
	docker ps

docker-network:
	docker network ls

docker-network-prune:
	docker network prune

docker-network-inspect:
	docker network inspect host

docker-shell:
	docker exec -it clipper bash

docker-clean:
	docker stop mongodb
	docker stop rabbitmq
	docker rm mongodb
	docker rm rabbitmq
	docker stop clipper
	docker rm clipper
	docker rmi clipper

docker-ip:
	docker-machine ip

docker-ps:
	 docker ps --all --format "table {{.ID}}\t{{.Names}}\t{{.Image}}\t{{.Status}}\t"

docker-ps-ports:
	 docker ps --all --format "table {{.Names}}\t{{.Ports}}\t"

test-ping:
	curl localhost:3000/ping

test-get-inventory:
	curl localhost:3000/clipper

test-update-inventory:
	curl -X PUT \
  	http://localhost:3000/clipper \
  	-H 'Content-Type: application/json' \
  	-d '{ \
  		"Countclippers": 1000 }'

test-place-order:
	curl -X POST \
  	http://localhost:3000/order \
  	-H 'Content-Type: application/json'

test-order-status:
	curl -X GET \
  	http://localhost:3000/order \
  	-H 'Content-Type: application/json'

test-process-order:
	curl -X POST \
  	http://localhost:3000/orders \
  	-H 'Content-Type: application/json'
