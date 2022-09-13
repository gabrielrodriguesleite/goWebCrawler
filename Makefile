subirContainerDB:
	@docker run -d --name mongodb -p 27017:27017 mongo

run:
	@go run main.go