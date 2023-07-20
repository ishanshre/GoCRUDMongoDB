createMongoDBcontainer:
	docker run -d -p 27017:27017 --name=MongoDBCRUD -v mongo_data:/data/db mongo

startContainer:
	docker start MongoDBCRUD

run:
	go run cmd/api/main.go    