protoc --go_out=fourth_people --proto_path=..\proto\ ..\proto\*.proto
go run server.go
