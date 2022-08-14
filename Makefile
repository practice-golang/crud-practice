build:
	go build -ldflags "-w -s -X 'main.appname=crud-practice'" -trimpath -o bin/

dist:
	go get -d github.com/mitchellh/gox
	go build -mod=readonly -o ./bin/ github.com/mitchellh/gox
	go mod tidy
	go env -w GOFLAGS=-trimpath
	./bin/gox -mod="readonly" -ldflags="-w -s -X 'main.appname=crud-practice'" -output="bin/{{.Dir}}_{{.OS}}_{{.Arch}}" -osarch="windows/amd64 linux/amd64 linux/arm darwin/amd64 darwin/arm64"
	rm ./bin/gox*

test:
	go test ./... -race -cover

clean:
	rm -rf bin/*
