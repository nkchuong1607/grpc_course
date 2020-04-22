
# Private: ca.key, server.key,server.pem, server.crt
# Share: ca.crt (client), server.csr

#server common name
SERVER_CN=localhost

# Generate Certificate Authority and Trust Certificate (ca.crt)
openssl genrsa -passout pass:1234 -des3 -out ca.key 1024
openssl req -passin pass:1234 -new -x509 -days 365 -key ca.key -out ca.crt -subj "/CN={$SERVER_CN}"

# Generate you server key
openssl genrsa -passout pass:1234 -des3 -out server.key 1024

# Generate your Certificate Signing Request (CSR)
openssl req -passin pass:1234 -new -key server.key -out server.csr -subj "/CN={$SERVER_CN}"

# Sign the certificate with the CA
openssl x509 -req -passin pass:1234 -days 365 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt

# Convert server certificate to server.pem
openssl pkcs8 -topk8 -nocrypt -passin pass:1234 -in server.key -out server.pem