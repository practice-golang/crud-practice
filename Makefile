build:
	go build -ldflags "-w -s" -trimpath -o bin/

test:
	go test ./... -race -cover

clean:
	rm -rf bin/*
