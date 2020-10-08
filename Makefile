pwd = $(shell pwd)

start-all:
	- docker-compose up

# caso os containers "travem a roda"
stop-all:
	- docker stop $(docker ps -a -q)

remove-all:
	- docker rm $(docker ps -a -q)