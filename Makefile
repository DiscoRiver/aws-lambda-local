build:
	echo "Compiling..."
	GOOS=linux GOARCH=amd64 go build -o build/main-linux-amd64 main.go

zip:
	zip build/main-linux-amd64.zip build/main-linux-amd64

all: build zip