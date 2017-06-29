all: clean cmd/platform-starter/data.go

clean:
	rm -f cmd/platform-starter/data.go


cmd/platform-starter/data.go:
	go-bindata -o cmd/platform-starter/data.go ./config ./hooks
