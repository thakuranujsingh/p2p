#Server run
**Run four server into saperate terminal**
   go run server/main.go --port="8080"
   go run server/main.go --port="8081"
   go run server/main.go --port="8082"
   go run server/main.go --port="8083"


#Client start

**Add transaction**
go run client/main.go --add 

**Get all transaction**
go run client/main.go --list

#Compile proto file
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/blockchain.proto