.PHONY: build clean deploy

build:
	#dep ensure -v
	env go build -ldflags="-s -w"  -o bin/wootscrape  cmd/main/main.go

install:
	mv ./bin/wootscrape /usr/local/bin

clean:
	rm -rf ./bin ./vendor Gopkg.lock
