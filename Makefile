docker-broker:
	sudo docker build -f Dockerfile.Broker . -t broker:latest
	sudo docker run --rm --name Broker -p 50051:50051 -p 50052:50052 broker:latest

docker-vanguardia:
	sudo docker build -f Dockerfile.Vanguardia . -t vanguardia:latest
	sudo docker run --rm --name Vanguardia --network="host" vanguardia:latest

docker-f1:
	sudo docker build -f Dockerfile.F1 . -t f1:latest
	sudo docker run --rm --name F1 -p 50051:50051 -p 50052:50052 f1:latest

docker-f2:
	sudo docker build -f Dockerfile.F2 . -t f2:latest
	sudo docker run --rm --name F2 -p 50051:50051 -p 50052:50052 f2:latest

docker-f3:
	sudo docker build -f Dockerfile.F3 . -t f3:latest
	sudo docker run --rm --name F3 -p 50051:50051 -p 50052:50052 f3:latest

docker-i1:
	sudo docker build -f Dockerfile.I1 . -t i1:latest
	sudo docker run --rm --name I1 -p 50051:50051 -p 50052:50052 i1:latest

#docker-i2: