.PHONY: gofull
gofull:
	CGO_ENABLED=0 GOOS=linux go build -i -o ./examples/go-full/app -ldflags '-s' -installsuffix cgo ./examples/go-full/main.go
	docker-compose -f docker-compose.yml -f docker-compose.gofull.yml up --build & (sleep 5 && echo 'GET http://localhost:8888/demo/' | vegeta attack)

.PHONY: generic
generic:
	docker-compose up

.PHONY: nodejs
nodejs:
	docker-compose -f docker-compose.yml -f docker-compose.nodejs.yml up --build


.PHONY: dotnet
dotnet:
	docker-compose -f docker-compose.yml -f docker-compose.dotnet.yml up --build