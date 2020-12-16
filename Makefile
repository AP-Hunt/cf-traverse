build:
	go build -o bin/cf-traverse .

install: build
	cf install-plugin -f ./bin/cf-traverse