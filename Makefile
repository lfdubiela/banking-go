DOCKER_CONTAINER:= docker exec banking-go-app bash -c

pwd = $(shell pwd)

start-all:
	- docker-compose up

stop-all:
	- docker stop $(docker ps -a -q)

remove-all:
	- docker rm $(docker ps -a -q)

test:
	$(DOCKER_CONTAINER) 'go test -covermode="count" ./...'

test-coverage:
	$(DOCKER_CONTAINER) './go-coverage.sh' && sensible-browser report/coverage.html