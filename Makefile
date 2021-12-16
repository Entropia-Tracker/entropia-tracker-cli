VERSION=`git describe --tags`
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

build:
	go build ${LDFLAGS} -o bin/entropia-tracker-cli ./cmd

.PHONY: bin
clean: bin
	rm -rf bin

test:
	go test -count=50 ./...

release: clean
	GOARCH=386 GOOS=windows go build ${LDFLAGS} -o bin/entropia-tracker-cli.386.exe ./cmd
	cd bin && sha512sum entropia-tracker-cli.386.exe > entropia-tracker-cli.386.exe.sha512
	cd bin && sha512sum -c entropia-tracker-cli.386.exe.sha512

	GOARCH=amd64 GOOS=windows go build ${LDFLAGS} -o bin/entropia-tracker-cli.amd64.exe ./cmd
	cd bin && sha512sum entropia-tracker-cli.amd64.exe > entropia-tracker-cli.amd64.exe.sha512
	cd bin && sha512sum -c entropia-tracker-cli.amd64.exe.sha512

	GOARCH=amd64 GOOS=linux go build ${LDFLAGS} -o bin/entropia-tracker-cli.unix ./cmd
	cd bin && sha512sum entropia-tracker-cli.unix > entropia-tracker-cli.unix.sha512
	cd bin && sha512sum -c entropia-tracker-cli.unix.sha512

	GOARCH=amd64 GOOS=darwin go build ${LDFLAGS} -o bin/entropia-tracker-cli.macos ./cmd
	cd bin && sha512sum entropia-tracker-cli.macos > entropia-tracker-cli.macos.sha512
	cd bin && sha512sum -c entropia-tracker-cli.macos.sha512
