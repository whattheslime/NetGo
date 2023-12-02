folder = build

all: windows linux

certifs:
	mkdir -p certs

	echo "[+] Make server cert"
	openssl req -new -nodes -x509 -out certs/server.pem -keyout certs/server.key -days 3650 -subj "/C=US/ST=NRW/L=Earth/O=Company/OU=IT/CN=www.random.com/emailAddress=john@doe.com"

	echo "[+] Make client cert"
	openssl req -new -nodes -x509 -out certs/client.pem -keyout certs/client.key -days 3650 -subj "/C=US/ST=NRW/L=Earth/O=Company/OU=IT/CN=www.random.com/emailAddress=john@doe.com"

windows:
	mkdir -p $(folder)
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $(folder)/netgo.exe .

linux:
	mkdir -p $(folder)
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $(folder)/netgo .

darwin:
	mkdir -p $(folder)
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o $(folder)/netgo .

clean:
	rm $(folder)/*