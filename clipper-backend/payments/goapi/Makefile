
all: clean

clean: 
	find . -name 'payments' -type f -exec rm -f {} \; 
	go clean

run:
	go run src/app/$(app).go

main:
	go run src/app/main.go

format:
	go fmt payments

install:
	go install payments

go-get:
	rm -rf src/github.com
	go get -v github.com/codegangsta/negroni
	go get -v github.com/gorilla/mux
	go get -v github.com/unrolled/render
	go get -v github.com/satori/go.uuid
	go get -v github.com/streadway/amqp
	go get -v gopkg.in/mgo.v2
	go get -v gopkg.in/mgo.v2/bson
	go get -v github.com/rs/cors

build:
	go build payments

start:
	./payments 

docker-build: 
	docker build -t payments .
	docker images

network-create:
	docker network create --driver bridge payments

network-inspect:
	docker network inspect payments

rabbitmq-run:
	docker run --name rabbitmq --network payments \
			   -p 8080:15672 -p 4369:4369 -p 5672:5672 \
			   -d rabbitmq:3-management
mongodb-run:
	docker run --name mongodb --network payments -p 27017:27017 -d mongo:3.7

docker-run:
	docker run --network payments \
			--name payments -p 5000:5000 -td payments
	docker ps

docker-network:
	docker network ls

docker-network-prune:
	docker network prune

docker-network-inspect:
	docker network inspect host

docker-shell:
	docker exec -it payments bash 

docker-clean:
	docker stop mongodb
	docker stop rabbitmq
	docker rm mongodb
	docker rm rabbitmq
	docker stop payments
	docker rm payments
	docker rmi payments

docker-ip:
	docker-machine ip

docker-ps:
	 docker ps --all --format "table {{.ID}}\t{{.Names}}\t{{.Image}}\t{{.Status}}\t"

docker-ps-ports:
	 docker ps --all --format "table {{.Names}}\t{{.Ports}}\t"

test-ping:
	curl localhost:5000/ping
