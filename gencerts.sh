#!/bin/bash
if ! [[ -d "certs" ]]
then
    mkdir certs
fi
rm certs/*

echo "make server cert"
openssl req -new -nodes -x509 -out certs/server.pem -keyout certs/server.key -days 3650 -subj "/C=US/ST=NRW/L=Earth/O=Company/OU=IT/CN=www.random.com/emailAddress=john@doe.com"
echo "make client cert"
openssl req -new -nodes -x509 -out certs/client.pem -keyout certs/client.key -days 3650 -subj "/C=US/ST=NRW/L=Earth/O=Company/OU=IT/CN=www.random.com/emailAddress=john@doe.com"
