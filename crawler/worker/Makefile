docker-build:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./crawler
	docker build -t nozomi0966/google_like_search_engine_crawler .

docker-push:
	docker push nozomi0966/google_like_search_engine_crawler 

build:
	go build 

