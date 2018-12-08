all: clean

clean: 
	find . -name '*.log' -type f -exec rm -f {} \; 

run:
	node app.js

docker-build: 
	docker build -t clipper-frontend .
	docker images

docker-run:
	docker run -d --name clipper-frontend -td -p 3500:3500 clipper-frontend

docker-shell:
	docker exec -it clipper-frontend bash

docker-clean:
	docker stop clipper-frontend
	docker rm clipper-frontend
