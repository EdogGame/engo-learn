all:
	go build -v -o bin/engo-learn ./src/engo-learn

mac:
	GOOS=darwin GOARCH=amd64 go build -v -o bin/engo-learn-darwin-amd64 ./src/engo-learn

windows:
	GOOS=windows GOARCH=386 go build -v -o bin/engo-learn-windows-386.exe ./src/engo-learn

windows64:
	GOOS=windows GOARCH=amd64 go build -v -o bin/engo-learn-windows-amd64.exe ./src/engo-learn