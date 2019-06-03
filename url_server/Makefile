docker-build:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./url_server
	docker build -t nozomi0966/google_like_search_engine_url_server .

docker-push:
	docker push nozomi0966/google_like_search_engine_url_server

build:
	go build 

