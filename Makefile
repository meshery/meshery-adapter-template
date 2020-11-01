docker:
	docker build -t layer5/meshery-<adaptor-name> .

docker-run:
	(docker rm -f meshery-<adaptor-name>) || true
	docker run --name meshery-<adaptor-name> -d \
	-p 10007:10007 \
	-e DEBUG=true \
	layer5/meshery-<adaptor-name>

run:
	DEBUG=true go run main.go