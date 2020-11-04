
openssl genrsa -out ssl/server.key 2048

openssl req -nodes -new -x509 -sha256 -days 1825 -config ssl/cert.conf -extensions 'req_ext' -key ssl/server.key -out ssl/server.crt