build: win linux mac
target = ./main.go

win:
	GOOS=windows GOARCH=amd64 go build -o ./builds/iciba.exe $(target)

linux:
	GOOS=linux GOARCH=amd64 go build -o ./builds/iciba $(target)

mac:
	GOOS=darwin GOARCH=amd64 go build -o ./builds/iciba_mac $(target)

test:
	go test -v ./...
