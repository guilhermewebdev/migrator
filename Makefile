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