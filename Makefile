build:
	- docker image build -t transfer-api .

run:
	- docker container run -p 8888:8888 transfer-api fresh

fresh:
	- docker run transfer-api