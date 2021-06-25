#.SILENT:
#
#.PHONY: build
#build:
#	go build -o forum -v ./cmd/
#
#.PHONY: git
#git:
#	git add .
#	git commit -m "$(comment)"
#	git push
#
#.PHONY: git-conf
#git-conf:
#	git config --global user.email "$(email)"
#	git config --global user.name "$(name)"
#
#.PHONY: docker-run
#docker-run:
#	docker build -t "$(image-name)" .
#	docker container run -p 17555:17555 --name "$(container-name)" "$(image-name)"
#
#.DEFAULT_GOAL := build

.PHONY: all clean

all: build run

clear:
	clear
	rm -rf ./assets/database.db
	rm -rf ./bin/app
	rm -rf ./assets/images

build:
	go build -v -o ./bin/app ./cmd/main.go

run:
	./bin/app

