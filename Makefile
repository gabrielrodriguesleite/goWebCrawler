subirContainerDB:
	@docker run -d --name mongodb -p 27017:27017 mongo

visitarDB:
	@echo "TERMINAL DO CONTAINER: instruções:"
	@echo -e 'execute:\n- mongosh\n- show dbs\n- use crawler\n- \
		show collections\n- db.links.countDocuments()\n- db.links.find({})'
	@docker exec -it mongodb /bin/bash

runCrawler:
	@go run main.go -action=webcrawler

runWebsite:
	@go run main.go