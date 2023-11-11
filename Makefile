# https://tutorialedge.net/golang/makefiles-for-go-developers/
compile:
	echo "\n\n Compiling for every OS and Platform\n\t Results in folder build"
	GOOS=freebsd GOARCH=amd64 go build -o build/termk-freebsd-amd64 main.go
	GOOS=linux GOARCH=amd64 go build -o build/termk-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o build/termk-windows-amd64.exe main.go
	GOOS=darwin  GOARCH=amd64 go build -o build/termk-darwin-amd64 main.go

build: clean
	mkdir -p build
	go build -o build/termk main.go
clean:
	@rm -r build/
	@echo "\n\n clean, run again make build for generate termK in actual platform\n\n"
help:
	@echo "make build \n\t generate termK for actual platform"
	@echo "make compile \n\t generate termK for windows linux bsd and darwin"
	@echo "make clean \n\t clean build"
	@echo "make run \n\t execute termK from source"
	@echo "make help \n\t show help"
run:
	go run main.go
