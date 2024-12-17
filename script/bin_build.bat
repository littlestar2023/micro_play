SET GOOS=linux
SET GOARCH=amd64

go build -o ../docker/user/user_service ../app/user/cmd/main.go
go build -o ../docker/task/task_service ../app/task/cmd/main.go
