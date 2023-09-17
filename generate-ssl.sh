mkcert -key-file pareto.com.pem -cert-file pareto-public.com.pem pareto.com

# for client request encryption
openssl genrsa -out private_key.pem 1024
openssl rsa -in private_key.pem -outform PEM -pubout -out public_key.pem
