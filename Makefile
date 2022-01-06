build:
	go build -ldflags "-w -s" -trimpath -o bin/

clean:
	rm -rf bin/*
