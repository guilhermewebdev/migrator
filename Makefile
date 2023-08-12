build:
	SRC_DIR="." ./bin/build
run:
	SRC_DIR="." ./bin/dev
clean:
	SRC_DIR="." ./bin/clean
unit:
	SRC_DIR="." ./bin/unit
e2e:
	SRC_DIR="." ./bin/e2e
fulltest:
	SRC_DIR="." ./bin/fulltest
install:
	sudo cp ./bin/migrate /usr/local/bin/
uninstall:
	sudo rm /usr/local/bin/migrate
images:
	docker build -t migrator:build -f docker/Dockerfile.build .
	docker build -t migrator:alpine -f docker/Dockerfile.alpine .
	docker build -t migrator:bullseye -t migrator:latest -f docker/Dockerfile.bullseye .
	docker build -t migrator:scratch -f docker/Dockerfile.scratch .
	docker build -t migrator:bookworm -f docker/Dockerfile.bookworm .
