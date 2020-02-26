package main

import (
	"futong_server_agent_go/serve"
)

/*
	交叉编译
	env GOOS=linux GOARCH=amd64 go build -o futongAgent.linux-amd64
	env GOOS=linux GOARCH=386 go build -o futongAgent.linux-386

	env GOOS=windows GOARCH=amd64 go build -o futongAgent.windows-amd64
*/

func main() {
	serve.Service.StartServe()
}


