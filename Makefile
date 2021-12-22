VERSION=`git describe --tags`
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

build:
	go build ${LDFLAGS} -o bin/entropia-tally-cli ./cmd

.PHONY: bin
clean: bin
	rm -rf bin

test:
	go test -count=50 ./...

release: clean
	GOARCH=386 GOOS=windows go build ${LDFLAGS} -o bin/entropia-tally-cli.386.exe ./cmd
	cd bin && sha512sum entropia-tally-cli.386.exe > entropia-tally-cli.386.exe.sha512
	cd bin && sha512sum -c entropia-tally-cli.386.exe.sha512

	GOARCH=amd64 GOOS=windows go build ${LDFLAGS} -o bin/entropia-tally-cli.amd64.exe ./cmd
	cd bin && sha512sum entropia-tally-cli.amd64.exe > entropia-tally-cli.amd64.exe.sha512
	cd bin && sha512sum -c entropia-tally-cli.amd64.exe.sha512

	GOARCH=amd64 GOOS=linux go build ${LDFLAGS} -o bin/entropia-tally-cli.unix ./cmd
	cd bin && sha512sum entropia-tally-cli.unix > entropia-tally-cli.unix.sha512
	cd bin && sha512sum -c entropia-tally-cli.unix.sha512

	GOARCH=amd64 GOOS=darwin go build ${LDFLAGS} -o bin/entropia-tally-cli.macos ./cmd
	cd bin && sha512sum entropia-tally-cli.macos > entropia-tally-cli.macos.sha512
	cd bin && sha512sum -c entropia-tally-cli.macos.sha512
