all:
	go build -o ./bin/klbs

install:
	go build -o ./bin/klbs
	cp ./bin/klbs ~/.local/bin/klbs

clean:
	rm -rf ./dist
	rm -rf ./bin
