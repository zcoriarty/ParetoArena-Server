source .env
# export JWT_SECRET=$(openssl rand -base64 256)
# echo $JWT_SECRET
# go get -v ./...
# go run ./entry/ generate_secret
# echo $JWT_SECRET

go run ./entry create_db
go run ./entry create_schema
go run ./entry create_superadmin -e test_super_admin@gmail.com -p password
# go run ./entry create_test_user -e connor@paretotest.com -p connorpass -a 0b34501e-33f3-42e4-96f5-ae57a17e436a

go run ./entry/main.go


