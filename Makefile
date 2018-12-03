all: clean

clean:
	find . -name 'userclipper' -type f -exec rm -f {} \;
	go clean

run:
	go run src/app/$(app).go

main:
	go run src/app/main.go

format:
	go fmt userclipper

install:
	go install userclipper

go-get:
	rm -rf src/github.com
	go get -v github.com/BurntSushi/toml
	go get -v github.com/gorilla/mux
	go get -v gopkg.in/mgo.v2
	go get -v gopkg.in/mgo.v2/bson

build:
	go build userclipper

start:
	./userclipper

docker-build:
	docker build -t userclipper .
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
	docker run --network clipper --name mongoclipper -p 3000:3000 -td --env MONGO_SERVER=mongodb --env MONGO_DB=cardholders_db userclipper
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
	docker stop userclipper
	docker rm userclipper
	docker rmi userclipper

docker-ip:
	docker-machine ip

docker-ps:
	 docker ps --all --format "table {{.ID}}\t{{.Names}}\t{{.Image}}\t{{.Status}}\t"

docker-ps-ports:
	 docker ps --all --format "table {{.Names}}\t{{.Ports}}\t"

test-ping:
	curl localhost:3000/ping

test-get-cardholders:
	curl localhost:3000/cardholders

test-get-cardholders-byid:
	curl localhost:3000/cardholders/{id}

test-create-user:
	curl -X POST \
  	http://localhost:3000/cardholders \
  	-H 'Content-Type: application/json'

test-update-user:
	curl -X PUT \
  	http://localhost:3000/cardholders \
  	-H 'Content-Type: application/json'

test-delete-user:
	curl -X DELETE \
  	http://localhost:3000/cardholders \
  	-H 'Content-Type: application/json'
