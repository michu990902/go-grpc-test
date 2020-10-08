# CA's private key and self signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -keyout ca-key.pem -out ca-cert.pem

# Server private key and certificate signing request
openssl req -newkey rsa:4096 -keyout server-key.pem -out server-req.pem

# Use CA'S private key to sign web server CSR and get back the signed certificate
openssl x509 -req -in server-req.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem

# test
openssl verify -CAfile ca-cert.pem server-cert.pem
