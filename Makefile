run-server :
	go run server/main.go

run-client :
	go run client/main.go

get-and-run-redis-docker :
	sudo docker run --name redis-server -p 6379:6379 -d redis	