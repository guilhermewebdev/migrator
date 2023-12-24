build:
	./bin/build
run:
	./bin/dev
clean:
	./bin/clean
unit:
	./bin/unit
e2e:
	./bin/e2e
fulltest:
	./bin/fulltest
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
	docker build -t migrator:debian -f docker/Dockerfile.debian .
	docker build -t migrator:slim -f docker/Dockerfile.slim .
	docker build -t migrator:latest -f docker/Dockerfile .
